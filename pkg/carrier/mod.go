package carrier

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/pairing"
	"go.dedis.ch/kyber/v4/sign/bdn"
	"io/ioutil"
	"math"
	"net"
	"sync"
)

type Carrier struct {
	counter uint64

	config Config

	id string

	listeners Listeners
	stores    Stores
	locks     Locks
	addresses Addresses

	nodeConn *net.TCPConn
	node     *Node

	neighbours map[string]*Neighbour

	// Registry of message handlers. Argument must be one of the enum types
	messageHandlers map[message.Type]func(message.Message) error

	suite   *pairing.SuiteBn256
	keypair *util.KeyPairString

	f int
	n int

	wg *sync.WaitGroup

	quit               chan bool
	broadcastDispenser chan message.Message
}

func Load(file string) (*Carrier, error) {
	// Read config file
	rawdata, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	config := &util.Config{}
	err = json.Unmarshal(rawdata, config)
	if err != nil {
		return nil, err
	}

	// Create carrier.Neighbour structs and collect them into a map for fast access
	neighbours := map[string]*Neighbour{}
	for _, n := range config.Neighbours {
		neighbours[n.ID] = NewNeighbour(n.ID, n.Address, n.PK)
	}

	// Build a new carrier node
	return NewCarrier(config.ID, config.Addresses.Client, config.Addresses.Carrier, config.Addresses.Front, config.Addresses.Decision, neighbours, &config.Keys), nil
}

func NewCarrier(id, clientToCarrierAddr, carrierToCarrierAddr, frontAddr, decisionAddr string, carriers map[string]*Neighbour, keypair *util.KeyPairString) *Carrier {

	c := &Carrier{}
	c.quit = make(chan bool, 1)
	c.id = id

	// TEMP
	c.counter = 0
	c.broadcastDispenser = make(chan message.Message, int64(math.Pow10(7)))

	c.suite = pairing.NewSuiteBn256()
	c.keypair = keypair

	c.node = NewNode(frontAddr)
	c.neighbours = carriers

	c.f = (len(carriers) - 1) / 3
	c.n = len(carriers)

	c.config = Config{
		carrierConnRetryDelay: util.CarrierConnRetryDelay,
		carrierConnMaxRetry:   util.CarrierConnMaxRetry,
		nodeConnRetryDelay:    util.NodeConnRetryDelay,
		nodeConnMaxRetry:      util.NodeConnMaxRetry,
	}

	c.messageHandlers = map[message.Type]func(message.Message) error{
		message.Init:    c.handleInitMessage,
		message.Echo:    c.handleEchoMessage,
		message.Request: c.handleRequestMessage,
		message.Resolve: c.handleResolveMessage,
	}

	c.locks = Locks{
		ValueStore:        &sync.RWMutex{},
		SignatureStore:    &sync.RWMutex{},
		SuperBlockSummary: &sync.RWMutex{},
		AcceptedHashStore: &sync.RWMutex{},
		DecisionLock:      &sync.RWMutex{},
	}

	c.stores = Stores{
		valueStore:        map[string][][]byte{},
		signatureStore:    map[string][]util.Signature{},
		superBlockSummary: map[string][]util.Signature{},
		acceptedHashStore: map[string][][]byte{},
	}

	c.addresses = Addresses{
		client2carrier:  clientToCarrierAddr,
		carrier2carrier: carrierToCarrierAddr,
		decision:        decisionAddr,
	}

	var err error

	// Check client listener addr
	_, err = util.ResolveTCPAddr(clientToCarrierAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
		return nil
	}

	// Check carrier listener addr
	_, err = util.ResolveTCPAddr(carrierToCarrierAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
		return nil
	}

	// Check front addr
	_, err = util.ResolveTCPAddr(frontAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
		return nil
	}

	// Check other carrier addrs
	for _, n := range c.neighbours {
		_, err := util.ResolveTCPAddr(n.Address)
		if err != nil {
			log.Error().Msgf(err.Error())
		}
	}

	return c
}

/*	Start listening to client requests
	Forward client requests
	We are not waiting for listeners to stop but I think it's fine
*/
func (c *Carrier) Start() *sync.WaitGroup {
	c.wg = &sync.WaitGroup{}
	c.wg.Add(1)

	var err error

	// Listen to nested SMR decisions
	decision, err := util.ResolveTCPAddr(c.getDecisionAddress())
	if err != nil {
		log.Error().Msgf(err.Error())
		c.wg.Done()
		return c.wg
	}
	c.listeners.decisionListener, err = util.ListenTCP(decision)
	if err != nil {
		log.Error().Msgf(err.Error())
		c.wg.Done()
		return c.wg
	}
	log.Info().Msgf("start listening to nested SMR decisions on %s", c.getDecisionAddress())
	go c.handleIncomingConnections(c.listeners.decisionListener, c.decodeNestedSMRDecisions)

	// Start client listener
	client, err := util.ResolveTCPAddr(c.getClientToCarrierAddress())
	if err != nil {
		log.Error().Msgf(err.Error())
		c.wg.Done()
		return c.wg
	}
	c.listeners.clientListener, err = util.ListenTCP(client)
	if err != nil {
		log.Error().Msgf(err.Error())
		c.wg.Done()
		return c.wg
	}
	log.Info().Msgf("start listening to client on %s", c.getClientToCarrierAddress())
	go c.handleIncomingConnections(c.listeners.clientListener, c.handleClientConn)

	// Start carrier listener
	carrier, err := util.ResolveTCPAddr(c.getCarrierToCarrierAddress())
	if err != nil {
		log.Error().Msgf(err.Error())
		c.wg.Done()
		return c.wg
	}
	c.listeners.carrierListener, err = util.ListenTCP(carrier)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	log.Info().Msgf("start listening to neighbours on %s", c.getCarrierToCarrierAddress())
	go c.handleIncomingConnections(c.listeners.carrierListener, c.handleCarrierConn)

	//Connect to node
	connect(c.node, c.config.nodeConnRetryDelay, c.config.nodeConnMaxRetry)

	// Set up connections to other neighbours
	for _, n := range c.neighbours {
		go connect(n, c.config.carrierConnRetryDelay, c.config.carrierConnMaxRetry)
	}

	//go c.logger()
	c.launchWorkerPool(10, c.broadcastWorker)

	return c.wg
}

func (c *Carrier) Stop() {
	log.Trace().Msgf("stop Carrier")
	err := c.listeners.clientListener.Close()
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	c.quit <- true
	c.wg.Done()
}

func (c *Carrier) GetAddress() string {
	return c.getCarrierToCarrierAddress()
}

// TODO
// Rust node expects to receive uint32 to describe size of tsx
// Implement this framing
// Implement decoding
// Implement reply
func (c *Carrier) NestedPropose(P SuperBlockSummary) error {

	err := c.node.GetEncoder().Encode(&P)
	if err != nil {
		return err
	}

	log.Info().Msgf("proposed %s to %s", P, c.node.GetAddress())
	return nil
}

func (c *Carrier) GetStringSK() string {
	return c.keypair.Sk
}

func (c *Carrier) GetStringPK() string {
	return c.keypair.Pk
}

func (c *Carrier) GetKyberSK() kyber.Scalar {
	skk, err := util.DecodeStringToBdnSK(c.keypair.Sk)
	if err != nil {
		panic("unable to decode SK") //TODO
	}

	return skk
}

func (c *Carrier) GetKyberPK() kyber.Point {
	pkk, err := util.DecodeStringToBdnPK(c.keypair.Pk)
	if err != nil {
		panic("unable to decode SK") //TODO
	}

	return pkk
}

func (c *Carrier) Sign(h string) string {
	hb, err := hex.DecodeString(h)
	if err != nil {
		panic("signing failed: failed to decode h")
	}
	s, err := bdn.Sign(c.suite, c.GetKyberSK(), hb)
	if err != nil {
		panic("signing failed")
	}

	return hex.EncodeToString(s)
}

func (c *Carrier) Verify(h string, s util.Signature) error {
	pk, err := c.GetPKFromID(s.SenderID)
	if err != nil {
		return err
	}

	hb, err := hex.DecodeString(h)
	if err != nil {
		panic("verify failed: failed to decode h")
	}
	sb, err := hex.DecodeString(s.S)
	if err != nil {
		panic("verify failed: failed to decode s")
	}
	err = bdn.Verify(c.suite, pk, hb, sb)
	return err
}

//func (c *Carrier) VerifyPK(pk string) (kyber.Point, error) {
//	//_, ok := c.neighbours[pk]
//	//if !ok {
//	//	return nil, fmt.Errorf("unrecognised sender")
//	//}
//
//	pkk, err := util.DecodeStringToBdnPK(pk)
//	if err != nil {
//		return nil, err
//	}
//
//	return pkk, nil
//}

func (c *Carrier) GetPKFromID(senderID string) (kyber.Point, error) {

	n, ok := c.neighbours[senderID]
	if !ok {
		return nil, fmt.Errorf("pk not found in store")
	}

	pkk, err := util.DecodeStringToBdnPK(n.PK)
	if err != nil {
		return nil, err
	}

	return pkk, nil
}

func (c *Carrier) GetSuite() pairing.Suite {
	return c.suite.Suite
}

func (c *Carrier) decide(oldD map[string][][]byte) {
	defer c.locks.DecisionLock.Unlock()
	D := oldD
	c.stores.acceptedHashStore = map[string][][]byte{}
	log.Info().Msgf("decided %s", D)
}

func (c *Carrier) getClientToCarrierAddress() string {
	return c.addresses.client2carrier
}

func (c *Carrier) getCarrierToCarrierAddress() string {
	return c.addresses.carrier2carrier
}

func (c *Carrier) getDecisionAddress() string {
	return c.addresses.decision
}

func (c *Carrier) GetID() string {
	return c.id
}

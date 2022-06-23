package carrier

import (
	"encoding/hex"
	"fmt"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/pairing"
	"go.dedis.ch/kyber/v4/sign/bdn"
	"math"
	"net"
	"sync"
	"time"
)

type Carrier struct {
	counter uint64

	config *util.Config

	listeners Listeners
	stores    Stores
	locks     Locks

	nodeConn *net.TCPConn
	node     *Node

	neighbours map[string]*Neighbour

	// Registry of message handlers. Argument must be one of the enum types
	messageHandlers map[message.Type]func(message.Message) error

	suite *pairing.SuiteBn256

	f int
	n int

	wg *sync.WaitGroup

	quit               chan bool
	broadcastDispenser chan message.Message
	sbsCounter         int
}

func NewCarrier(config *util.Config) *Carrier {
	neighbours := map[string]*Neighbour{}
	for _, n := range config.Neighbours {
		neighbours[n.ID] = NewNeighbour(n.ID, n.Address, n.PK)
	}

	c := &Carrier{}
	c.config = config
	c.quit = make(chan bool, 1)

	// TEMP
	c.counter = 0
	c.broadcastDispenser = make(chan message.Message, int64(math.Pow10(7)))
	c.sbsCounter = 0

	c.suite = pairing.NewSuiteBn256()

	c.node = NewNode(config.Addresses.Front)
	c.neighbours = neighbours

	c.f = (len(neighbours) - 1) / 3
	c.n = len(neighbours)

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

		decidedHashStore: map[string]any{},
	}

	//var err error

	//// Check client listener addr
	//_, err = util.ResolveTCPAddr(clientToCarrierAddr)
	//if err != nil {
	//	log.Error().Msgf(err.Error())
	//	return nil
	//}
	//
	//// TODO move this
	//// Check carrier listener addr
	//_, err = util.ResolveTCPAddr(carrierToCarrierAddr)
	//if err != nil {
	//	log.Error().Msgf(err.Error())
	//	return nil
	//}
	//
	//// Check front addr
	//_, err = util.ResolveTCPAddr(frontAddr)
	//if err != nil {
	//	log.Error().Msgf(err.Error())
	//	return nil
	//}
	//
	//// Check other carrier addrs
	//for _, n := range c.neighbours {
	//	_, err := util.ResolveTCPAddr(n.Address)
	//	if err != nil {
	//		log.Error().Msgf(err.Error())
	//	}
	//}

	return c
}

/*	Start listening to client requests
	Forward client requests
	We are not waiting for listeners to stop but I think it's fine
*/
func (c *Carrier) Start() *sync.WaitGroup {
	if c.forwardMode() {
		log.Info().Msgf("ForwardMode is turned on - logging of tsx is at debug level to avoid flooding. If you want to see individual logs, set log level to debug or higher.")
	}

	c.wg = &sync.WaitGroup{}
	c.wg.Add(1)

	var err error

	// Listen to nested SMR decisions
	c.listeners.decisionListener, err = c.startListener(c.getDecisionAddress())
	if err != nil {
		log.Error().Msgf(err.Error())
		c.wg.Done()
		return c.wg
	}
	log.Info().Msgf("start listening to nested SMR decisions on %s", c.getDecisionAddress())
	go c.handleIncomingConnections(c.listeners.decisionListener, c.decodeNestedSMRDecisions)

	// Listen to client
	c.listeners.clientListener, err = c.startListener(c.getClientToCarrierAddress())
	if err != nil {
		log.Error().Msgf(err.Error())
		c.wg.Done()
		return c.wg
	}
	log.Info().Msgf("start listening to client on %s", c.getClientToCarrierAddress())
	go c.handleIncomingConnections(c.listeners.clientListener, c.handleClientConn)

	// Listen to carrier
	c.listeners.carrierListener, err = c.startListener(c.getCarrierToCarrierAddress())
	if err != nil {
		log.Error().Msgf(err.Error())
		c.wg.Done()
		return c.wg
	}
	log.Info().Msgf("start listening to neighbours on %s", c.getCarrierToCarrierAddress())
	go c.handleIncomingConnections(c.listeners.carrierListener, c.handleCarrierConn)

	//Connect to node
	go connect(c.node, time.Duration(c.config.Settings.NodeConnRetryDelay)*time.Millisecond, c.config.Settings.NodeConnMaxRetry)

	// Set up connections to other neighbours
	for _, n := range c.neighbours {
		go connect(n, time.Duration(c.config.Settings.CarrierConnRetryDelay)*time.Millisecond, c.config.Settings.CarrierConnMaxRetry)
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

func (c *Carrier) NestedPropose(P SuperBlockSummary) error {

	err := c.node.GetEncoder().Encode(&P)
	if err != nil {
		return err
	}

	log.Info().Msgf("proposed SuperBlock %d to %s", c.sbsCounter, c.node.GetAddress())
	return nil
}

func (c *Carrier) GetStringSK() string {
	return c.config.Keys.Sk
}

func (c *Carrier) GetStringPK() string {
	return c.config.Keys.Pk
}

func (c *Carrier) GetKyberSK() kyber.Scalar {
	skk, err := util.DecodeStringToBdnSK(c.config.Keys.Sk)
	if err != nil {
		panic("unable to decode SK")
	}

	return skk
}

func (c *Carrier) GetKyberPK() kyber.Point {
	pkk, err := util.DecodeStringToBdnPK(c.config.Keys.Pk)
	if err != nil {
		panic("unable to decode SK")
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
		return fmt.Errorf("failed to decode h: %s", err.Error())
	}
	sb, err := hex.DecodeString(s.S)
	if err != nil {
		return fmt.Errorf("failed to decode s: %s", err.Error())
	}
	err = bdn.Verify(c.suite, pk, hb, sb)
	return err
}

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

func decide(D map[string][][]byte) {

	// Process decided values
	for h, _ := range D {
		log.Info().Msgf("committed %s", h)
	}
}

func (c *Carrier) getClientToCarrierAddress() string {
	return c.config.Addresses.Client
}

func (c *Carrier) getCarrierToCarrierAddress() string {
	return c.config.Addresses.Carrier
}

func (c *Carrier) getDecisionAddress() string {
	return c.config.Addresses.Decision
}

func (c *Carrier) GetID() string {
	return c.config.ID
}

func (c *Carrier) startListener(address string) (*net.TCPListener, error) {
	resolvedAddress, err := util.ResolveTCPAddr(address)
	resolvedAddress.IP = net.ParseIP("0.0.0.0") // We can only host on localhost
	if err != nil {
		return nil, err
	}
	listener, err := util.ListenTCP(resolvedAddress)

	return listener, err
}

func (c *Carrier) forwardMode() bool {
	return c.config.Settings.ForwardMode == 1
}

func (c *Carrier) getTsxSize() int {
	return c.config.Settings.TsxSize
}

func (c *Carrier) getMempoolThreshold() int {
	return c.config.Settings.InitThreshold
}

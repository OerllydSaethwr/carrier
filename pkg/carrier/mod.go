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
	"net"
	"sync"
	"time"
)

type Carrier struct {
	conf Config

	clientListener   *net.TCPListener
	carrierListener  *net.TCPListener
	decisionListener *net.TCPListener

	client2carrierAddr  string
	carrier2carrierAddr string
	frontAddr           string
	decisionAddr        string

	nodeConn *net.TCPConn

	carriers map[string]string

	// We are relying on the pointers to addresses being equal here. This means the addresses have to originate from
	// the same source, which for now is only in the NewCarrier function. Keep in mind if you want to compute
	// addresses, this will break
	carrierConns map[string]*net.TCPConn

	mempool chan []byte

	// Registry of message handlers. Argument must be one of the enum types
	messageHandlers map[message.Type]func(message.Message) error

	valueStore             map[string][][]byte
	signatureStore         map[string][]util.Signature
	superBlockSummary      []SuperBlockSummaryItem
	superBlockSummaryStore map[string]struct{}
	acceptedHashStore      map[string][][]byte

	suite   *pairing.SuiteBn256
	keypair *util.KeyPairString

	locks Locks

	f int
	n int

	wg *sync.WaitGroup

	secret string
	quit   chan bool
}

type Locks struct {
	CarrierConns      *sync.RWMutex
	ValueStore        *sync.RWMutex
	SignatureStore    *sync.RWMutex
	SuperBlockSummary *sync.RWMutex
}

type Config struct {
	carrierConnRetryDelay time.Duration
	carrierConnMaxRetry   uint
}

type SuperBlockSummaryItem struct {
	ID string           `json:"ID"`
	H  string           `json:"h"`
	S  []util.Signature `json:"s"`
}

type SuperBlockSummary []SuperBlockSummaryItem

func NewCarrier(clientToCarrierAddr, carrierToCarrierAddr, frontAddr, decisionAddr string, carriers map[string]string, keypair *util.KeyPairString) *Carrier {
	conf := Config{
		carrierConnRetryDelay: util.CarrierConnRetryDelay,
		carrierConnMaxRetry:   util.CarrierConnMaxRetry,
	}

	c := &Carrier{}
	c.conf = conf
	c.carrierConns = map[string]*net.TCPConn{}
	c.mempool = make(chan []byte, 100)
	c.quit = make(chan bool, 1)

	c.suite = pairing.NewSuiteBn256()
	c.keypair = keypair

	c.messageHandlers = map[message.Type]func(message.Message) error{}
	c.messageHandlers[message.Init] = c.handleInitMessage
	c.messageHandlers[message.Echo] = c.handleEchoMessage
	c.messageHandlers[message.Request] = c.handleRequestMessage
	c.messageHandlers[message.Resolve] = c.handleResolveMessage

	c.locks = Locks{
		CarrierConns:      &sync.RWMutex{},
		ValueStore:        &sync.RWMutex{},
		SignatureStore:    &sync.RWMutex{},
		SuperBlockSummary: &sync.RWMutex{},
	}

	c.valueStore = map[string][][]byte{}
	c.signatureStore = map[string][]util.Signature{}
	c.superBlockSummary = make([]SuperBlockSummaryItem, 0)
	c.acceptedHashStore = map[string][][]byte{}

	c.client2carrierAddr = clientToCarrierAddr
	c.carrier2carrierAddr = carrierToCarrierAddr
	c.frontAddr = frontAddr
	c.decisionAddr = decisionAddr
	c.carriers = carriers

	c.f = (len(carriers) - 1) / 3
	c.n = len(carriers)

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
	for _, strAddr := range c.carriers {
		_, err := util.ResolveTCPAddr(strAddr)
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

	// Connect to node
	front, err := util.ResolveTCPAddr(c.frontAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
		c.wg.Done()
		return c.wg
	}
	c.nodeConn, err = util.DialTCP(front)
	if err != nil {
		log.Error().Msgf("Failed to connect to node %s", err.Error())
		// We will retry later
	}

	// Listen to nested SMR decisions
	decision, err := util.ResolveTCPAddr(c.decisionAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
		c.wg.Done()
		return c.wg
	}
	c.decisionListener, err = util.ListenTCP(decision)
	if err != nil {
		log.Error().Msgf(err.Error())
		c.wg.Done()
		return c.wg
	}
	log.Info().Msgf("Start listening to nested SMR decisions on %s", c.decisionAddr)
	go c.handleIncomingConnections(c.decisionListener, c.decodeNestedSMRDecisions)

	// Start client listener
	client, err := util.ResolveTCPAddr(c.client2carrierAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
		c.wg.Done()
		return c.wg
	}
	c.clientListener, err = util.ListenTCP(client)
	if err != nil {
		log.Error().Msgf(err.Error())
		c.wg.Done()
		return c.wg
	}
	log.Info().Msgf("Start listening to client on %s", c.client2carrierAddr)
	go c.handleIncomingConnections(c.clientListener, c.handleClientConn)

	// Start carrier listener
	carrier, err := util.ResolveTCPAddr(c.carrier2carrierAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
		c.wg.Done()
		return c.wg
	}
	c.carrierListener, err = util.ListenTCP(carrier)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	log.Info().Msgf("Start listening to carriers on %s", c.carrier2carrierAddr)
	go c.handleIncomingConnections(c.carrierListener, c.handleCarrierConn)

	// Set up connections to other carriers
	for _, address := range c.carriers {
		go c.setupCarrierConnection(address) //TODO goroutine
	}

	return c.wg
}

func (c *Carrier) Stop() {
	log.Trace().Msgf("Stop Carrier")
	err := c.clientListener.Close()
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	c.quit <- true
	c.wg.Done()
}

func (c *Carrier) GetAddress() string {
	return c.carrier2carrierAddr
}

// TODO make this a proper retrying function and gracefully fail if cannot be established
func (c *Carrier) retryNodeConnection() {
	var err error
	front, err := util.ResolveTCPAddr(c.frontAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}

	c.nodeConn, err = util.DialTCP(front)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	log.Info().Msgf("Connect to node %s", c.frontAddr)
}

func (c *Carrier) NestedPropose(P SuperBlockSummary) error {
	if c.nodeConn == nil {
		c.retryNodeConnection()
	}

	encoder := json.NewEncoder(c.nodeConn)
	err := encoder.Encode(P)
	if err != nil {
		return err
	}
	//
	//encoder := gob.NewEncoder(c.nodeConn)
	//err := encoder.Encode(P)
	//if err != nil {
	//	return err
	//}

	log.Info().Msgf("Proposed unknown bytes of data to %s", c.frontAddr)
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
	pk, err := c.VerifyPK(s.Pk)
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

func (c *Carrier) VerifyPK(pk string) (kyber.Point, error) {
	_, ok := c.carriers[pk]
	if !ok {
		return nil, fmt.Errorf("unrecognised sender")
	}

	pkk, err := util.DecodeStringToBdnPK(pk)
	if err != nil {
		return nil, err
	}

	return pkk, nil
}

func (c *Carrier) GetSuite() pairing.Suite {
	return c.suite.Suite
}

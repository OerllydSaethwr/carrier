package carrier

import (
	"encoding/json"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/pairing"
	"net"
	"sync"
	"time"
)

type Carrier struct {
	conf Config

	clientListener  *net.TCPListener
	carrierListener *net.TCPListener

	client2carrierAddr  string
	carrier2carrierAddr string
	frontAddr           string

	nodeConn *net.TCPConn

	carriers map[string]kyber.Point

	// We are relying on the pointers to addresses being equal here. This means the addresses have to originate from
	// the same source, which for now is only in the NewCarrier function. Keep in mind if you want to compute
	// addresses, this will break
	carrierConns map[string]*net.TCPConn

	mempool chan []byte

	// Registry of message handlers. Argument must be one of the enum types
	messageHandlers map[message.Type]func(message.Message) error

	valueStore        map[string][][]byte
	signatureStore    map[string][][]byte
	superBlockSummary []*SuperBlockSummary

	suite   *pairing.SuiteBn256
	keypair *util.KeyPair

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

type SuperBlockSummary struct {
	H          []byte   `json:"h"`
	Signatures [][]byte `json:"signatures"`
}

func NewCarrier(clientToCarrierAddr, carrierToCarrierAddr, frontAddr string, carriers map[string]kyber.Point, keypair *util.KeyPair) *Carrier {
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
	c.signatureStore = map[string][][]byte{}
	c.superBlockSummary = make([]*SuperBlockSummary, 0)

	c.client2carrierAddr = clientToCarrierAddr
	c.carrier2carrierAddr = carrierToCarrierAddr
	c.frontAddr = frontAddr
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
	for strAddr, _ := range carriers {
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
	for address := range c.carriers {
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

func (c *Carrier) NestedPropose(P []*SuperBlockSummary) error {
	if c.nodeConn == nil {
		c.retryNodeConnection()
	}

	data, err := json.Marshal(P)
	if err != nil {
		return err
	}

	_, err = c.nodeConn.Write(data)
	if err != nil {
		return err
	}

	log.Info().Msgf("Proposed %d bytes of data to %s", len(data), c.frontAddr)
	return nil
}

/* TODO
1. add command to generate key files
2. batch transactions in local mempool
- add stuff
- remove stuff when threshold is hit
3. process that watches mempool and invokes functions after threshold

*/

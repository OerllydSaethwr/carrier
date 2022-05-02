package carrier

import (
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
	"net"
	"sync"
	"time"
)

type Carrier struct {
	conf Config

	clientListener  *net.TCPListener
	carrierListener *net.TCPListener

	client2carrierAddr  *net.TCPAddr
	carrier2carrierAddr *net.TCPAddr
	frontAddr           *net.TCPAddr

	nodeConn *net.TCPConn

	carrierAddrs []*net.TCPAddr

	// We are relying on the pointers to addresses being equal here. This means the addresses have to originate from
	// the same source, which for now is only in the NewCarrier function. Keep in mind if you want to compute
	// addresses, this will break
	carrierConns map[*net.TCPAddr]*net.TCPConn

	mempool chan []byte

	// Registry of message handlers. Argument must be one of the enum types
	messageHandlers map[MessageType]func(any) error

	wg *sync.WaitGroup

	secret string
	quit   chan bool
}

type Config struct {
	carrierConnRetryDelay time.Duration
	carrierConnMaxRetry   uint
}

// CarrierAddrs helps with importing addresses of other carriers
type CarrierAddrs struct {
	CarrierAddrs []string `json:"carriers"`
}

func NewCarrier(clientToCarrierAddr, carrierToCarrierAddr, frontAddr string, carrierAddrs []string) *Carrier {
	conf := Config{
		carrierConnRetryDelay: util.CarrierConnRetryDelay,
		carrierConnMaxRetry:   util.CarrierConnMaxRetry,
	}

	c := &Carrier{}
	c.conf = conf
	c.carrierConns = map[*net.TCPAddr]*net.TCPConn{}
	c.mempool = make(chan []byte, 100)
	c.quit = make(chan bool, 1)

	c.messageHandlers = map[MessageType]func(any) error{}
	c.messageHandlers[Init] = c.handleInitMessage
	c.messageHandlers[Echo] = c.handleEchoMessage
	c.messageHandlers[Request] = c.handleRequestMessage
	c.messageHandlers[Resolve] = c.handleResolveMessage
	//TODO secret

	var err error
	c.client2carrierAddr, err = net.ResolveTCPAddr(util.Network, clientToCarrierAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
		return nil
	}
	c.carrier2carrierAddr, err = net.ResolveTCPAddr(util.Network, carrierToCarrierAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
		return nil
	}
	c.frontAddr, err = net.ResolveTCPAddr(util.Network, frontAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
		return nil
	}

	c.carrierAddrs = make([]*net.TCPAddr, 0)
	for _, strAddr := range carrierAddrs {
		addr, err := net.ResolveTCPAddr(util.Network, strAddr)
		if err != nil {
			log.Error().Msgf(err.Error())
		} else {
			c.carrierAddrs = append(c.carrierAddrs, addr)
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
	c.nodeConn, err = util.DialTCP(c.frontAddr)
	if err != nil {
		log.Error().Msgf("Failed to connect to node %s", err.Error())
		// We will retry later
	}

	// Start client listener
	c.clientListener, err = util.ListenTCP(c.client2carrierAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
		c.wg.Done()
		return c.wg
	}
	log.Info().Msgf("Start listening to client on %s", c.client2carrierAddr.String())
	go c.handleIncomingConnections(c.clientListener, c.handleClientConn)

	// Start carrier listener
	c.carrierListener, err = util.ListenTCP(c.carrier2carrierAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	log.Info().Msgf("Start listening to carriers on %s", c.client2carrierAddr.String())
	go c.handleIncomingConnections(c.carrierListener, c.handleCarrierConn)

	// Set up connections to other carriers
	for _, carrierAddr := range c.carrierAddrs {
		go c.setupCarrierConnection(carrierAddr)
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

func (c *Carrier) retryNodeConnection() {
	var err error
	c.nodeConn, err = util.DialTCP(c.frontAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	log.Info().Msgf("Connect to node %s", c.frontAddr)
}

/* TODO
1. add command to generate key files
2. batch transactions in local mempool
- add stuff
- remove stuff when threshold is hit
3. process that watches mempool and invokes functions after threshold

*/

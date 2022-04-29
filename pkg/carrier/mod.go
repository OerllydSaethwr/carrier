package carrier

import (
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
	"io"
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
	carrierConns []*net.TCPConn

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

func NewCarrier(wg *sync.WaitGroup, clientToCarrierAddr, carrierToCarrierAddr, frontAddr string, carrierAddrs []string) *Carrier {
	conf := Config{
		carrierConnRetryDelay: util.CarrierConnRetryDelay,
		carrierConnMaxRetry:   util.CarrierConnMaxRetry,
	}

	c := &Carrier{}
	c.conf = conf
	c.carrierConns = make([]*net.TCPConn, 0)
	c.wg = wg
	c.quit = make(chan bool, 1)
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
func (c *Carrier) Start() {
	c.wg.Add(1)

	var err error

	// Connect to node
	c.nodeConn, err = util.DialTCP(c.frontAddr)
	if err != nil {
		log.Error().Msgf("Failed to connect to node %s", err.Error())
		// We will retry later
	}

	// Start client listener
	c.clientListener, err = net.ListenTCP(util.Network, c.client2carrierAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}
	log.Info().Msgf("Start listening to client on %s", c.client2carrierAddr.String())
	go c.startProcessor(c.clientListener, c.processClientConn)

	// Start carrier listener
	c.carrierListener, err = net.ListenTCP(util.Network, c.carrier2carrierAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	log.Info().Msgf("Start listening to carriers on %s", c.client2carrierAddr.String())
	go c.startProcessor(c.carrierListener, c.processCarrierConn)

	// Set up connections to other carriers
	for _, carrierAddr := range c.carrierAddrs {
		go c.setupCarrierConnection(carrierAddr)
	}
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

func (c *Carrier) startProcessor(l *net.TCPListener, process func(conn net.Conn)) {
	for {
		select {
		case <-c.quit:
			return
		default:
			conn, err := l.AcceptTCP()
			if err != nil {
				log.Error().Msgf(err.Error())
				return
			}
			go process(conn)
		}
	}
}

func (c *Carrier) processClientConn(conn net.Conn) {
	// If we didn't manage to connect to the node before, try one last time
	if c.nodeConn == nil {
		c.retryNodeConnection()
	}
	// Make a buffer to hold incoming data.
	buf := make([]byte, util.Tsx_size) //TODO make this configurable
	// Read the incoming connection into the buffer.

	var err error
	var n int
	for n, err = io.ReadAtLeast(conn, buf, util.Tsx_size); err == nil; n, err = io.ReadAtLeast(conn, buf, util.Tsx_size) {
		var err2 error
		log.Info().Msgf("Read %d bytes from %s", n, conn.RemoteAddr())
		if c.nodeConn != nil {
			_, err2 = c.nodeConn.Write(buf)
			log.Info().Msgf("Forwarded %d bytes to %s", n, c.nodeConn.RemoteAddr())
		}
		if err2 != nil {
			log.Error().Msgf(err2.Error())
		}
	}
	conn.Close()
	log.Info().Msgf(err.Error())
	log.Info().Msgf("Close client connection %s", conn.RemoteAddr())
}

func (c *Carrier) processCarrierConn(conn net.Conn) {
	return
}

func (c *Carrier) retryNodeConnection() {
	var err error
	c.nodeConn, err = util.DialTCP(c.frontAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	log.Info().Msgf("Connect to node %s", c.frontAddr)
}

func (c *Carrier) setupCarrierConnection(carrierAddr *net.TCPAddr) {
	// If carrierConnMaxRetry is 0, we keep retrying indefinitely
	for i := uint(0); c.conf.carrierConnMaxRetry == 0 || i < c.conf.carrierConnMaxRetry; i++ {
		conn, err := util.DialTCP(carrierAddr)
		if err == nil {
			c.carrierConns = append(c.carrierConns, conn)
			log.Info().Msgf("Connect to carrier %s | attempt %d/%d", carrierAddr.String(), i+1, c.conf.carrierConnMaxRetry)
			return
		} else {
			log.Info().Msgf("Failed to connect to carrier %s | attempt %d/%d", carrierAddr.String(), i+1, c.conf.carrierConnMaxRetry)
			time.Sleep(c.conf.carrierConnRetryDelay)
		}
	}
}

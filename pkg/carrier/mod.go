package carrier

import (
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
	"io"
	"net"
	"sync"
)

type Carrier struct {
	clientListener  *net.TCPListener
	carrierListener *net.TCPListener

	client2carrierAddr  *net.TCPAddr
	carrier2carrierAddr *net.TCPAddr
	frontAddr           *net.TCPAddr

	nodeConn *net.TCPConn

	wg *sync.WaitGroup

	secret string
	quit   chan bool
}

func NewCarrier(wg *sync.WaitGroup, clientToCarrierAddr, carrierToCarrierAddr, frontAddr string) *Carrier {
	c := &Carrier{}
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

	c.nodeConn, err = net.DialTCP(util.Network, nil, c.frontAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
		log.Error().Msgf("Failed to connect to node")
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
	c.clientListener, err = net.ListenTCP("tcp", c.client2carrierAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}
	log.Info().Msgf("Start listening on %s", c.client2carrierAddr.String())
	go c.startProcessor(c.clientListener, c.processClientConn)
}

func (c *Carrier) Stop() {
	log.Trace().Msgf("Stop Carrier")
	c.clientListener.Close()
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
	buf := make([]byte, 9) //TODO make this configurable
	// Read the incoming connection into the buffer.
	reader := io.LimitReader(conn, int64(len(buf)))

	for _, err := reader.Read(buf); err == nil; _, err = reader.Read(buf) {
		var err2 error
		if c.nodeConn != nil {
			_, err2 = c.nodeConn.Write(buf)
		}
		if err2 != nil {
			log.Error().Msgf(err2.Error())
		}
	}
	conn.Close()
	log.Info().Msgf("Close client connection %s", conn.RemoteAddr())
}

func (c *Carrier) processCarrierConn(conn net.Conn) {
	return
}

func (c *Carrier) retryNodeConnection() {
	var err error
	c.nodeConn, err = net.DialTCP(util.Network, nil, c.frontAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
}

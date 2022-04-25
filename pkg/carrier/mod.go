package carrier

import (
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
	c.client2carrierAddr, err = net.ResolveTCPAddr("tcp4", clientToCarrierAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
		return nil
	}
	c.carrier2carrierAddr, err = net.ResolveTCPAddr("tcp4", carrierToCarrierAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
		return nil
	}
	c.frontAddr, err = net.ResolveTCPAddr("tcp4", carrierToCarrierAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
		return nil
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
	log.Info().Msgf("Listening on %s", c.client2carrierAddr.String())
	go c.startProcessor(c.clientListener, c.processClientConn)
}

func (c *Carrier) Stop() {
	log.Trace().Msgf("Stopping Carrier")
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
	// Make a buffer to hold incoming data.
	buf := make([]byte, 5)
	// Read the incoming connection into the buffer.
	reader := io.LimitReader(conn, int64(len(buf)))
	_, err := reader.Read(buf)
	if err != nil {
		log.Trace().Msgf(err.Error())
		return
	}
	//if err != nil {
	//	log.Trace().Msgf("Error reading: %s", err.Error())
	//}
	// Close the connection when you're done with it.
	conn.Close()
	log.Info().Msgf("Reading from connection %s", conn)
}

func (c *Carrier) processCarrierConn(conn net.Conn) {
	return
}

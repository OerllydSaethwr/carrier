package carrier

import (
	"fmt"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
	"net"
	"sync"
)

type Carrier struct {
	clientListener  Listener
	carrierListener Listener
	processClient   chan net.Conn
	processCarrier  chan net.Conn

	carrierAddr string
	carrierPort int
	frontAddr   string
	frontPort   int

	wg *sync.WaitGroup

	secret string
	quit   chan bool
}

func NewCarrier(wg *sync.WaitGroup, clientToCarrierAddr, carrierToCarrierAddr, frontAddr string) *Carrier {
	hostcl2ca, portcl2ca, err := util.SplitHostPort(clientToCarrierAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	hostca2ca, portca2ca, err := util.SplitHostPort(carrierToCarrierAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	hostca2f, portca2f, err := util.SplitHostPort(frontAddr)
	if err != nil {
		log.Error().Msgf(err.Error())
	}

	fmt.Println(hostca2f, portca2f) //TODO remove

	processClient := make(chan net.Conn)
	processCarrier := make(chan net.Conn)

	p := Carrier{
		clientListener:  NewTCPListener(processClient, hostcl2ca, portcl2ca),
		carrierListener: NewTCPListener(processCarrier, hostca2ca, portca2ca),
		processClient:   processClient,
		processCarrier:  processCarrier,

		wg: wg,

		secret: xid.New().String(), //TODO pass in
		quit:   make(chan bool, 1),
	}

	return &p
}

/*	Start listening to client requests
	Forward client requests
	We are not waiting for listeners to stop but I think it's fine
*/
func (c *Carrier) Start() {
	c.wg.Add(1)
	go c.clientListener.Listen()
	go c.startProcessor(c.clientListener, c.processClientConn)
}

func (c *Carrier) Stop() {
	log.Trace().Msgf("Stopping Carrier")

	c.clientListener.Stop()
	c.quit <- true
	c.wg.Done()
}

func (c *Carrier) startProcessor(listener Listener, process func(conn net.Conn)) {
	for {
		select {
		case <-c.quit:
			return
		default:
			process(listener.GetConn())
		}
	}
}

func (c *Carrier) processClientConn(conn net.Conn) {

	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		log.Trace().Msgf("Error reading:", err.Error())
	}
	// Close the connection when you're done with it.
	conn.Close()

}

func (c *Carrier) processCarrierConn(conn net.Conn) {
	return
}

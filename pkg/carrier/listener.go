package carrier

import (
	. "github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
	"net"
	"os"
	"strconv"
)

type Listener interface {
	Listen()
	Stop()
	GetConn() net.Conn

	// utils
	GetName() string
}

// TCPListener Generic wrapper around net listener
type TCPListener struct {
	quit    chan bool
	process chan net.Conn

	host string
	port int

	// utils
	name string
}

func NewTCPListener(process chan net.Conn, host string, port int) *TCPListener {
	return &TCPListener{
		quit:    make(chan bool, 1),
		process: process,
		host:    host,
		port:    port,
		name:    xid.New().String(),
	}
}

func (tcpl *TCPListener) Stop() {
	log.Info().Msgf("Stopping Listener " + tcpl.GetName())
	tcpl.quit <- true
}

func (tcpl *TCPListener) Listen() {
	l, err := net.Listen(TYPE, tcpl.host+":"+strconv.Itoa(tcpl.port))
	if err != nil {
		log.Error().Msgf(err.Error())
		os.Exit(1) //TODO maybe this is too harsh
	}

	defer l.Close()
	log.Info().Msgf("Listening on " + tcpl.host + ":" + strconv.Itoa(tcpl.port))

	for {
		select {
		case <-tcpl.quit:
			return
		default:
			conn, err := l.Accept()
			if err != nil {
				log.Error().Msgf(err.Error())
				os.Exit(1)
			}
			tcpl.process <- conn
		}
	}
}

// HandleRequest Handles incoming requests.
//func (tcpl *TCPListener) HandleRequest(conn net.Conn) {
//	// Make a buffer to hold incoming data.
//	buf := make([]byte, 1024)
//	// Read the incoming connection into the buffer.
//	_, err := conn.Read(buf)
//	if err != nil {
//		log.Trace().Msgf("Error reading:", err.Error())
//	}
//	// Send a response back to person contacting us.
//	conn.Write([]byte("Message received."))
//	// Close the connection when you're done with it.
//	conn.Close()
//}

/* Utils */

func (tcpl TCPListener) GetName() string {
	return tcpl.name
}

func (tcpl TCPListener) GetConn() net.Conn {
	return <-tcpl.process
}

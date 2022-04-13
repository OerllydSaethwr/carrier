package carrier

import (
	"github.com/rs/zerolog/log"
	. "gitlab.epfl.ch/valaczka/carrier/pkg/util"
	"net"
	"os"
	"strconv"
)

type Listener interface {
	Start()
	Stop()

	// utils
	GetName() string
}

// TCPListener Wrapper around net listener
type TCPListener struct {
	quit chan bool

	// utils
	name string
	port int
}

func (tcpl *TCPListener) Start() {
	go tcpl.Listen()
}

func (tcpl *TCPListener) Stop() {
	log.Info().Msgf("Stopping Listener " + tcpl.GetName())
	tcpl.quit <- true
}

func (tcpl *TCPListener) Listen() {
	for {
		select {
		case <-tcpl.quit:
			return
		default:
			l, err := net.Listen(TYPE, HOST+":"+strconv.Itoa(tcpl.port))
			if err != nil {
				log.Info().Msgf(err.Error())
				os.Exit(1)
			}
			// Close the listener when the application closes.
			defer l.Close()
			log.Info().Msgf("Listening on " + HOST + ":" + strconv.Itoa(tcpl.port))
			for {
				// Listen for an incoming connection.
				conn, err := l.Accept()
				if err != nil {
					log.Info().Msgf(err.Error())
					os.Exit(1)
				}
				// Handle connections in a new goroutine.
				go tcpl.HandleRequest(conn)
			}
		}
	}
}

// HandleRequest Handles incoming requests.
func (tcpl *TCPListener) HandleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		log.Trace().Msgf("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}

/* Utils */

func (tcpl TCPListener) GetName() string {
	return tcpl.name
}

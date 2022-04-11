package impl

import (
	"fmt"
	"net"
	"os"

	"github.com/rs/xid"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8000"
	CONN_TYPE = "tcp"
)

type addr string

type Proxy interface {
	Start()
	Stop()
}

type proxy struct {
	Proxy

	quit          chan bool
	targetMempool addr

	secret string
}

//  TODO: just forward everything to an address
func NewProxy() Proxy {
	p := proxy{
		quit:          make(chan bool, 1),
		secret:        xid.New().String(), //TODO
		targetMempool: "0.0.0.0",
	}

	return &p
}

/*	Start listening to client requests
	Forward client requests
*/
func (p *proxy) Start() {
	go p.listen()
}

func (p *proxy) Stop() {
	fmt.Println("Stopping proxy")
	p.quit <- true
}

func (p *proxy) listen() {
	for {
		select {
		case <-p.quit:
			return
		default:
			l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
			if err != nil {
				fmt.Println("Error listening:", err.Error())
				os.Exit(1)
			}
			// Close the listener when the application closes.
			defer l.Close()
			fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
			for {
				// Listen for an incoming connection.
				conn, err := l.Accept()
				if err != nil {
					fmt.Println("Error accepting: ", err.Error())
					os.Exit(1)
				}
				// Handle connections in a new goroutine.
				go p.handleRequest(conn)
			}
		}
	}
}

// Handles incoming requests.
func (p *proxy) handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}

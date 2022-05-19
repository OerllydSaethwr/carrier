package main

import (
	"github.com/rs/zerolog/log"
	"net"
	"os"
)

const (
	//	CONN_HOST = "localhost"
	//	CONN_PORT = "9000"
	CONN_TYPE = "tcp"
)

func main() {
	CONN_HOST := os.Args[1]
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST)
	if err != nil {
		log.Error().Msgf("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	log.Info().Msgf("Listening on " + CONN_HOST)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			log.Error().Msgf("Error accepting: ", err.Error())
			os.Exit(1)
		}
		log.Info().Msgf("Accepted")
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {

}

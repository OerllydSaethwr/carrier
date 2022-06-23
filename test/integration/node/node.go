package main

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier/codec"
	"github.com/rs/zerolog/log"
	"net"
	"os"
	"time"
)

const (
	//	CONN_HOST = "localhost"
	//	CONN_PORT = "9000"
	CONN_TYPE = "tcp"
)

func main() {
	CONN_HOST := os.Args[1]
	decision := os.Args[2]
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
		go handleRequest(conn, decision)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn, da string) {
	decoder := codec.NewBinaryDecoder(conn)

	decision := setupDecisionConn(da)
	encoder := codec.NewBinaryEncoder(decision)
	for {

		var buf []byte
		err := decoder.Decode(&buf)
		if err != nil {
			log.Error().Msgf(err.Error())
		}

		time.Sleep(time.Millisecond * 100)
		err = encoder.Encode(buf)
		if err != nil {
			log.Error().Msgf(err.Error())
		}
		log.Info().Msgf("Sent %s to %s", buf, da)

		// Legacy code, could be useful
		// If forward mode is on, don't try to decode the superblock summary
		//if util.ForwardMode {
		//	buf := make([]byte, util.TsxSize)
		//	err := decoder.Decode(&buf)
		//	if err != nil {
		//		log.Error().Msgf(err.Error())
		//	}
		//	time.Sleep(time.Millisecond*100)
		//	err = encoder.Encode(buf)
		//	if err != nil {
		//		log.Error().Msgf(err.Error())
		//	}
		//	continue
		//}
		//// Otherwise, decode it
		//var superBlockSummary carrier.SuperBlockSummary
		//err := decoder.Decode(&superBlockSummary)
		//if err != nil {
		//	log.Error().Msgf(err.Error())
		//}
		//log.Info().Msgf("Read %s from %s", superBlockSummary, conn.RemoteAddr())

		//time.Sleep(time.Millisecond * 100)
		//err = encoder.Encode(&superBlockSummary)
		//if err != nil {
		//	log.Error().Msgf(err.Error())
		//}
		//log.Info().Msgf("Sent %s to %s", superBlockSummary, da)
	}
}

func setupDecisionConn(a string) *net.TCPConn {
	ar, err := net.ResolveTCPAddr("tcp", a)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	var conn *net.TCPConn
	for {
		conn, err = net.DialTCP("tcp", nil, ar)
		if err != nil {
			log.Error().Msgf(err.Error())
			time.Sleep(time.Second)
		} else {
			break
		}
	}

	return conn
}

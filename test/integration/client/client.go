package main

import (
	"github.com/rs/zerolog/log"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	var err error
	var conn *net.TCPConn

	transaction := make([]byte, 9)
	servAddr := os.Args[1]
	rate := 5

	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		log.Error().Msgf("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	for {
		conn, err = net.DialTCP("tcp", nil, tcpAddr)
		if err == nil {
			break
		}
		log.Info().Msgf("Attempting to connect: %s", tcpAddr)
		time.Sleep(time.Second)
	}
	log.Info().Msgf("Connected")

	for {
		rand.Read(transaction)
		_, err = conn.Write(transaction)
		if err != nil {
			log.Error().Msgf("Write to server failed:", err.Error())
			os.Exit(1)
		}
		log.Info().Msgf("Send %d bytes to %s", len(transaction), servAddr)
		time.Sleep(time.Duration(int64(time.Second) / int64(rate)))
	}

	conn.Close()
}

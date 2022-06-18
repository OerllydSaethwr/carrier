package main

import (
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	var err error
	var conn *net.TCPConn

	transaction := make([]byte, util.TsxSize)
	servAddr := os.Args[1]
	var counter uint = 0
	rate := 200000

	zerolog.SetGlobalLevel(zerolog.Disabled)

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

	go func() {
		for {
			time.Sleep(time.Second)
			println(counter)
		}
	}()

	for {
		rand.Read(transaction)
		_, err = conn.Write(util.Frame(transaction))
		if err != nil {
			log.Error().Msgf("Write to server failed:", err.Error())
			os.Exit(1)
		}
		log.Info().Msgf("Send %d bytes to %s", len(transaction), servAddr)
		counter++
		time.Sleep(time.Duration(int64(time.Second) / int64(rate)))
	}

	conn.Close()
}

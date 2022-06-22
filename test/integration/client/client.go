package main

import (
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

// Client takes 3 arguments: server-address, tsx-size and rate
func main() {
	var err error
	var conn *net.TCPConn

	servAddr := os.Args[1]

	tsxSize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}
	rate, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}

	transaction := make([]byte, tsxSize)
	var counter uint = 0

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

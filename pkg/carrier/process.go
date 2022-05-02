package carrier

import (
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
	"net"
	"time"
)

/* Functions in this file are typically invoked as their own goroutines and loop either indefinitely or until their goal is achieved */

func (c *Carrier) handleIncomingConnections(l *net.TCPListener, handler func(conn net.Conn)) {
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
			go handler(conn)
		}
	}
}

func (c *Carrier) setupCarrierConnection(carrierAddr *net.TCPAddr) {
	// If carrierConnMaxRetry is 0, we keep retrying indefinitely
	for i := uint(0); c.conf.carrierConnMaxRetry == 0 || i < c.conf.carrierConnMaxRetry; i++ {
		conn, err := util.DialTCP(carrierAddr)
		if err == nil {
			c.carrierConns[carrierAddr] = conn
			log.Info().Msgf("Connect to carrier %s | attempt %d/%d", carrierAddr.String(), i+1, c.conf.carrierConnMaxRetry)
			return
		} else {
			log.Info().Msgf("Failed to connect to carrier %s | attempt %d/%d", carrierAddr.String(), i+1, c.conf.carrierConnMaxRetry)
			time.Sleep(c.conf.carrierConnRetryDelay)
		}
	}
}

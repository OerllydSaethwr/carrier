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

func connect(n Remote, retryDelay time.Duration, maxRetry uint) {
	// If xxxConnMaxRetry is 0, we keep retrying indefinitely
	address, err := util.ResolveTCPAddr(n.GetAddress())
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	for i := uint(0); maxRetry == 0 || i < maxRetry; i++ {
		conn, err := util.DialTCP(address)
		if err != nil {
			log.Info().Msgf("failed to connect to %s %s: %s | attempt %d/%d", n.GetType(), address.String(), err.Error(), i+1, maxRetry)
			time.Sleep(retryDelay)
		} else {
			log.Info().Msgf("connect to %s %s | attempt %d/%d", n.GetType(), address.String(), i+1, maxRetry)
			n.SetConnAndEncoderAndSignalAlive(conn)

			return
		}
	}
}

func (c *Carrier) checkAcceptedHashStoreAndDecide() {
	c.locks.AcceptedHashStore.Lock()
	defer c.locks.AcceptedHashStore.Unlock()

	for _, v := range c.stores.acceptedHashStore {
		if v == nil {
			return
		}
	}

	c.decide(c.stores.acceptedHashStore)
}

func (c *Carrier) logger() {
	for {
		time.Sleep(time.Second)
		println(c.counter)
	}
}

func (c *Carrier) bufferMaker() {
	for {
		c.bufferDispenser <- make([]byte, util.TsxSize)
	}
}

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

func connect(n Remote, retryDelay time.Duration, maxRetry int) {
	// If xxxConnMaxRetry is 0, we keep retrying indefinitely
	address, err := util.ResolveTCPAddr(n.GetAddress())
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	for i := 0; maxRetry == 0 || i < maxRetry; i++ {
		conn, err := util.DialTCP(address)
		if err != nil {
			log.Info().Msgf("FAIL - connect to %s %s: %s | attempt %d/%d", n.GetType(), address.String(), err.Error(), i+1, maxRetry)
			time.Sleep(retryDelay)
		} else {
			log.Info().Msgf("SUCCESS - connect to %s %s | attempt %d/%d", n.GetType(), address.String(), i+1, maxRetry)
			n.SetConnAndEncoderAndSignalAlive(conn)

			return
		}
	}
}

func (c *Carrier) checkAcceptedHashStoreAndDecide() {
	c.locks.AcceptedHashStore.Lock()
	defer c.locks.AcceptedHashStore.Unlock()

	decidedHashes := make([]string, 0)
	for h, v := range c.stores.acceptedHashStore {

		// If we don't have all hashes, abort
		if v == nil {
			return
		}
		decidedHashes = append(decidedHashes, h)
	}

	// Decide
	defer c.locks.DecisionLock.Unlock()
	D := c.stores.acceptedHashStore
	c.stores.acceptedHashStore = map[string][][]byte{}
	for _, hash := range decidedHashes {
		c.stores.decidedHashStore[hash] = struct{}{}
	}

	log.Info().Msgf("total hashes decided: %d", len(c.stores.decidedHashStore))

	decide(D)
}

func (c *Carrier) logger() {
	for {
		time.Sleep(time.Second)
		println(len(c.stores.valueStore))
	}
}

func (c *Carrier) launchWorkerPool(poolSize int, task func()) {
	for i := 0; i < poolSize; i++ {
		go task()
	}
}

func (c *Carrier) broadcastWorker() {
	for {
		c.executeBroadcast(<-c.broadcastDispenser)
	}
}

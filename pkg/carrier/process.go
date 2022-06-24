package carrier

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier/remote"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
	"net"
	"time"
)

/* Functions in this file are typically invoked as their own goroutines and loop either indefinitely or until their goal is achieved */

func (c *Carrier) HandleIncomingConnections(l *net.TCPListener, handler func(conn net.Conn)) {
	for {
		select {
		case <-c.Quit:
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

func connect(n remote.Remote, retryDelay time.Duration, maxRetry int) {
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

func (c *Carrier) CheckAcceptedHashStoreAndDecide() {
	c.Locks.AcceptedHashStore.Lock()
	defer c.Locks.AcceptedHashStore.Unlock()

	decidedHashes := make([]string, 0)
	for h, v := range c.Stores.acceptedHashStore {

		// If we don't have all hashes, abort
		if v == nil {
			return
		}
		decidedHashes = append(decidedHashes, h)
	}

	// Decide
	defer c.Locks.DecisionLock.Unlock()
	D := c.Stores.acceptedHashStore
	c.Stores.acceptedHashStore = map[string][][]byte{}
	for _, hash := range decidedHashes {
		c.Stores.decidedHashStore[hash] = struct{}{}
	}

	log.Info().Msgf("total hashes decided: %d", len(c.Stores.decidedHashStore))

	Decide(D)
}

func (c *Carrier) Logger() {
	for {
		time.Sleep(time.Second)
		println(len(c.Stores.valueStore))
	}
}

func (c *Carrier) LaunchWorkerPool(poolSize int, task func()) {
	for i := 0; i < poolSize; i++ {
		go task()
	}
}

func (c *Carrier) BroadcastWorker() {
	for {
		c.ExecuteBroadcast(<-c.BroadcastDispenser)
	}
}

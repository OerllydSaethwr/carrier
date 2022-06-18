// @BETA

package carrier

import (
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"net"
)

// Conn is a new wrapper around net.Conn
// The idea is to implement a reliable connection that will detect errors and attempt to reconnect if the connection breaks down.
type Conn struct {
	conn    net.Conn
	sink    chan []byte
	address string
}

func Connect(address string, sinkBufferSize int) *Conn {
	// Connect

	return nil
}

func (c *Conn) Write(buf []byte) {
	c.sink <- buf
}

func (c *Conn) connect() error {
	address, err := util.ResolveTCPAddr(c.address)
	if err != nil {
		return err
	}

	conn, err := util.DialTCP(address)
	if err != nil {
		return err
	} else {
		c.conn = conn
		return nil
	}
}

func (c *Conn) runSinkConsumer() {
	for {
		if c.conn == nil {
			c.connect()
		} else {
			_, err := c.conn.Write(<-c.sink)
			if err != nil {
				c.conn = nil
			}
		}
	}
}

//func connect(n Remote, retryDelay time.Duration, maxRetry uint) {
//	// If xxxConnMaxRetry is 0, we keep retrying indefinitely
//	address, err := util.ResolveTCPAddr(n.GetAddress())
//	if err != nil {
//		log.Error().Msgf(err.Error())
//	}
//	for i := uint(0); maxRetry == 0 || i < maxRetry; i++ {
//		conn, err := util.DialTCP(address)
//		if err != nil {
//			log.Info().Msgf("FAIL - connect to %s %s: %s | attempt %d/%d", n.GetType(), address.String(), err.Error(), i+1, maxRetry)
//			time.Sleep(retryDelay)
//		} else {
//			log.Info().Msgf("SUCCESS - connect to %s %s | attempt %d/%d", n.GetType(), address.String(), i+1, maxRetry)
//			n.SetConnAndEncoderAndSignalAlive(conn)
//
//			return
//		}
//	}
//}

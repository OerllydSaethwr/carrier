package carrier

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/rs/zerolog/log"
	"net"
	"sync/atomic"
)

/* Functions in this file are typically invoked as their own goroutines and loop while the connection is open */

func (c *Carrier) handleClientConn(conn net.Conn) {
	// Make a buffer to hold incoming data.
	// Read the incoming connection into the buffer.
	decoder := NewBinaryDecoder(conn)
outerLoop:
	for {

		// If we're in ForwardMode, forward messages to node without any processing
		if c.forwardMode() {
			buf := make([]byte, c.getTsxSize())
			err := decoder.Decode(&buf)
			if err != nil {
				log.Error().Msgf(err.Error())
				break outerLoop
			}

			err = c.node.GetEncoder().Encode(buf)
			if err != nil {
				log.Error().Msgf(err.Error())
				break outerLoop
			}
			log.Debug().Msgf("forward tsx to %s", c.node.GetAddress())
			continue
		}

		// Otherwise, do normal processing
		initMessage := message.NewInitMessage(
			make([][]byte, 0),
			c.GetID(),
		)
		for i := 0; i < c.getMempoolThreshold(); i++ {
			buf := make([]byte, c.getTsxSize())

			err := decoder.Decode(&buf)
			//buf := make([]byte, util.TsxSize) //TODO make this configurable
			//_, err := io.ReadAtLeast(conn, buf, util.TsxSize)
			if err != nil {
				log.Error().Msgf(err.Error())
				break outerLoop
			}

			initMessage.V = append(initMessage.V, buf)
		}

		log.Debug().Msgf("V %d", len(c.stores.valueStore))
		log.Debug().Msgf("S %d", len(c.stores.signatureStore))
		log.Debug().Msgf("P %d", len(c.stores.superBlockSummary))
		log.Debug().Msgf("D %d", len(c.stores.acceptedHashStore))

		atomic.AddUint64(&c.counter, 1)

		c.broadcast(initMessage)
	}

	log.Info().Msgf("close client connection %s", conn.RemoteAddr())
	err := conn.Close()
	if err != nil {
		log.Error().Msgf(err.Error())
	}
}

func (c *Carrier) handleCarrierConn(conn net.Conn) {
	decoder := NewBinaryDecoder(conn)
	for {

		// We expect packets framed using util.Frame - they will contain a uint32 (4 bytes) describing the length of the incoming stream
		var m message.Message
		err := decoder.Decode(&m)
		if err != nil {
			log.Error().Msgf(err.Error())
			continue
		}

		log.Debug().Msgf("received %s from %s", m.GetType(), m.GetSenderID())

		// Send to message handler
		err = c.messageHandlers[m.GetType()](m)
		if err != nil {
			log.Error().Msgf("should not get here - garbage messages should be caught during decoding")
			panic("message handler returned error: " + err.Error())
		}

	}
}

func (c *Carrier) decodeNestedSMRDecisions(conn net.Conn) {
	decoder := NewBinaryDecoder(conn)
	for {
		// Only decode byte array if we're in forward mode
		if c.forwardMode() {
			var buf []byte
			err := decoder.Decode(&buf)
			if err != nil {
				log.Error().Msgf(err.Error())
			}
			log.Debug().Msgf("received nested SMR decision from %s", c.node.GetAddress())
			continue
		}

		var N SuperBlockSummary
		err := decoder.Decode(&N)
		if err != nil {
			log.Error().Msgf(err.Error())

			// Ignore garbage messages
			continue
		}
		log.Info().Msgf("received nested SMR decision from %s", c.node.GetAddress())
		c.handleNestedSMRDecision(N)
	}
}

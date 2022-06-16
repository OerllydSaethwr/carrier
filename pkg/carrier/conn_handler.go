package carrier

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
	"io"
	"net"
	"sync/atomic"
)

/* Functions in this file are typically invoked as their own goroutines and loop while the connection is open */

func (c *Carrier) handleClientConn(conn net.Conn) {
	// Make a buffer to hold incoming data.
	// Read the incoming connection into the buffer.
outerLoop:
	for {
		initMessage := message.NewInitMessage(
			make([][]byte, 0),
			c.GetID(),
		)
		for i := 0; i < util.MempoolThreshold; i++ {
			buf := make([]byte, util.TsxSize) //TODO make this configurable
			_, err := io.ReadAtLeast(conn, buf, util.TsxSize)
			if err != nil {
				log.Info().Msgf(err.Error())
				break outerLoop
			}

			initMessage.V = append(initMessage.V, buf)
		}

		log.Debug().Msgf("V %d", len(c.stores.valueStore))
		log.Debug().Msgf("S %d", len(c.stores.signatureStore))
		log.Debug().Msgf("P %d", len(c.stores.superBlockSummary))
		log.Debug().Msgf("D %d", len(c.stores.acceptedHashStore))

		atomic.AddUint64(&c.counter, 1)

		//TODO tomorrow broadcast slows down client to about 500,000 tsx, carrier to 200,000. definite bottleneck, check marshalling
		c.broadcast(initMessage)
	}

	err := conn.Close()
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	log.Info().Msgf("close client connection %s", conn.RemoteAddr())
}

func (c *Carrier) handleCarrierConn(conn net.Conn) {
	decoder := NewBinaryDecoder(conn)
	for {

		// We expect packets framed using util.Frame - they will contain a uint32 (4 bytes) describing the length of the incoming stream
		var m message.Message
		err := decoder.Decode(&m)

		log.Info().Msgf("received %s from %s", m.GetType(), m.GetSenderID())

		// Send to message handler
		err = c.messageHandlers[m.GetType()](m)
		if err != nil {
			log.Error().Msgf(err.Error())
			panic("message handler returned error")
		}

	}
}

func (c *Carrier) decodeNestedSMRDecisions(conn net.Conn) {
	decoder := NewBinaryDecoder(conn)
	for {
		var superBlockSummary SuperBlockSummary
		err := decoder.Decode(&superBlockSummary)
		if err != nil {
			//log.Error().Msgf(err.Error())
			//continue //TODO
			panic(err.Error())
		}
		log.Info().Msgf("received nested SMR decision from %s", conn.RemoteAddr())
		err = c.handleNestedSMRDecision(superBlockSummary)
		if err != nil {
			//log.Error().Msgf(err.Error()) //TODO
			panic(err.Error())
		}
	}
}

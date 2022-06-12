package carrier

import (
	"encoding/gob"
	"encoding/json"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
	"io"
	"net"
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

		c.counter++

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

	// Create a single decoder for a single incoming connection
	decoder := gob.NewDecoder(conn)
	for {
		rawMessage := &message.TransportMessage{}
		err := decoder.Decode(rawMessage)
		if err != nil {
			log.Error().Msgf(err.Error())
			return
		}

		var m message.Message
		switch rawMessage.Type {
		case message.Init:
			m = &message.InitMessage{}
			err = json.Unmarshal(rawMessage.Payload, m)
		case message.Echo:
			m = &message.EchoMessage{}
			err = json.Unmarshal(rawMessage.Payload, m)
		case message.Request:
			m = &message.RequestMessage{}
			err = json.Unmarshal(rawMessage.Payload, m)
		case message.Resolve:
			m = &message.ResolveMessage{}
			err = json.Unmarshal(rawMessage.Payload, m)
		}

		if err != nil {
			log.Error().Msgf(err.Error())
		}

		log.Info().Msgf("received %s from %s", m.GetType(), m.GetSenderID())
		err = c.messageHandlers[rawMessage.Type](m)
		if err != nil {
			log.Error().Msgf(err.Error())
			panic("message handler returned error")
		}
	}
}

func (c *Carrier) decodeNestedSMRDecisions(conn net.Conn) {
	decoder := json.NewDecoder(conn)
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

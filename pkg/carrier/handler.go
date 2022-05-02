package carrier

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
	"io"
	"net"
)

/* Functions in this file are typically invoked as their own goroutines and loop while the connection is open */

func (c *Carrier) handleClientConn(conn net.Conn) {
	// If we didn't manage to connect to the node before, try one last time
	if c.nodeConn == nil {
		c.retryNodeConnection()
	}

	// Make a buffer to hold incoming data.
	// Read the incoming connection into the buffer.
outerLoop:
	for {
		initMessage := &InitMessage{V: make([][]byte, 0)}
		for i := 0; i < util.MempoolThreshold; i++ {
			buf := make([]byte, util.TsxSize) //TODO make this configurable
			_, err := io.ReadAtLeast(conn, buf, util.TsxSize)
			if err != nil {
				log.Info().Msgf(err.Error())
				break outerLoop
			}

			initMessage.V = append(initMessage.V, buf)
		}

		message := &Message{MessageType: Init}
		payload, err := json.Marshal(initMessage)
		if err != nil {
			log.Error().Msgf(err.Error())
			break outerLoop
		}
		message.Payload = payload
		c.broadcast(message)
	}

	//var err error
	//var n int
	//var i = 0
	//for n, err = io.ReadAtLeast(conn, buf, util.TsxSize); err == nil; n, err = io.ReadAtLeast(conn, buf, util.TsxSize) {
	//	var err2 error
	//	log.Info().Msgf("Read %d bytes from %s", n, conn.RemoteAddr())
	//
	//	if c.nodeConn != nil {
	//		_, err2 = c.nodeConn.Write(buf)
	//		log.Info().Msgf("Forwarded %d bytes to %s", n, c.nodeConn.RemoteAddr())
	//	}
	//	if err2 != nil {
	//		log.Error().Msgf(err2.Error())
	//	}
	//	i += n
	//
	//	// Threshold hit, initiate consensus
	//	if i == len(buf) {
	//		message := InitMessage{V: buf}
	//
	//		c.broadcast()
	//	}
	//}
	//log.Info().Msgf(err.Error())

	err2 := conn.Close()
	if err2 != nil {
		log.Error().Msgf(err2.Error())
	}
	log.Info().Msgf("Close client connection %s", conn.RemoteAddr())
}

func (c *Carrier) handleCarrierConn(conn net.Conn) {
	for {
		decoder := gob.NewDecoder(conn)
		rawMessage := &Message{}
		err := decoder.Decode(rawMessage)
		if err != nil {
			log.Error().Msgf(err.Error())
			return
		}

		var message any
		switch rawMessage.MessageType {
		case Init:
			message = &InitMessage{}
			err = json.Unmarshal(rawMessage.Payload, message)
			//c.messageHandlers[Init](message)
		case Echo:
			message = &EchoMessage{}
			err = json.Unmarshal(rawMessage.Payload, message)
		case Request:
			message = &RequestMessage{}
			err = json.Unmarshal(rawMessage.Payload, message)
		case Resolve:
			message = &ResolveMessage{}
			err = json.Unmarshal(rawMessage.Payload, message)
		}

		if err != nil {
			log.Error().Msgf(err.Error())
		}

		err = c.messageHandlers[rawMessage.MessageType](message)
		if err != nil {
			log.Error().Msgf(err.Error())
		}
	}
}

func (c *Carrier) handleInitMessage(rawMessage any) error {
	_, ok := rawMessage.(*InitMessage)
	if !ok {
		return errors.New("expected InitMessage")
	}

	log.Info().Msgf("Received InitMessage")

	return nil
}

func (c *Carrier) handleEchoMessage(rawMessage any) error {
	_, ok := rawMessage.(InitMessage)
	if !ok {
		return errors.New("expected EchoMessage")
	}

	log.Info().Msgf("Received EchoMessage")

	return nil
}

func (c *Carrier) handleRequestMessage(rawMessage any) error {
	_, ok := rawMessage.(InitMessage)
	if !ok {
		return errors.New("expected RequestMessage")
	}

	log.Info().Msgf("Received RequestMessage")

	return nil
}

func (c *Carrier) handleResolveMessage(rawMessage any) error {
	_, ok := rawMessage.(InitMessage)
	if !ok {
		return errors.New("expected ResolveMessage")
	}

	log.Info().Msgf("Received ResolveMessage")

	return nil
}

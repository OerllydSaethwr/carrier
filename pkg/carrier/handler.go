package carrier

import (
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
	"go.dedis.ch/kyber/v4/sign/bdn"
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
		initMessage := &message.InitMessage{
			V:      make([][]byte, 0),
			Sender: c.GetAddress(),
		}
		for i := 0; i < util.MempoolThreshold; i++ {
			buf := make([]byte, util.TsxSize) //TODO make this configurable
			_, err := io.ReadAtLeast(conn, buf, util.TsxSize)
			if err != nil {
				log.Info().Msgf(err.Error())
				break outerLoop
			}

			initMessage.V = append(initMessage.V, buf)
		}

		log.Info().Msgf("Send InitMessage")
		c.broadcast(initMessage)
	}

	err2 := conn.Close()
	if err2 != nil {
		log.Error().Msgf(err2.Error())
	}
	log.Info().Msgf("Close client connection %s", conn.RemoteAddr())
}

func (c *Carrier) handleCarrierConn(conn net.Conn) {
	for {
		decoder := gob.NewDecoder(conn)
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

		err = c.messageHandlers[rawMessage.Type](m)
		if err != nil {
			log.Error().Msgf(err.Error())
			panic("message handler returned error")
		}
	}
}

func (c *Carrier) handleInitMessage(rawMessage message.Message) error {
	initM, ok := rawMessage.(*message.InitMessage)
	if !ok {
		return fmt.Errorf("expected InitMessage")
	}
	log.Info().Msgf("Received InitMessage from %s", initM.GetSender())

	h := initM.Hash()
	s, err := bdn.Sign(c.suite, c.keypair.Sk, h)
	if err != nil {
		return err
	}

	echoM := &message.EchoMessage{
		H:      h,
		S:      s,
		Sender: c.GetAddress(),
	}

	c.broadcast(echoM)

	hstr := hex.EncodeToString(h)
	c.locks.ValueStore.Lock()
	c.valueStore[hstr] = initM.V
	c.locks.ValueStore.Unlock()

	return nil
}

func (c *Carrier) handleEchoMessage(rawMessage message.Message) error {
	var err error
	echoM, ok := rawMessage.(*message.EchoMessage)
	if !ok {
		return fmt.Errorf("expected EchoMessage")
	}
	log.Info().Msgf("Received EchoMessage from %s", echoM.GetSender())

	c.locks.CarrierConns.Lock()
	pk := c.carriers[echoM.GetSender()]
	c.locks.CarrierConns.Unlock()

	err = bdn.Verify(c.suite, pk, echoM.H, echoM.S)
	if err == nil {
		hstr := hex.EncodeToString(echoM.H)
		c.locks.SignatureStore.Lock()
		c.signatureStore[hstr] = append(c.signatureStore[hstr], echoM.S)
		c.locks.SignatureStore.Unlock()

		if len(c.signatureStore[hstr]) == c.f+1 {
			newSBSum := &SuperBlockSummary{
				H:          echoM.H,
				Signatures: c.signatureStore[hstr],
			}
			c.locks.SuperBlockSummary.Lock()
			c.superBlockSummary = append(c.superBlockSummary, newSBSum)
			c.locks.SuperBlockSummary.Unlock()
		}
	}

	c.locks.SuperBlockSummary.Lock()
	if len(c.superBlockSummary) == c.n-c.f {
		err = c.NestedPropose(c.superBlockSummary)
		c.superBlockSummary = make([]*SuperBlockSummary, 0)
	}
	c.locks.SuperBlockSummary.Unlock()
	if err != nil {
		return err
	}

	return nil
}

func (c *Carrier) handleRequestMessage(rawMessage message.Message) error {
	_, ok := rawMessage.(*message.RequestMessage)
	if !ok {
		return fmt.Errorf("expected RequestMessage")
	}

	log.Info().Msgf("Received RequestMessage")

	return nil
}

func (c *Carrier) handleResolveMessage(rawMessage message.Message) error {
	_, ok := rawMessage.(*message.ResolveMessage)
	if !ok {
		return fmt.Errorf("expected ResolveMessage")
	}

	log.Info().Msgf("Received ResolveMessage")

	return nil
}

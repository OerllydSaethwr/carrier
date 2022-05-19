package carrier

import (
	"encoding/gob"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/rs/zerolog/log"
	"net"
)

func (c *Carrier) broadcast(message message.Message) {
	transportMessage := message.Marshal()

	for _, addr := range c.carrierAddrs {
		c.send(addr, transportMessage)
	}
}

func (c *Carrier) send(dest *net.TCPAddr, message *message.TransportMessage) {
	conn, ok := c.carrierConns[dest]
	if !ok {
		log.Error().Msgf("Cannot find connection to address %s", dest.String())
		panic(1)
		return
	}

	encoder := gob.NewEncoder(conn)

	// Send to dest
	err := encoder.Encode(message)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
}

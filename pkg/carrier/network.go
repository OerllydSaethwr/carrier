package carrier

import (
	"encoding/gob"
	"github.com/rs/zerolog/log"
	"net"
)

func (c *Carrier) broadcast(message *Message) {
	for _, addr := range c.carrierAddrs {
		c.send(addr, message)
	}
}

func (c *Carrier) send(dest *net.TCPAddr, message *Message) {
	conn, ok := c.carrierConns[dest]
	if !ok {
		log.Error().Msgf("Cannot find connection to address %s", dest.String())
		return
	}

	encoder := gob.NewEncoder(conn)

	// Send to dest
	err := encoder.Encode(message)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
}

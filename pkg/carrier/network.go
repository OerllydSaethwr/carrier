package carrier

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/rs/zerolog/log"
)

func (c *Carrier) broadcast(message message.Message) {
	c.broadcastDispenser <- message
}

// For communicating with carriers
func (c *Carrier) executeBroadcast(message message.Message) {
	log.Debug().Msgf("broadcast %s", message.GetType())

	buf := message.BinaryMarshal()

	for _, n := range c.neighbours {
		n.Send(buf)
	}
}

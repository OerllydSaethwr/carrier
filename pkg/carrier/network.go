package carrier

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/rs/zerolog/log"
)

func (c *Carrier) Broadcast(message message.Message) {
	c.BroadcastDispenser <- message
}

// For communicating with carriers
func (c *Carrier) ExecuteBroadcast(message message.Message) {
	log.Debug().Msgf("broadcast %s", message.GetType())

	buf := message.BinaryMarshal()

	for _, n := range c.Neighbours {
		n.Send(buf)
	}
}

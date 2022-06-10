package carrier

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/rs/zerolog/log"
)

// For communicating with carriers
func (c *Carrier) broadcast(message message.Message) {
	log.Info().Msgf("broadcast %s", message.GetType())
	transportMessage := message.Marshal()

	for _, n := range c.neighbours {
		n.send(transportMessage)
	}
}

func (n *Neighbour) send(message *message.TransportMessage) {
	if !n.IsAlive() {
		n.WaitUntilAlive()
	}

	// Send to dest
	err := n.GetEncoder().Encode(message)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
}

func (n *Neighbour) marshalAndSend(message message.Message) {
	n.send(message.Marshal())
}

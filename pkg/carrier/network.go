package carrier

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
)

func (c *Carrier) broadcast(message message.Message) {
	c.broadcastDispenser <- message
}

// For communicating with carriers
func (c *Carrier) executeBroadcast(message message.Message) {
	log.Info().Msgf("broadcast %s", message.GetType())

	buf := message.BinaryMarshal()

	//println(len(buf))
	//n0 := buf[8]
	//n1 := util.UnmarshalUInt64(buf[9:17])
	//n2 := util.UnmarshalUInt64(buf[17:])
	//println(n0)
	//println(n1)
	//println(n2)
	//transportMessage := message.Marshal()

	for _, n := range c.neighbours {
		n.send(buf)
	}
}

// @Bottleneck B1 - double marshalling (try using gob only)
// DO NOT USE - send expects a framedBuf, it will not do the framing. Use framedAndSend instead
func (n *Neighbour) send(buf []byte) {

	// TODO figure something out with this
	// Send to dest
	err := n.GetEncoder().Encode(buf)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
}

func (n *Neighbour) frameAndSend(buf []byte) {
	n.send(util.Frame(buf))
}

func (n *Neighbour) marshalAndSend(message message.Message) {
	//n.send(message.Marshal())
}

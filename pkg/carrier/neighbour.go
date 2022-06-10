package carrier

import (
	"encoding/gob"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
	"net"
)

type Neighbour struct {
	util.Neighbour
	conn           *net.TCPConn
	encoder        *gob.Encoder
	decoder        *gob.Decoder
	waitUntilAlive chan int // Dummy channel that we use to block other processes until the connection becomes live
}

func NewNeighbour(id, address, pk string) *Neighbour {
	n := &Neighbour{
		Neighbour: util.Neighbour{
			ID:      id,
			Address: address,
			PK:      pk,
		},
		conn:           nil,
		encoder:        nil,
		decoder:        nil,
		waitUntilAlive: make(chan int),
	}

	return n
}

func (n *Neighbour) GetID() string {
	return n.ID
}

func (n *Neighbour) GetAddress() string {
	return n.Address
}

func (n *Neighbour) GetPK() string {
	return n.PK
}

func (n *Neighbour) GetConn() *net.TCPConn {
	return n.conn
}

func (n *Neighbour) GetEncoder() *gob.Encoder {
	return n.encoder
}

// @Unused
func (n *Neighbour) GetDecoder() *gob.Decoder {
	return n.decoder
}

func (n *Neighbour) IsAlive() bool {
	return n.conn != nil
}

func (n *Neighbour) WaitUntilAlive() {
	log.Info().Msgf("waiting for connection with %s to come alive...", n.GetAddress())
	<-n.waitUntilAlive
	return
}

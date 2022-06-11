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

// GetEncoder will block until connection is alive
func (n *Neighbour) GetEncoder() Encoder {
	n.WaitUntilAlive()
	return n.encoder
}

func (n *Neighbour) IsAlive() bool {
	return n.conn != nil
}

func (n *Neighbour) WaitUntilAlive() {
	if !n.IsAlive() {
		log.Info().Msgf("waiting for connection with %s to come alive...", n.GetAddress())
		<-n.waitUntilAlive
	}
	return
}

func (n *Neighbour) SetConnAndEncoderAndSignalAlive(conn *net.TCPConn) {
	n.conn = conn
	n.encoder = gob.NewEncoder(conn)
	close(n.waitUntilAlive)
}

func (n *Neighbour) GetType() string {
	return "carrier"
}

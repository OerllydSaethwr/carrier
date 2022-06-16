package carrier

import (
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
	"net"
	"sync"
)

type Neighbour struct {
	util.Neighbour
	encoder        *BinaryEncoder
	waitUntilAlive chan int // Dummy channel that we use to block other processes until the connection becomes live
	connLock       *sync.RWMutex
}

func NewNeighbour(id, address, pk string) *Neighbour {
	n := &Neighbour{
		Neighbour: util.Neighbour{
			ID:      id,
			Address: address,
			PK:      pk,
		},
		encoder:        NewBinaryEncoder(nil),
		waitUntilAlive: make(chan int),
		connLock:       &sync.RWMutex{},
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
	return n.encoder.conn != nil
}

func (n *Neighbour) WaitUntilAlive() {
	if !n.IsAlive() {
		log.Info().Msgf("waiting for connection with %s (carrier) to come alive...", n.GetAddress())
		<-n.waitUntilAlive
	}
	return
}

func (n *Neighbour) SetConnAndEncoderAndSignalAlive(conn *net.TCPConn) {
	n.encoder = NewBinaryEncoder(conn)
	close(n.waitUntilAlive)
}

func (n *Neighbour) GetType() string {
	return "carrier"
}

func (n *Neighbour) GetConnLock() *sync.RWMutex {
	return n.connLock
}

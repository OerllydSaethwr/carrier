package remote

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier/codec"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
	"net"
	"sync"
)

type Neighbour struct {
	util.Neighbour
	Encoder            *codec.BinaryEncoder
	WaitUntilAliveChan chan int // Dummy channel that we use to block other processes until the connection becomes live
	ConnLock           *sync.RWMutex
}

func NewNeighbour(id, address, pk string) *Neighbour {
	n := &Neighbour{
		Neighbour: util.Neighbour{
			ID:      id,
			Address: address,
			PK:      pk,
		},
		Encoder:            codec.NewBinaryEncoder(nil),
		WaitUntilAliveChan: make(chan int),
		ConnLock:           &sync.RWMutex{},
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
func (n *Neighbour) GetEncoder() codec.Encoder {
	n.WaitUntilAlive()
	return n.Encoder
}

func (n *Neighbour) IsAlive() bool {
	return n.Encoder.Conn != nil
}

func (n *Neighbour) WaitUntilAlive() {
	if !n.IsAlive() {
		log.Info().Msgf("waiting for connection with %s (carrier) to come alive...", n.GetAddress())
		<-n.WaitUntilAliveChan
	}
	return
}

func (n *Neighbour) SetConnAndEncoderAndSignalAlive(conn *net.TCPConn) {
	n.Encoder = codec.NewBinaryEncoder(conn)
	close(n.WaitUntilAliveChan)
}

func (n *Neighbour) GetType() string {
	return "carrier"
}

func (n *Neighbour) GetConnLock() *sync.RWMutex {
	return n.ConnLock
}

func (n *Neighbour) Send(buf []byte) {

	// Send to dest
	err := n.GetEncoder().Encode(buf)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
}

func (n *Neighbour) MarshalAndSend(message message.Message) {
	n.Send(message.BinaryMarshal())
}

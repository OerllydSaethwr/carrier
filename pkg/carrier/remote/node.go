package remote

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier/codec"
	"github.com/rs/zerolog/log"
	"net"
)

type Node struct {
	Address            string
	Encoder            *codec.BinaryEncoder
	WaitUntilAliveChan chan int
}

func NewNode(address string) *Node {
	return &Node{
		Address:            address,
		Encoder:            codec.NewBinaryEncoder(nil),
		WaitUntilAliveChan: make(chan int),
	}
}

func (n *Node) GetAddress() string {
	return n.Address
}

// GetEncoder will block until connection is alive
func (n *Node) GetEncoder() codec.Encoder {
	n.WaitUntilAlive()
	return n.Encoder
}

func (n *Node) IsAlive() bool {
	return n.Encoder.Conn != nil
}

func (n *Node) WaitUntilAlive() {
	if !n.IsAlive() {
		log.Info().Msgf("waiting for connection with %s (node) to come alive...", n.GetAddress())
		<-n.WaitUntilAliveChan
	}
	return
}

func (n *Node) SetConnAndEncoderAndSignalAlive(conn *net.TCPConn) {
	n.Encoder = codec.NewBinaryEncoder(conn)
	close(n.WaitUntilAliveChan)
}

func (n *Node) GetType() string {
	return "node"
}

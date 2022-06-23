package remote

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier/codec"
	"github.com/rs/zerolog/log"
	"net"
)

type Node struct {
	address        string
	encoder        *codec.BinaryEncoder
	waitUntilAlive chan int
}

func NewNode(address string) *Node {
	return &Node{
		address:        address,
		encoder:        codec.NewBinaryEncoder(nil),
		waitUntilAlive: make(chan int),
	}
}

func (n *Node) GetAddress() string {
	return n.address
}

// GetEncoder will block until connection is alive
func (n *Node) GetEncoder() codec.Encoder {
	n.WaitUntilAlive()
	return n.encoder
}

func (n *Node) IsAlive() bool {
	return n.encoder.Conn != nil
}

func (n *Node) WaitUntilAlive() {
	if !n.IsAlive() {
		log.Info().Msgf("waiting for connection with %s (node) to come alive...", n.GetAddress())
		<-n.waitUntilAlive
	}
	return
}

func (n *Node) SetConnAndEncoderAndSignalAlive(conn *net.TCPConn) {
	n.encoder = codec.NewBinaryEncoder(conn)
	close(n.waitUntilAlive)
}

func (n *Node) GetType() string {
	return "node"
}

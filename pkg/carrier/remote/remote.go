package remote

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier/codec"
	"net"
)

type Remote interface {
	GetAddress() string
	GetEncoder() codec.Encoder
	SetConnAndEncoderAndSignalAlive(conn *net.TCPConn)
	GetType() string
}

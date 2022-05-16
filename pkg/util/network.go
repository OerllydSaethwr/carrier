package util

import (
	"net"
)

// ListenTCP Simple wrawpper around net.ListenTCP
func ListenTCP(laddr *net.TCPAddr) (*net.TCPListener, error) {
	return net.ListenTCP(Network, laddr)
}

// DialTCP Simple wrapper around net.DialTCP
func DialTCP(raddr *net.TCPAddr) (*net.TCPConn, error) {
	return net.DialTCP(Network, nil, raddr)
}

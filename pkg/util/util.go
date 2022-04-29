package util

import (
	"net"
	"strconv"
)

func SplitHostPort(hostport string) (string, int, error) {
	host, portstring, err := net.SplitHostPort(hostport)
	if err != nil {
		return "", 0, err
	}
	port, err := strconv.Atoi(portstring)

	return host, port, err
}

// Wrapper around net.DialTCP
func DialTCP(raddr *net.TCPAddr) (*net.TCPConn, error) {
	return net.DialTCP(Network, nil, raddr)
}

package util

import (
	"time"
)

const (
	Network = "tcp4" // Type of network, choices are in net.DialTCP

	PortsPerCarrier = 3 // Used by the config generator to know how many ports to reserve for each carrier

	//LogTimeFormat = zerolog.TimeFormatUnixMs
	LogTimeFormat = time.RFC3339
)

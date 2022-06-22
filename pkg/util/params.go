package util

import (
	"github.com/rs/zerolog"
	"time"
)

const (
	Network               = "tcp4" // Type of network, choices are in net.DialTCP
	TsxSize               = 512    // Size of transactions in bytes
	Rate                  = 10000000
	Nodes                 = 4
	CarrierConnRetryDelay = time.Second // Delay between retries
	CarrierConnMaxRetry   = 0           // Number of max retries - 0 means keep trying forever
	NodeConnRetryDelay    = time.Second // Delay between retries
	NodeConnMaxRetry      = 0           // Number of max retries - 0 means keep trying forever
	MempoolThreshold      = 1000        // Threshold in number of transactions before we initiate a consensus with carriers

	BasePort        = 6000 // Used by the config generator for the base port
	PortsPerCarrier = 3    // Used by the config generator to know how many ports to reserve for each carrier
	FrontPort       = 9000 // Base front port for nodes

	Decision = 3000 // Addresses used by the remote benchmarks. Make sure these are the same as in the settings.json
	Client   = 5000
	Carrier  = 4000

	ForwardMode = false // If set to true, carrier will not do any processing and forward client tsx as they are

	SinkWriteTimeout = time.Second // Drop message after this duration if the sink is full

	LogLevel = zerolog.DebugLevel // Coordinating log settings between main program and tests
	//LogTimeFormat = zerolog.TimeFormatUnixMs
	LogTimeFormat = time.RFC3339
)

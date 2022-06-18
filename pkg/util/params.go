package util

import "time"

const (
	Network               = "tcp4"      // Type of network, choices are in net.DialTCP
	TsxSize               = 9           // Size of transactions in bytes
	CarrierConnRetryDelay = time.Second // Delay between retries
	CarrierConnMaxRetry   = 0           // Number of max retries - 0 means keep trying forever
	NodeConnRetryDelay    = time.Second // Delay between retries
	NodeConnMaxRetry      = 0           // Number of max retries - 0 means keep trying forever
	MempoolThreshold      = 10          // Threshold in number of transactions before we initiate a consensus with carriers

	BasePort        = 6000 // Used by the config generator for the base port
	PortsPerCarrier = 3    // Used by the config generator to know how many ports to reserve for each carrier
	FrontPort       = 9000 // Base front port for nodes

	Decision = 3000 // Addresses used by the remote benchmarks. Make sure these are the same as in the settings.json
	Client   = 5000
	Carrier  = 4000

	ForwardMode = false // If set to true, carrier will not do any processing and forward client tsx as they are
)

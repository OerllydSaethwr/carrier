package util

import "time"

const (
	Network               = "tcp4"      // Type of network, choices are in net.DialTCP
	TsxSize               = 9           // Size of transactions in bytes
	CarrierConnRetryDelay = time.Second // Delay between retries
	CarrierConnMaxRetry   = 10          // Number of max retries
	NodeConnRetryDelay    = time.Second // Delay between retries
	NodeConnMaxRetry      = 10          // Number of max retries
	MempoolThreshold      = 10          // Threshold in number of transactions before we initiate a consensus with carriers //TODO carriers can't keep up with batch rate lower than 10 and data will get garbled
	BasePort              = 6000        // Used by the config generator for the base port
	PortsPerCarrier       = 3           // Used by the config generator to know how many ports to reserve for each carrier
	FrontPort             = 9000        // Base front port for nodes
)

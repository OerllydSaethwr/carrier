package util

import "time"

const (
	Network               = "tcp4"      // Type of network, choices are in net.DialTCP
	TsxSize               = 9           // Size of transactions in bytes
	CarrierConnRetryDelay = time.Second // Delay between retries
	CarrierConnMaxRetry   = 10          // Number of max retries
	MempoolThreshold      = 10          // Threshold in number of transactions before we initiate a consensus with carriers
)

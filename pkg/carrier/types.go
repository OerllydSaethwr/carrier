package carrier

import (
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"net"
	"sync"
	"time"
)

type Locks struct {
	ValueStore        *sync.RWMutex
	SignatureStore    *sync.RWMutex
	SuperBlockSummary *sync.RWMutex
	AcceptedHashStore *sync.RWMutex

	// This lock is a known bottleneck
	DecisionLock *sync.RWMutex
}

type Stores struct {
	valueStore        map[string][][]byte
	signatureStore    map[string][]util.Signature
	superBlockSummary SuperBlockSummary
	acceptedHashStore AcceptedHashStore
}

type Listeners struct {
	clientListener   *net.TCPListener
	carrierListener  *net.TCPListener
	decisionListener *net.TCPListener
}

type ConnRetrySettings struct {
	carrierConnRetryDelay time.Duration
	carrierConnMaxRetry   uint
	nodeConnRetryDelay    time.Duration
	nodeConnMaxRetry      uint
}

type Addresses struct {
	client2carrier  string
	carrier2carrier string
	decision        string
}

type SuperBlockSummary struct {
	id      uint32
	payload map[string][]util.Signature
}

type AcceptedHashStore struct {
	id      uint32
	payload map[string][][]byte
}

type Remote interface {
	GetAddress() string
	GetEncoder() Encoder
	SetConnAndEncoderAndSignalAlive(conn *net.TCPConn)
	GetType() string
}

type Encoder interface {
	Encode(e any) error
}

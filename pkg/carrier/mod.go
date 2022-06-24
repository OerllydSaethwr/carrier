package carrier

import (
	"encoding/hex"
	"fmt"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/remote"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/superblock"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/pairing"
	"go.dedis.ch/kyber/v4/sign/bdn"
	"math"
	"net"
	"sync"
	"time"
)

type Carrier struct {
	Counter uint64

	Config *util.Config

	Listeners Listeners
	Stores    Stores
	Locks     Locks

	NodeConn *net.TCPConn
	Node     *remote.Node

	Neighbours map[string]*remote.Neighbour

	// Registry of message handlers. Argument must be one of the enum types
	MessageHandlers map[message.Type]func(message.Message) error

	Suite *pairing.SuiteBn256

	F int
	N int

	Wg *sync.WaitGroup

	Quit               chan bool
	BroadcastDispenser chan message.Message
	SbsCounter         int
}

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
	superBlockSummary map[string][]util.Signature
	acceptedHashStore map[string][][]byte

	decidedHashStore map[string]interface{}
}

type Listeners struct {
	clientListener   *net.TCPListener
	carrierListener  *net.TCPListener
	decisionListener *net.TCPListener
}

func NewCarrier(config *util.Config) *Carrier {
	neighbours := map[string]*remote.Neighbour{}
	for _, n := range config.Neighbours {
		neighbours[n.ID] = remote.NewNeighbour(n.ID, n.Address, n.PK)
	}

	c := &Carrier{}
	c.Config = config
	c.Quit = make(chan bool, 1)

	// TEMP
	c.Counter = 0
	c.BroadcastDispenser = make(chan message.Message, int64(math.Pow10(7)))
	c.SbsCounter = 0

	c.Suite = pairing.NewSuiteBn256()

	c.Node = remote.NewNode(config.Addresses.Front)
	c.Neighbours = neighbours

	c.F = (len(neighbours) - 1) / 3
	c.N = len(neighbours)

	c.MessageHandlers = map[message.Type]func(message.Message) error{
		message.Init:    c.HandleInitMessage,
		message.Echo:    c.HandleEchoMessage,
		message.Request: c.HandleRequestMessage,
		message.Resolve: c.HandleResolveMessage,
	}

	c.Locks = Locks{
		ValueStore:        &sync.RWMutex{},
		SignatureStore:    &sync.RWMutex{},
		SuperBlockSummary: &sync.RWMutex{},
		AcceptedHashStore: &sync.RWMutex{},
		DecisionLock:      &sync.RWMutex{},
	}

	c.Stores = Stores{
		valueStore:        map[string][][]byte{},
		signatureStore:    map[string][]util.Signature{},
		superBlockSummary: map[string][]util.Signature{},
		acceptedHashStore: map[string][][]byte{},

		decidedHashStore: map[string]any{},
	}

	//var err error

	//// Check client listener addr
	//_, err = util.ResolveTCPAddr(clientToCarrierAddr)
	//if err != nil {
	//	log.Error().Msgf(err.Error())
	//	return nil
	//}
	//
	//// TODO move this
	//// Check carrier listener addr
	//_, err = util.ResolveTCPAddr(carrierToCarrierAddr)
	//if err != nil {
	//	log.Error().Msgf(err.Error())
	//	return nil
	//}
	//
	//// Check front addr
	//_, err = util.ResolveTCPAddr(frontAddr)
	//if err != nil {
	//	log.Error().Msgf(err.Error())
	//	return nil
	//}
	//
	//// Check other carrier addrs
	//for _, n := range c.neighbours {
	//	_, err := util.ResolveTCPAddr(n.Address)
	//	if err != nil {
	//		log.Error().Msgf(err.Error())
	//	}
	//}

	return c
}

/*	Start listening to client requests
	Forward client requests
	We are not waiting for listeners to stop but I think it's fine
*/
func (c *Carrier) Start() *sync.WaitGroup {
	if c.ForwardMode() {
		log.Info().Msgf("ForwardMode is turned on - logging of tsx is at debug level to avoid flooding. If you want to see individual logs, set log level to debug or higher.")
	}

	log.Info().Msgf("init-threshold: %d", c.GetMempoolThreshold())

	c.Wg = &sync.WaitGroup{}
	c.Wg.Add(1)

	var err error

	// Listen to nested SMR decisions
	c.Listeners.decisionListener, err = c.StartListener(c.GetDecisionAddress())
	if err != nil {
		log.Error().Msgf(err.Error())
		c.Wg.Done()
		return c.Wg
	}
	log.Info().Msgf("start listening to nested SMR decisions on %s", c.GetDecisionAddress())
	go c.HandleIncomingConnections(c.Listeners.decisionListener, c.DecodeNestedSMRDecisions)

	// Listen to client
	c.Listeners.clientListener, err = c.StartListener(c.GetClientToCarrierAddress())
	if err != nil {
		log.Error().Msgf(err.Error())
		c.Wg.Done()
		return c.Wg
	}
	log.Info().Msgf("start listening to client on %s", c.GetClientToCarrierAddress())
	go c.HandleIncomingConnections(c.Listeners.clientListener, c.HandleClientConn)

	// Listen to carrier
	c.Listeners.carrierListener, err = c.StartListener(c.GetCarrierToCarrierAddress())
	if err != nil {
		log.Error().Msgf(err.Error())
		c.Wg.Done()
		return c.Wg
	}
	log.Info().Msgf("start listening to neighbours on %s", c.GetCarrierToCarrierAddress())
	go c.HandleIncomingConnections(c.Listeners.carrierListener, c.HandleCarrierConn)

	//Connect to node
	go connect(c.Node, time.Duration(c.Config.Settings.NodeConnRetryDelay)*time.Millisecond, c.Config.Settings.NodeConnMaxRetry)

	// Set up connections to other neighbours
	for _, n := range c.Neighbours {
		go connect(n, time.Duration(c.Config.Settings.CarrierConnRetryDelay)*time.Millisecond, c.Config.Settings.CarrierConnMaxRetry)
	}

	//go c.logger()
	c.LaunchWorkerPool(10, c.BroadcastWorker)

	return c.Wg
}

func (c *Carrier) Stop() {
	log.Trace().Msgf("stop Carrier")
	err := c.Listeners.clientListener.Close()
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	c.Quit <- true
	c.Wg.Done()
}

func (c *Carrier) GetAddress() string {
	return c.GetCarrierToCarrierAddress()
}

func (c *Carrier) NestedPropose(P superblock.SuperBlockSummary) error {

	err := c.Node.GetEncoder().Encode(&P)
	if err != nil {
		return err
	}

	log.Info().Msgf("proposed SuperBlock %d to %s", c.SbsCounter, c.Node.GetAddress())
	return nil
}

func (c *Carrier) GetStringSK() string {
	return c.Config.Keys.Sk
}

func (c *Carrier) GetStringPK() string {
	return c.Config.Keys.Pk
}

func (c *Carrier) GetKyberSK() kyber.Scalar {
	skk, err := util.DecodeStringToBdnSK(c.Config.Keys.Sk)
	if err != nil {
		panic("unable to decode SK")
	}

	return skk
}

func (c *Carrier) GetKyberPK() kyber.Point {
	pkk, err := util.DecodeStringToBdnPK(c.Config.Keys.Pk)
	if err != nil {
		panic("unable to decode SK")
	}

	return pkk
}

func (c *Carrier) Sign(h string) string {
	hb, err := hex.DecodeString(h)
	if err != nil {
		panic("signing failed: failed to decode h")
	}
	s, err := bdn.Sign(c.Suite, c.GetKyberSK(), hb)
	if err != nil {
		panic("signing failed")
	}

	return hex.EncodeToString(s)
}

func (c *Carrier) Verify(h string, s util.Signature) error {
	pk, err := c.GetPKFromID(s.SenderID)
	if err != nil {
		return err
	}

	hb, err := hex.DecodeString(h)
	if err != nil {
		return fmt.Errorf("failed to decode h: %s", err.Error())
	}
	sb, err := hex.DecodeString(s.S)
	if err != nil {
		return fmt.Errorf("failed to decode s: %s", err.Error())
	}
	err = bdn.Verify(c.Suite, pk, hb, sb)
	return err
}

func (c *Carrier) GetPKFromID(senderID string) (kyber.Point, error) {

	n, ok := c.Neighbours[senderID]
	if !ok {
		return nil, fmt.Errorf("pk not found in store")
	}

	pkk, err := util.DecodeStringToBdnPK(n.PK)
	if err != nil {
		return nil, err
	}

	return pkk, nil
}

func (c *Carrier) GetSuite() pairing.Suite {
	return c.Suite.Suite
}

func Decide(D map[string][][]byte) {

	// Process decided values
	for h, _ := range D {
		log.Info().Msgf("committed %s", h)
	}
}

func (c *Carrier) GetClientToCarrierAddress() string {
	return c.Config.Addresses.Client
}

func (c *Carrier) GetCarrierToCarrierAddress() string {
	return c.Config.Addresses.Carrier
}

func (c *Carrier) GetDecisionAddress() string {
	return c.Config.Addresses.Decision
}

func (c *Carrier) GetID() string {
	return c.Config.ID
}

func (c *Carrier) StartListener(address string) (*net.TCPListener, error) {
	resolvedAddress, err := util.ResolveTCPAddr(address)
	resolvedAddress.IP = net.ParseIP("0.0.0.0") // We can only host on localhost
	if err != nil {
		return nil, err
	}
	listener, err := util.ListenTCP(resolvedAddress)

	return listener, err
}

func (c *Carrier) ForwardMode() bool {
	return c.Config.Settings.ForwardMode == 1
}

func (c *Carrier) GetTsxSize() int {
	return c.Config.Settings.TsxSize
}

func (c *Carrier) GetMempoolThreshold() int {
	return c.Config.Settings.InitThreshold
}

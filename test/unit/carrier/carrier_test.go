package carrier

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/remote"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/superblock"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/pairing"
	"net"
	"reflect"
	"sync"
	"testing"
)

func TestCarrier_GetAddress(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if got := c.GetAddress(); got != tt.want {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarrier_GetID(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if got := c.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarrier_GetKyberPK(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	tests := []struct {
		name   string
		fields fields
		want   kyber.Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if got := c.GetKyberPK(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKyberPK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarrier_GetKyberSK(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	tests := []struct {
		name   string
		fields fields
		want   kyber.Scalar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if got := c.GetKyberSK(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKyberSK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarrier_GetPKFromID(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	type args struct {
		senderID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    kyber.Point
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			got, err := c.GetPKFromID(tt.args.senderID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPKFromID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPKFromID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarrier_GetStringPK(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if got := c.GetStringPK(); got != tt.want {
				t.Errorf("GetStringPK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarrier_GetStringSK(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if got := c.GetStringSK(); got != tt.want {
				t.Errorf("GetStringSK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarrier_GetSuite(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	tests := []struct {
		name   string
		fields fields
		want   pairing.Suite
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if got := c.GetSuite(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSuite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarrier_NestedPropose(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	type args struct {
		P superblock.SuperBlockSummary
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if err := c.NestedPropose(tt.args.P); (err != nil) != tt.wantErr {
				t.Errorf("NestedPropose() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCarrier_Sign(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	type args struct {
		h string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if got := c.Sign(tt.args.h); got != tt.want {
				t.Errorf("Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarrier_Start(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	tests := []struct {
		name   string
		fields fields
		want   *sync.WaitGroup
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if got := c.Start(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarrier_Stop(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			c.Stop()
		})
	}
}

func TestCarrier_Verify(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	type args struct {
		h string
		s util.Signature
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if err := c.Verify(tt.args.h, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCarrier_broadcast(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	type args struct {
		message message.Message
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			c.Broadcast(tt.args.message)
		})
	}
}

func TestCarrier_broadcastWorker(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			c.BroadcastWorker()
		})
	}
}

func TestCarrier_checkAcceptedHashStoreAndDecide(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			c.CheckAcceptedHashStoreAndDecide()
		})
	}
}

func TestCarrier_decodeNestedSMRDecisions(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			c.DecodeNestedSMRDecisions(tt.args.conn)
		})
	}
}

func TestCarrier_executeBroadcast(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	type args struct {
		message message.Message
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			c.ExecuteBroadcast(tt.args.message)
		})
	}
}

func TestCarrier_forwardMode(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if got := c.ForwardMode(); got != tt.want {
				t.Errorf("forwardMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarrier_getCarrierToCarrierAddress(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if got := c.GetCarrierToCarrierAddress(); got != tt.want {
				t.Errorf("getCarrierToCarrierAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarrier_getClientToCarrierAddress(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if got := c.GetClientToCarrierAddress(); got != tt.want {
				t.Errorf("getClientToCarrierAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarrier_getDecisionAddress(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if got := c.GetDecisionAddress(); got != tt.want {
				t.Errorf("getDecisionAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarrier_getMempoolThreshold(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if got := c.GetMempoolThreshold(); got != tt.want {
				t.Errorf("getMempoolThreshold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarrier_getTsxSize(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if got := c.GetTsxSize(); got != tt.want {
				t.Errorf("getTsxSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarrier_handleCarrierConn(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			c.HandleCarrierConn(tt.args.conn)
		})
	}
}

func TestCarrier_handleClientConn(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			c.HandleClientConn(tt.args.conn)
		})
	}
}

func TestCarrier_handleEchoMessage(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	type args struct {
		rawMessage message.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if err := c.HandleEchoMessage(tt.args.rawMessage); (err != nil) != tt.wantErr {
				t.Errorf("handleEchoMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCarrier_handleIncomingConnections(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	type args struct {
		l       *net.TCPListener
		handler func(conn net.Conn)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			c.HandleIncomingConnections(tt.args.l, tt.args.handler)
		})
	}
}

func TestCarrier_handleInitMessage(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	type args struct {
		rawMessage message.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if err := c.HandleInitMessage(tt.args.rawMessage); (err != nil) != tt.wantErr {
				t.Errorf("handleInitMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCarrier_handleNestedSMRDecision(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	type args struct {
		N superblock.SuperBlockSummary
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			c.HandleNestedSMRDecision(tt.args.N)
		})
	}
}

func TestCarrier_handleRequestMessage(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	type args struct {
		rawMessage message.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if err := c.HandleRequestMessage(tt.args.rawMessage); (err != nil) != tt.wantErr {
				t.Errorf("handleRequestMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCarrier_handleResolveMessage(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	type args struct {
		rawMessage message.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			if err := c.HandleResolveMessage(tt.args.rawMessage); (err != nil) != tt.wantErr {
				t.Errorf("handleResolveMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCarrier_launchWorkerPool(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	type args struct {
		poolSize int
		task     func()
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			c.LaunchWorkerPool(tt.args.poolSize, tt.args.task)
		})
	}
}

func TestCarrier_logger(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			c.Logger()
		})
	}
}

func TestCarrier_startListener(t *testing.T) {
	type fields struct {
		counter            uint64
		config             *util.Config
		listeners          carrier.Listeners
		stores             carrier.Stores
		locks              carrier.Locks
		nodeConn           *net.TCPConn
		node               *remote.Node
		neighbours         map[string]*remote.Neighbour
		messageHandlers    map[message.Type]func(message.Message) error
		suite              *pairing.SuiteBn256
		f                  int
		n                  int
		wg                 *sync.WaitGroup
		quit               chan bool
		broadcastDispenser chan message.Message
		sbsCounter         int
	}
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *net.TCPListener
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Carrier{
				Counter:            tt.fields.counter,
				Config:             tt.fields.config,
				Listeners:          tt.fields.listeners,
				Stores:             tt.fields.stores,
				Locks:              tt.fields.locks,
				NodeConn:           tt.fields.nodeConn,
				Node:               tt.fields.node,
				Neighbours:         tt.fields.neighbours,
				MessageHandlers:    tt.fields.messageHandlers,
				Suite:              tt.fields.suite,
				F:                  tt.fields.f,
				N:                  tt.fields.n,
				Wg:                 tt.fields.wg,
				Quit:               tt.fields.quit,
				BroadcastDispenser: tt.fields.broadcastDispenser,
				SbsCounter:         tt.fields.sbsCounter,
			}
			got, err := c.StartListener(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("startListener() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("startListener() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_Write(t *testing.T) {
	type fields struct {
		conn    net.Conn
		sink    chan []byte
		address string
	}
	type args struct {
		buf []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Conn{
				Conn:    tt.fields.conn,
				Sink:    tt.fields.sink,
				Address: tt.fields.address,
			}
			c.Write(tt.args.buf)
		})
	}
}

func TestConn_connect(t *testing.T) {
	type fields struct {
		conn    net.Conn
		sink    chan []byte
		address string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Conn{
				Conn:    tt.fields.conn,
				Sink:    tt.fields.sink,
				Address: tt.fields.address,
			}
			if err := c.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConn_runSinkConsumer(t *testing.T) {
	type fields struct {
		conn    net.Conn
		sink    chan []byte
		address string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carrier.Conn{
				Conn:    tt.fields.conn,
				Sink:    tt.fields.sink,
				Address: tt.fields.address,
			}
			c.RunSinkConsumer()
		})
	}
}

func TestConnect(t *testing.T) {
	type args struct {
		address        string
		sinkBufferSize int
	}
	tests := []struct {
		name string
		args args
		want *carrier.Conn
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := carrier.Connect(tt.args.address, tt.args.sinkBufferSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCarrier(t *testing.T) {
	type args struct {
		config *util.Config
	}
	tests := []struct {
		name string
		args args
		want *carrier.Carrier
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := carrier.NewCarrier(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCarrier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_connect(t *testing.T) {
	type args struct {
		s1 string
		i1 int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			carrier.Connect(tt.args.s1, tt.args.i1)
		})
	}
}

func Test_decide(t *testing.T) {
	type args struct {
		D map[string][][]byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			carrier.Decide(tt.args.D)
		})
	}
}

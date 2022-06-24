package remote

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier/codec"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/remote"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"net"
	"reflect"
	"sync"
	"testing"
)

func TestNeighbour_GetAddress(t *testing.T) {
	type fields struct {
		Neighbour      util.Neighbour
		encoder        *codec.BinaryEncoder
		waitUntilAlive chan int
		connLock       *sync.RWMutex
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
			n := &remote.Neighbour{
				Neighbour:          tt.fields.Neighbour,
				Encoder:            tt.fields.encoder,
				WaitUntilAliveChan: tt.fields.waitUntilAlive,
				ConnLock:           tt.fields.connLock,
			}
			if got := n.GetAddress(); got != tt.want {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeighbour_GetConnLock(t *testing.T) {
	type fields struct {
		Neighbour      util.Neighbour
		encoder        *codec.BinaryEncoder
		waitUntilAlive chan int
		connLock       *sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   *sync.RWMutex
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &remote.Neighbour{
				Neighbour:          tt.fields.Neighbour,
				Encoder:            tt.fields.encoder,
				WaitUntilAliveChan: tt.fields.waitUntilAlive,
				ConnLock:           tt.fields.connLock,
			}
			if got := n.GetConnLock(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConnLock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeighbour_GetEncoder(t *testing.T) {
	type fields struct {
		Neighbour      util.Neighbour
		encoder        *codec.BinaryEncoder
		waitUntilAlive chan int
		connLock       *sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   codec.Encoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &remote.Neighbour{
				Neighbour:          tt.fields.Neighbour,
				Encoder:            tt.fields.encoder,
				WaitUntilAliveChan: tt.fields.waitUntilAlive,
				ConnLock:           tt.fields.connLock,
			}
			if got := n.GetEncoder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeighbour_GetID(t *testing.T) {
	type fields struct {
		Neighbour      util.Neighbour
		encoder        *codec.BinaryEncoder
		waitUntilAlive chan int
		connLock       *sync.RWMutex
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
			n := &remote.Neighbour{
				Neighbour:          tt.fields.Neighbour,
				Encoder:            tt.fields.encoder,
				WaitUntilAliveChan: tt.fields.waitUntilAlive,
				ConnLock:           tt.fields.connLock,
			}
			if got := n.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeighbour_GetPK(t *testing.T) {
	type fields struct {
		Neighbour      util.Neighbour
		encoder        *codec.BinaryEncoder
		waitUntilAlive chan int
		connLock       *sync.RWMutex
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
			n := &remote.Neighbour{
				Neighbour:          tt.fields.Neighbour,
				Encoder:            tt.fields.encoder,
				WaitUntilAliveChan: tt.fields.waitUntilAlive,
				ConnLock:           tt.fields.connLock,
			}
			if got := n.GetPK(); got != tt.want {
				t.Errorf("GetPK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeighbour_GetType(t *testing.T) {
	type fields struct {
		Neighbour      util.Neighbour
		encoder        *codec.BinaryEncoder
		waitUntilAlive chan int
		connLock       *sync.RWMutex
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
			n := &remote.Neighbour{
				Neighbour:          tt.fields.Neighbour,
				Encoder:            tt.fields.encoder,
				WaitUntilAliveChan: tt.fields.waitUntilAlive,
				ConnLock:           tt.fields.connLock,
			}
			if got := n.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeighbour_IsAlive(t *testing.T) {
	type fields struct {
		Neighbour      util.Neighbour
		encoder        *codec.BinaryEncoder
		waitUntilAlive chan int
		connLock       *sync.RWMutex
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
			n := &remote.Neighbour{
				Neighbour:          tt.fields.Neighbour,
				Encoder:            tt.fields.encoder,
				WaitUntilAliveChan: tt.fields.waitUntilAlive,
				ConnLock:           tt.fields.connLock,
			}
			if got := n.IsAlive(); got != tt.want {
				t.Errorf("IsAlive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeighbour_MarshalAndSend(t *testing.T) {
	type fields struct {
		Neighbour      util.Neighbour
		encoder        *codec.BinaryEncoder
		waitUntilAlive chan int
		connLock       *sync.RWMutex
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
			n := &remote.Neighbour{
				Neighbour:          tt.fields.Neighbour,
				Encoder:            tt.fields.encoder,
				WaitUntilAliveChan: tt.fields.waitUntilAlive,
				ConnLock:           tt.fields.connLock,
			}
			n.MarshalAndSend(tt.args.message)
		})
	}
}

func TestNeighbour_Send(t *testing.T) {
	type fields struct {
		Neighbour      util.Neighbour
		encoder        *codec.BinaryEncoder
		waitUntilAlive chan int
		connLock       *sync.RWMutex
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
			n := &remote.Neighbour{
				Neighbour:          tt.fields.Neighbour,
				Encoder:            tt.fields.encoder,
				WaitUntilAliveChan: tt.fields.waitUntilAlive,
				ConnLock:           tt.fields.connLock,
			}
			n.Send(tt.args.buf)
		})
	}
}

func TestNeighbour_SetConnAndEncoderAndSignalAlive(t *testing.T) {
	type fields struct {
		Neighbour      util.Neighbour
		encoder        *codec.BinaryEncoder
		waitUntilAlive chan int
		connLock       *sync.RWMutex
	}
	type args struct {
		conn *net.TCPConn
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
			n := &remote.Neighbour{
				Neighbour:          tt.fields.Neighbour,
				Encoder:            tt.fields.encoder,
				WaitUntilAliveChan: tt.fields.waitUntilAlive,
				ConnLock:           tt.fields.connLock,
			}
			n.SetConnAndEncoderAndSignalAlive(tt.args.conn)
		})
	}
}

func TestNeighbour_WaitUntilAlive(t *testing.T) {
	type fields struct {
		Neighbour      util.Neighbour
		encoder        *codec.BinaryEncoder
		waitUntilAlive chan int
		connLock       *sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &remote.Neighbour{
				Neighbour:          tt.fields.Neighbour,
				Encoder:            tt.fields.encoder,
				WaitUntilAliveChan: tt.fields.waitUntilAlive,
				ConnLock:           tt.fields.connLock,
			}
			n.WaitUntilAlive()
		})
	}
}

func TestNewNeighbour(t *testing.T) {
	type args struct {
		id      string
		address string
		pk      string
	}
	tests := []struct {
		name string
		args args
		want *remote.Neighbour
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := remote.NewNeighbour(tt.args.id, tt.args.address, tt.args.pk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNeighbour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNode(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want *remote.Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := remote.NewNode(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_GetAddress(t *testing.T) {
	type fields struct {
		address        string
		encoder        *codec.BinaryEncoder
		waitUntilAlive chan int
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
			n := &remote.Node{
				Address:            tt.fields.address,
				Encoder:            tt.fields.encoder,
				WaitUntilAliveChan: tt.fields.waitUntilAlive,
			}
			if got := n.GetAddress(); got != tt.want {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_GetEncoder(t *testing.T) {
	type fields struct {
		address        string
		encoder        *codec.BinaryEncoder
		waitUntilAlive chan int
	}
	tests := []struct {
		name   string
		fields fields
		want   codec.Encoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &remote.Node{
				Address:            tt.fields.address,
				Encoder:            tt.fields.encoder,
				WaitUntilAliveChan: tt.fields.waitUntilAlive,
			}
			if got := n.GetEncoder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_GetType(t *testing.T) {
	type fields struct {
		address        string
		encoder        *codec.BinaryEncoder
		waitUntilAlive chan int
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
			n := &remote.Node{
				Address:            tt.fields.address,
				Encoder:            tt.fields.encoder,
				WaitUntilAliveChan: tt.fields.waitUntilAlive,
			}
			if got := n.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_IsAlive(t *testing.T) {
	type fields struct {
		address        string
		encoder        *codec.BinaryEncoder
		waitUntilAlive chan int
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
			n := &remote.Node{
				Address:            tt.fields.address,
				Encoder:            tt.fields.encoder,
				WaitUntilAliveChan: tt.fields.waitUntilAlive,
			}
			if got := n.IsAlive(); got != tt.want {
				t.Errorf("IsAlive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_SetConnAndEncoderAndSignalAlive(t *testing.T) {
	type fields struct {
		address        string
		encoder        *codec.BinaryEncoder
		waitUntilAlive chan int
	}
	type args struct {
		conn *net.TCPConn
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
			n := &remote.Node{
				Address:            tt.fields.address,
				Encoder:            tt.fields.encoder,
				WaitUntilAliveChan: tt.fields.waitUntilAlive,
			}
			n.SetConnAndEncoderAndSignalAlive(tt.args.conn)
		})
	}
}

func TestNode_WaitUntilAlive(t *testing.T) {
	type fields struct {
		address        string
		encoder        *codec.BinaryEncoder
		waitUntilAlive chan int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &remote.Node{
				Address:            tt.fields.address,
				Encoder:            tt.fields.encoder,
				WaitUntilAliveChan: tt.fields.waitUntilAlive,
			}
			n.WaitUntilAlive()
		})
	}
}

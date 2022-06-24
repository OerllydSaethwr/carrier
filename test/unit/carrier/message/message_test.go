package message

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"reflect"
	"testing"
)

func TestBinaryMarshalH(t *testing.T) {
	type args struct {
		H string
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := message.BinaryMarshalH(tt.args.H)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryMarshalH() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BinaryMarshalH() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBinaryMarshalS(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := message.BinaryMarshalS(tt.args.S)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryMarshalS() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BinaryMarshalS() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBinaryMarshalSenderID(t *testing.T) {
	type args struct {
		SenderID string
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := message.BinaryMarshalSenderID(tt.args.SenderID)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryMarshalSenderID() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BinaryMarshalSenderID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBinaryMarshalV(t *testing.T) {
	type args struct {
		V [][]byte
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := message.BinaryMarshalV(tt.args.V)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryMarshalV() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BinaryMarshalV() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBinaryUnmarshal(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name  string
		args  args
		want  message.Type
		want1 message.Message
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := message.BinaryUnmarshal(tt.args.buf)
			if got != tt.want {
				t.Errorf("BinaryUnmarshal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BinaryUnmarshal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEchoMessage_BinaryMarshal(t *testing.T) {
	type fields struct {
		H string
		S util.Signature
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message.EchoMessage{
				H: tt.fields.H,
				S: tt.fields.S,
			}
			if got := msg.BinaryMarshal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryMarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEchoMessage_BinaryUnmarshal(t *testing.T) {
	type fields struct {
		H string
		S util.Signature
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
			msg := &message.EchoMessage{
				H: tt.fields.H,
				S: tt.fields.S,
			}
			msg.BinaryUnmarshal(tt.args.buf)
		})
	}
}

func TestEchoMessage_GetSenderID(t *testing.T) {
	type fields struct {
		H string
		S util.Signature
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
			msg := &message.EchoMessage{
				H: tt.fields.H,
				S: tt.fields.S,
			}
			if got := msg.GetSenderID(); got != tt.want {
				t.Errorf("GetSenderID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEchoMessage_GetType(t *testing.T) {
	type fields struct {
		H string
		S util.Signature
	}
	tests := []struct {
		name   string
		fields fields
		want   message.Type
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message.EchoMessage{
				H: tt.fields.H,
				S: tt.fields.S,
			}
			if got := msg.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEchoMessage_Hash(t *testing.T) {
	type fields struct {
		H string
		S util.Signature
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
			msg := &message.EchoMessage{
				H: tt.fields.H,
				S: tt.fields.S,
			}
			if got := msg.Hash(); got != tt.want {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEchoMessage_Marshal(t *testing.T) {
	type fields struct {
		H string
		S util.Signature
	}
	tests := []struct {
		name   string
		fields fields
		want   *message.TransportMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message.EchoMessage{
				H: tt.fields.H,
				S: tt.fields.S,
			}
			if got := msg.Marshal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEchoMessage_Payload(t *testing.T) {
	type fields struct {
		H string
		S util.Signature
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message.EchoMessage{
				H: tt.fields.H,
				S: tt.fields.S,
			}
			if got := msg.Payload(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Payload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitMessage_BinaryMarshal(t *testing.T) {
	type fields struct {
		V        [][]byte
		SenderID string
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message.InitMessage{
				V:        tt.fields.V,
				SenderID: tt.fields.SenderID,
			}
			if got := msg.BinaryMarshal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryMarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitMessage_BinaryUnmarshal(t *testing.T) {
	type fields struct {
		V        [][]byte
		SenderID string
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
			msg := &message.InitMessage{
				V:        tt.fields.V,
				SenderID: tt.fields.SenderID,
			}
			msg.BinaryUnmarshal(tt.args.buf)
		})
	}
}

func TestInitMessage_GetSenderID(t *testing.T) {
	type fields struct {
		V        [][]byte
		SenderID string
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
			msg := &message.InitMessage{
				V:        tt.fields.V,
				SenderID: tt.fields.SenderID,
			}
			if got := msg.GetSenderID(); got != tt.want {
				t.Errorf("GetSenderID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitMessage_GetType(t *testing.T) {
	type fields struct {
		V        [][]byte
		SenderID string
	}
	tests := []struct {
		name   string
		fields fields
		want   message.Type
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message.InitMessage{
				V:        tt.fields.V,
				SenderID: tt.fields.SenderID,
			}
			if got := msg.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitMessage_Hash(t *testing.T) {
	type fields struct {
		V        [][]byte
		SenderID string
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
			msg := &message.InitMessage{
				V:        tt.fields.V,
				SenderID: tt.fields.SenderID,
			}
			if got := msg.Hash(); got != tt.want {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitMessage_Marshal(t *testing.T) {
	type fields struct {
		V        [][]byte
		SenderID string
	}
	tests := []struct {
		name   string
		fields fields
		want   *message.TransportMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message.InitMessage{
				V:        tt.fields.V,
				SenderID: tt.fields.SenderID,
			}
			if got := msg.Marshal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitMessage_Payload(t *testing.T) {
	type fields struct {
		V        [][]byte
		SenderID string
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message.InitMessage{
				V:        tt.fields.V,
				SenderID: tt.fields.SenderID,
			}
			if got := msg.Payload(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Payload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEchoMessage(t *testing.T) {
	type args struct {
		h string
		s util.Signature
	}
	tests := []struct {
		name string
		args args
		want *message.EchoMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := message.NewEchoMessage(tt.args.h, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEchoMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInitMessage(t *testing.T) {
	type args struct {
		v        [][]byte
		senderID string
	}
	tests := []struct {
		name string
		args args
		want *message.InitMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := message.NewInitMessage(tt.args.v, tt.args.senderID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInitMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRequestMessage(t *testing.T) {
	type args struct {
		h        string
		senderID string
	}
	tests := []struct {
		name string
		args args
		want *message.RequestMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := message.NewRequestMessage(tt.args.h, tt.args.senderID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRequestMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewResolveMessage(t *testing.T) {
	type args struct {
		h        string
		v        [][]byte
		senderID string
	}
	tests := []struct {
		name string
		args args
		want *message.ResolveMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := message.NewResolveMessage(tt.args.h, tt.args.v, tt.args.senderID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResolveMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestMessage_BinaryMarshal(t *testing.T) {
	type fields struct {
		H        string
		SenderID string
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message.RequestMessage{
				H:        tt.fields.H,
				SenderID: tt.fields.SenderID,
			}
			if got := msg.BinaryMarshal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryMarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestMessage_BinaryUnmarshal(t *testing.T) {
	type fields struct {
		H        string
		SenderID string
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
			msg := &message.RequestMessage{
				H:        tt.fields.H,
				SenderID: tt.fields.SenderID,
			}
			msg.BinaryUnmarshal(tt.args.buf)
		})
	}
}

func TestRequestMessage_GetSenderID(t *testing.T) {
	type fields struct {
		H        string
		SenderID string
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
			msg := &message.RequestMessage{
				H:        tt.fields.H,
				SenderID: tt.fields.SenderID,
			}
			if got := msg.GetSenderID(); got != tt.want {
				t.Errorf("GetSenderID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestMessage_GetType(t *testing.T) {
	type fields struct {
		H        string
		SenderID string
	}
	tests := []struct {
		name   string
		fields fields
		want   message.Type
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message.RequestMessage{
				H:        tt.fields.H,
				SenderID: tt.fields.SenderID,
			}
			if got := msg.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestMessage_Hash(t *testing.T) {
	type fields struct {
		H        string
		SenderID string
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
			msg := &message.RequestMessage{
				H:        tt.fields.H,
				SenderID: tt.fields.SenderID,
			}
			if got := msg.Hash(); got != tt.want {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestMessage_Marshal(t *testing.T) {
	type fields struct {
		H        string
		SenderID string
	}
	tests := []struct {
		name   string
		fields fields
		want   *message.TransportMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message.RequestMessage{
				H:        tt.fields.H,
				SenderID: tt.fields.SenderID,
			}
			if got := msg.Marshal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestMessage_Payload(t *testing.T) {
	type fields struct {
		H        string
		SenderID string
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message.RequestMessage{
				H:        tt.fields.H,
				SenderID: tt.fields.SenderID,
			}
			if got := msg.Payload(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Payload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolveMessage_BinaryMarshal(t *testing.T) {
	type fields struct {
		H        string
		V        [][]byte
		SenderID string
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message.ResolveMessage{
				H:        tt.fields.H,
				V:        tt.fields.V,
				SenderID: tt.fields.SenderID,
			}
			if got := msg.BinaryMarshal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryMarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolveMessage_BinaryUnmarshal(t *testing.T) {
	type fields struct {
		H        string
		V        [][]byte
		SenderID string
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
			msg := &message.ResolveMessage{
				H:        tt.fields.H,
				V:        tt.fields.V,
				SenderID: tt.fields.SenderID,
			}
			msg.BinaryUnmarshal(tt.args.buf)
		})
	}
}

func TestResolveMessage_GetSenderID(t *testing.T) {
	type fields struct {
		H        string
		V        [][]byte
		SenderID string
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
			msg := &message.ResolveMessage{
				H:        tt.fields.H,
				V:        tt.fields.V,
				SenderID: tt.fields.SenderID,
			}
			if got := msg.GetSenderID(); got != tt.want {
				t.Errorf("GetSenderID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolveMessage_GetType(t *testing.T) {
	type fields struct {
		H        string
		V        [][]byte
		SenderID string
	}
	tests := []struct {
		name   string
		fields fields
		want   message.Type
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message.ResolveMessage{
				H:        tt.fields.H,
				V:        tt.fields.V,
				SenderID: tt.fields.SenderID,
			}
			if got := msg.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolveMessage_Hash(t *testing.T) {
	type fields struct {
		H        string
		V        [][]byte
		SenderID string
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
			msg := &message.ResolveMessage{
				H:        tt.fields.H,
				V:        tt.fields.V,
				SenderID: tt.fields.SenderID,
			}
			if got := msg.Hash(); got != tt.want {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolveMessage_Marshal(t *testing.T) {
	type fields struct {
		H        string
		V        [][]byte
		SenderID string
	}
	tests := []struct {
		name   string
		fields fields
		want   *message.TransportMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message.ResolveMessage{
				H:        tt.fields.H,
				V:        tt.fields.V,
				SenderID: tt.fields.SenderID,
			}
			if got := msg.Marshal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolveMessage_Payload(t *testing.T) {
	type fields struct {
		H        string
		V        [][]byte
		SenderID string
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message.ResolveMessage{
				H:        tt.fields.H,
				V:        tt.fields.V,
				SenderID: tt.fields.SenderID,
			}
			if got := msg.Payload(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Payload() = %v, want %v", got, tt.want)
			}
		})
	}
}

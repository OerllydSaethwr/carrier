package message

import (
	"encoding/json"
	"github.com/OerllydSaethwr/carrier/pkg/util"
)

type Type string

const (
	Init    Type = "init"
	Echo    Type = "echo"
	Request Type = "request"
	Resolve Type = "resolve"
)

type TransportMessage struct {
	Type    Type            `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type Message interface {
	// Payload @Deprecated
	Payload() []byte
	// Only use on InitMessage!
	Hash() string
	// Marshal @Deprecated
	Marshal() *TransportMessage
	GetSenderID() string
	GetType() Type
	BinaryMarshal() []byte
	BinaryUnmarshal(buf []byte)
}

func BinaryUnmarshal(buf []byte) (Type, Message) {
	types := []Type{Init, Echo, Request, Resolve}
	msgs := []Message{&InitMessage{}, &EchoMessage{}, &RequestMessage{}, &ResolveMessage{}}

	// 1. (Actual) First byte conveys the type information
	t := types[buf[0]]
	msg := msgs[buf[0]]

	// 2. Send the rest of the buffer for the concrete implementation to unmarshal
	msg.BinaryUnmarshal(buf[1:])

	return t, msg
}

func BinaryMarshalSenderID(SenderID string) ([]byte, []byte) {
	// 2. Put size of SenderID as 8 bytes - uint64
	senderIDb := []byte(SenderID)
	senderIDsize := util.MarshalUInt64(uint64(len(senderIDb)))
	return senderIDsize, senderIDb
}

func BinaryMarshalH(H string) ([]byte, []byte) {
	// 3. Put size of H as 8 bytes - uint64
	Hb := []byte(H)
	Hsize := util.MarshalUInt64(uint64(len(Hb)))
	return Hsize, Hb
}

func BinaryMarshalS(S string) ([]byte, []byte) {
	// 3. Put size of H as 8 bytes - uint64
	Sb := []byte(S)
	Ssize := util.MarshalUInt64(uint64(len(Sb)))
	return Ssize, Sb
}

func BinaryMarshalV(V [][]byte) ([]byte, []byte) {
	var lenSumCounter int
	for _, e := range V {
		lenSumCounter += len(e)
	}
	lenSum := util.MarshalUInt64(uint64(lenSumCounter))

	buf := make([]byte, 0)
	for _, e := range V {
		buf = append(buf, e...)
	}

	return lenSum, buf
}

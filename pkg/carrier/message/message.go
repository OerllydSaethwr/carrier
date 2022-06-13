package message

import (
	"encoding/json"
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
	Payload() []byte
	Hash() string
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

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
}

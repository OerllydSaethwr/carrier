package carrier

import "encoding/json"

type MessageType string

const (
	Init    MessageType = "init"
	Echo    MessageType = "echo"
	Request MessageType = "request"
	Resolve MessageType = "resolve"
)

//type Message interface {
//	GetType() MessageType
//}

type Message struct {
	MessageType MessageType     `json:"type"`
	Payload     json.RawMessage `json:"payload"`
}

type V struct {
	V [][]byte `json:"v"`
}

type H struct {
	H []byte `json:"h"`
}

type S struct {
	S []byte `json:"s"`
}

type InitMessage V

type RequestMessage H

type EchoMessage struct {
	H
	S
}

type ResolveMessage struct {
	H
	V
}

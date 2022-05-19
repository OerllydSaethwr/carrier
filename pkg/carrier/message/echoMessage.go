package message

type EchoMessage struct {
	H []byte `json:"h"`
	S []byte `json:"s"`
}

func (msg *EchoMessage) Payload() []byte {
	return nil
}

func (msg *EchoMessage) Hash() []byte {
	return nil
}

func (msg *EchoMessage) Marshal() *TransportMessage {
	return nil
}

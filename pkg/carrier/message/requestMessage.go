package message

type RequestMessage struct {
	H []byte `json:"h"`
}

func (msg *RequestMessage) Payload() []byte {
	return nil
}

func (msg *RequestMessage) Hash() []byte {
	return nil
}

func (msg *RequestMessage) Marshal() *TransportMessage {
	return nil
}

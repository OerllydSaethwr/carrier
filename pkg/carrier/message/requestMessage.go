package message

type RequestMessage struct {
	H []byte `json:"h"`
}

func (msg *RequestMessage) Payload() []byte {
	panic("not implemented")
	return nil
}

func (msg *RequestMessage) Hash() []byte {
	panic("not implemented")
	return nil
}

func (msg *RequestMessage) Marshal() *TransportMessage {
	panic("not implemented")
	return nil
}

func (msg *RequestMessage) GetSender() string {
	panic("not implemented")
	return ""
}

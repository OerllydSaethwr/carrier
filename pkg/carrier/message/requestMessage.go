package message

type RequestMessage struct {
	H string `json:"h"`
}

func (msg *RequestMessage) Payload() []byte {
	panic("not implemented")
	return nil
}

func (msg *RequestMessage) Hash() string {
	panic("not implemented")
	return ""
}

func (msg *RequestMessage) Marshal() *TransportMessage {
	panic("not implemented")
	return nil
}

func (msg *RequestMessage) GetSender() string {
	panic("not implemented")
	return ""
}

package message

type RequestMessage struct {
	H      string `json:"h"`
	Sender string `json:"sender"`
}

func NewRequestMessage(h string, sender string) *RequestMessage {
	return &RequestMessage{
		H:      h,
		Sender: sender,
	}
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
	return msg.Sender
}

func (msg *RequestMessage) GetType() Type {
	return Request
}

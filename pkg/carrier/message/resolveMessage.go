package message

type ResolveMessage struct {
	H []byte   `json:"h"`
	V [][]byte `json:"v"`
}

func (msg *ResolveMessage) Payload() []byte {
	return nil
}

func (msg *ResolveMessage) Hash() []byte {
	return nil
}

func (msg *ResolveMessage) Marshal() *TransportMessage {
	return nil
}

func (msg *ResolveMessage) GetSender() string {
	panic("not implemented")
	return ""
}

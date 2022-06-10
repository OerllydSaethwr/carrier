package message

import "encoding/json"

type ResolveMessage struct {
	H string   `json:"h"`
	V [][]byte `json:"v"`
}

func NewResolveMessage(h string, v [][]byte) *ResolveMessage {
	return &ResolveMessage{
		H: h,
		V: v,
	}
}

func (msg *ResolveMessage) Payload() []byte {
	panic("not implemented")
	return nil
}

func (msg *ResolveMessage) Hash() string {
	panic("not implemented")
	return ""
}

func (msg *ResolveMessage) Marshal() *TransportMessage {
	transportMessage := &TransportMessage{Type: msg.GetType()}
	payload, err := json.Marshal(msg)
	transportMessage.Payload = payload
	if err != nil {
		//log.Error().Msgf(err.Error()) //TODO Don't panic
		panic(err)
	}
	return transportMessage
}

func (msg *ResolveMessage) GetSenderID() string {
	panic("not implemented")
	return ""
}

func (msg *ResolveMessage) GetType() Type {
	return Resolve
}

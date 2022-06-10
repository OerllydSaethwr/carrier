package message

import "encoding/json"

type RequestMessage struct {
	H        string `json:"h"`
	SenderID string `json:"senderID"`
}

func NewRequestMessage(h string, senderID string) *RequestMessage {
	return &RequestMessage{
		H:        h,
		SenderID: senderID,
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
	transportMessage := &TransportMessage{Type: msg.GetType()}
	payload, err := json.Marshal(msg)
	transportMessage.Payload = payload
	if err != nil {
		//log.Error().Msgf(err.Error()) //TODO Don't panic
		panic(err)
	}
	return transportMessage
}

func (msg *RequestMessage) GetSenderID() string {
	return msg.SenderID
}

func (msg *RequestMessage) GetType() Type {
	return Request
}

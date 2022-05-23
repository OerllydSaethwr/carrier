package message

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type InitMessage struct {
	V      [][]byte `json:"v"`
	Sender string   `json:"sender"`
}

func NewInitMessage(v [][]byte, sender string) *InitMessage {
	return &InitMessage{
		V:      v,
		Sender: sender,
	}
}

func (msg *InitMessage) Payload() []byte {
	return bytes.Join(msg.V, nil)
}

func (msg *InitMessage) Hash() string {
	h := sha256.Sum256(msg.Payload())
	return hex.EncodeToString(h[:])
}

func (msg *InitMessage) Marshal() *TransportMessage {
	transportMessage := &TransportMessage{Type: Init}
	payload, err := json.Marshal(msg)
	transportMessage.Payload = payload
	if err != nil {
		//log.Error().Msgf(err.Error()) //TODO Don't panic
		panic(err)
	}
	return transportMessage
}

func (msg *InitMessage) GetSender() string {
	return msg.Sender
}

func (msg *InitMessage) GetType() Type {
	return Init
}

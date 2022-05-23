package message

import (
	"encoding/json"
	"github.com/OerllydSaethwr/carrier/pkg/util"
)

type EchoMessage struct {
	H      string         `json:"h"`
	S      util.Signature `json:"s"`
	Sender string         `json:"sender"`
}

func (msg *EchoMessage) Payload() []byte {
	panic("not implemented")
	return nil
}

func (msg *EchoMessage) Hash() string {
	panic("not implemented")
	return ""
}

func (msg *EchoMessage) Marshal() *TransportMessage {
	transportMessage := &TransportMessage{Type: Echo}
	payload, err := json.Marshal(msg)
	transportMessage.Payload = payload
	if err != nil {
		//log.Error().Msgf(err.Error()) //TODO Don't panic
		panic(err)
	}
	return transportMessage
}

func (msg *EchoMessage) GetSender() string {
	return msg.Sender
}

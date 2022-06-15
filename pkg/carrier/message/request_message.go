package message

import (
	"encoding/json"
	"github.com/OerllydSaethwr/carrier/pkg/util"
)

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

func (msg *RequestMessage) BinaryMarshal() []byte {
	buf := make([]byte, 0)

	// 1. Put message type as a single byte. RequestMessage has type 2
	buf = append(buf, byte(2))

	// 2. Put size of SenderID as 8 bytes - uint64
	senderIDb := []byte(msg.SenderID)
	senderIDsize := util.MarshalUInt64(uint64(len(senderIDb)))
	//println("i " + strconv.Itoa(len(senderIDb)))
	buf = append(buf, senderIDsize...)

	// 3. Put size of H as 8 bytes - uint64
	Hb := []byte(msg.H)
	Hsize := util.MarshalUInt64(uint64(len(Hb)))
	buf = append(buf, Hsize...)

	// 4. Put byte representation of SenderID
	buf = append(buf, senderIDb...)

	// 5. Put byte representation of H
	buf = append(buf, Hb...)

	return buf
}

func (msg *RequestMessage) BinaryUnmarshal(buf []byte) {
	// 2. Decode SenderIDsize - 8 bytes
	senderIDsize := util.UnmarshalUInt64(buf[:8])

	// 3. Decode size of H - 8 bytes
	Hsize := util.UnmarshalUInt64(buf[8:16])

	// 4. Decode SenderID - senderIDsize bytes
	senderIDb := buf[16 : 16+senderIDsize]
	msg.SenderID = string(senderIDb)

	// 5. Decode H - Hsize bytes
	Hb := buf[16+senderIDsize : 16+senderIDsize+Hsize]
	msg.H = string(Hb)
}

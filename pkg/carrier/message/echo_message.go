package message

import (
	"encoding/json"
	"github.com/OerllydSaethwr/carrier/pkg/util"
)

type EchoMessage struct {
	H string         `json:"h"`
	S util.Signature `json:"s"`
}

func NewEchoMessage(h string, s util.Signature) *EchoMessage {
	return &EchoMessage{
		H: h,
		S: s,
	}
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
	transportMessage := &TransportMessage{Type: msg.GetType()}
	payload, err := json.Marshal(msg)
	transportMessage.Payload = payload
	if err != nil {
		//log.Error().Msgf(err.Error()) //TODO Don't panic
		panic(err)
	}
	return transportMessage
}

func (msg *EchoMessage) GetSenderID() string {
	return msg.S.SenderID
}

func (msg *EchoMessage) GetType() Type {
	return Echo
}

func (msg *EchoMessage) BinaryMarshal() []byte {
	buf := make([]byte, 0)

	// 1. Put message type as a single byte. EchoMessage has type 1
	buf = append(buf, byte(1))

	// 2. Put size of SenderID as 8 bytes - uint64
	senderIDb := []byte(msg.S.SenderID)
	senderIDsize := util.MarshalUInt64(uint64(len(senderIDb)))
	//println("i " + strconv.Itoa(len(senderIDb)))
	buf = append(buf, senderIDsize...)

	// 3. Put size of H as 8 bytes - uint64
	Hb := []byte(msg.H)
	Hsize := util.MarshalUInt64(uint64(len(Hb)))
	buf = append(buf, Hsize...)

	// 4. Put size of S as 8 bytes - uint64
	Sb := []byte(msg.S.S)
	Ssize := util.MarshalUInt64(uint64(len(Sb)))
	buf = append(buf, Ssize...)

	// 5. Put byte representation of SenderID
	buf = append(buf, senderIDb...)

	// 6. Put byte representation of H
	buf = append(buf, Hb...)

	// 7. Put byte representation of S
	buf = append(buf, Sb...)

	return buf
}

func (msg *EchoMessage) BinaryUnmarshal(buf []byte) {
	// 2. Decode SenderIDsize - 8 bytes
	senderIDsize := util.UnmarshalUInt64(buf[:8])

	// 3. Decode size of H - 8 bytes
	Hsize := util.UnmarshalUInt64(buf[8:16])

	// 4. Decode size of S - 8 bytes
	Ssize := util.UnmarshalUInt64(buf[16:24])

	// 5. Decode SenderID - senderIDsize bytes
	senderIDb := buf[24 : 24+senderIDsize]
	msg.S.SenderID = string(senderIDb)

	// 6. Decode H - Hsize bytes
	Hb := buf[24+senderIDsize : 24+senderIDsize+Hsize]
	msg.H = string(Hb)

	// 7. Decode S - Ssize bytes
	Sb := buf[24+senderIDsize+Hsize : 24+senderIDsize+Hsize+Ssize]
	msg.S.S = string(Sb)
}

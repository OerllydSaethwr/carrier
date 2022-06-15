package message

import (
	"encoding/json"
	"github.com/OerllydSaethwr/carrier/pkg/util"
)

type ResolveMessage struct {
	H        string   `json:"h"`
	V        [][]byte `json:"v"`
	SenderID string   `json:"senderID"`
}

func NewResolveMessage(h string, v [][]byte, senderID string) *ResolveMessage {
	return &ResolveMessage{
		H:        h,
		V:        v,
		SenderID: senderID,
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

func (msg *ResolveMessage) BinaryMarshal() []byte {
	buf := make([]byte, 0)

	// 1. Put message type as a single byte. InitMessage has type 0
	buf = append(buf, byte(0))

	// 2. Put size of SenderID as 8 bytes - uint64
	senderIDb := []byte(msg.SenderID)
	senderIDsize := util.MarshalUInt64(uint64(len(senderIDb)))
	buf = append(buf, senderIDsize...)

	// 3. Put size of H as 8 bytes - uint64
	Hb := []byte(msg.H)
	Hsize := util.MarshalUInt64(uint64(len(Hb)))
	buf = append(buf, Hsize...)

	// 4. Put size of V as 8 bytes - uint64
	var lenSumCounter int
	for _, e := range msg.V {
		lenSumCounter += len(e)
	}
	lenSum := util.MarshalUInt64(uint64(lenSumCounter))
	buf = append(buf, lenSum...)

	// 5. Put byte representation of SenderID
	buf = append(buf, senderIDb...)

	// 6. Put byte representation of H
	buf = append(buf, Hb...)

	// 7. Put flattened byte array of V
	// Tsx is currently fixed so each element of V should have the same size. This will be crucial when decoding the struct.
	// @Critical C3 - if you change this condition, marshalling-unmarshalling will break
	for _, e := range msg.V {
		buf = append(buf, e...)
	}

	return buf
}

func (msg *ResolveMessage) BinaryUnmarshal(buf []byte) {
	// 2. Decode SenderIDsize - 8 bytes
	senderIDsize := util.UnmarshalUInt64(buf[:8])

	// 3. Decode size of H - 8 bytes
	Hsize := util.UnmarshalUInt64(buf[8:16])

	// 4. Decode size of V - 8 bytes
	lenSum := util.UnmarshalUInt64(buf[16:24])

	// 5. Decode SenderID - senderIDsize bytes
	senderIDb := buf[24 : 24+senderIDsize]
	msg.SenderID = string(senderIDb)

	// 6. Decode H - Hsize bytes
	Hb := buf[24+senderIDsize : 24+senderIDsize+Hsize]
	msg.H = string(Hb)

	// 7. Decode V - lenSum bytes
	Vb := buf[24+senderIDsize+Hsize : 24+senderIDsize+Hsize+lenSum]
	msg.V = util.Build(Vb)
}

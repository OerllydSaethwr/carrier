package message

import (
	"encoding/json"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
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
		log.Error().Msgf(err.Error())
	}
	return transportMessage
}

func (msg *ResolveMessage) GetSenderID() string {
	return msg.SenderID
}

func (msg *ResolveMessage) GetType() Type {
	return Resolve
}

func (msg *ResolveMessage) BinaryMarshal() []byte {
	buf := make([]byte, 0)

	// 1. Put message type as a single byte. InitMessage has type 0
	buf = append(buf, byte(0))

	// 2. Put tsxSize as 4 bytes - uint32
	tsxSizeb := util.MarshalUInt32(0)
	if len(msg.V) > 0 {
		tsxSizeb = util.MarshalUInt32(uint32(len(msg.V[0])))
	}
	buf = append(buf, tsxSizeb...)

	// 3. Put size of SenderID as 8 bytes - uint64
	senderIDb := []byte(msg.SenderID)
	senderIDsize := util.MarshalUInt64(uint64(len(senderIDb)))
	buf = append(buf, senderIDsize...)

	// 4. Put size of H as 8 bytes - uint64
	Hb := []byte(msg.H)
	Hsize := util.MarshalUInt64(uint64(len(Hb)))
	buf = append(buf, Hsize...)

	// 5. Put size of V as 8 bytes - uint64
	var lenSumCounter int
	for _, e := range msg.V {
		lenSumCounter += len(e)
	}
	lenSum := util.MarshalUInt64(uint64(lenSumCounter))
	buf = append(buf, lenSum...)

	// 6. Put byte representation of SenderID
	buf = append(buf, senderIDb...)

	// 7. Put byte representation of H
	buf = append(buf, Hb...)

	// 8. Put flattened byte array of V
	// Tsx is currently fixed so each element of V should have the same size. This will be crucial when decoding the struct.
	// @Critical C3 - if you change this condition, marshalling-unmarshalling will break
	for _, e := range msg.V {
		buf = append(buf, e...)
	}

	return buf
}

func (msg *ResolveMessage) BinaryUnmarshal(buf []byte) {
	// 2. Decode tsxSize - 4 bytes
	curr := 0
	next := 4
	tsxSize := util.UnmarshalUInt32(buf[curr:next])

	// 3. Decode SenderIDsize - 8 bytes
	curr = next
	next += 8
	senderIDsize := util.UnmarshalUInt64(buf[curr:next])

	// 4. Decode size of H - 8 bytes
	curr = next
	next += 8
	Hsize := util.UnmarshalUInt64(buf[curr:next])

	// 5. Decode size of V - 8 bytes
	curr = next
	next += 8
	lenSum := util.UnmarshalUInt64(buf[curr:next])

	// 6. Decode SenderID - senderIDsize bytes
	curr = next
	next += int(senderIDsize)
	senderIDb := buf[curr:next]
	msg.SenderID = string(senderIDb)

	// 7. Decode H - Hsize bytes
	curr = next
	next += int(Hsize)
	Hb := buf[curr:next]
	msg.H = string(Hb)

	// 8. Decode V - lenSum bytes
	curr = next
	next += int(lenSum)
	Vb := buf[curr:next]
	msg.V = util.Build(Vb, int(tsxSize))
}

package message

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
)

type InitMessage struct {
	V        [][]byte `json:"v"`
	SenderID string   `json:"senderID"`
}

func NewInitMessage(v [][]byte, senderID string) *InitMessage {
	return &InitMessage{
		V:        v,
		SenderID: senderID,
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
	transportMessage := &TransportMessage{Type: msg.GetType()}
	payload, err := json.Marshal(msg)
	transportMessage.Payload = payload
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	return transportMessage
}

func (msg *InitMessage) GetSenderID() string {
	return msg.SenderID
}

func (msg *InitMessage) GetType() Type {
	return Init
}

// BinaryMarshal Output will look like this (sizes in bytes):
// TotalSize - Type - IDSize - DataSize - ID - Data
//         8 -    1 -      8 -        8 -  IDSize - DataSize
func (msg *InitMessage) BinaryMarshal() []byte {
	buf := make([]byte, 0)

	// 1. Put message type as a single byte. InitMessage has type 0
	buf = append(buf, byte(0))

	// 2. Put tsxSize as 4 bytes - uint32
	tsxSizeb := util.MarshalUInt32(0)
	if len(msg.V) > 0 {

		// Tricky here - get the tsx size from the size of the first entry in V.
		// It would be nice to make this flexible (and also not too much effort)
		tsxSizeb = util.MarshalUInt32(uint32(len(msg.V[0])))
	}
	buf = append(buf, tsxSizeb...)

	// 3. Put size of SenderID as 8 bytes - uint64
	senderIDb := []byte(msg.SenderID)
	senderIDsize := util.MarshalUInt64(uint64(len(senderIDb)))
	//println("i " + strconv.Itoa(len(senderIDb)))
	buf = append(buf, senderIDsize...)

	// 4. Put size of V as 8 bytes - uint64
	var lenSumCounter int
	for _, e := range msg.V {
		lenSumCounter += len(e)
	}
	lenSum := util.MarshalUInt64(uint64(lenSumCounter))
	buf = append(buf, lenSum...)

	// 5. Put byte representation of SenderID
	buf = append(buf, senderIDb...)

	// 6. Put flattened byte array of V
	// Tsx is currently fixed so each element of V should have the same size. This will be crucial when decoding the struct.
	// @Critical C3 - if you change this condition, marshalling-unmarshalling will break
	for _, e := range msg.V {
		buf = append(buf, e...)
	}

	return buf
}

// BinaryUnmarshal should only be called by message.BinaryUnmarshal
func (msg *InitMessage) BinaryUnmarshal(buf []byte) {

	// 2. Decode tsxSize - 4 bytes
	curr := 0
	next := 4
	tsxSize := util.UnmarshalUInt32(buf[curr:next])

	// 2. Decode SenderIDsize - 8 bytes
	curr = next
	next += 8
	senderIDsize := util.UnmarshalUInt64(buf[curr:next])

	// 3. Decode size of V - 8 bytes
	curr = next
	next += 8
	lenSum := util.UnmarshalUInt64(buf[curr:next])

	// 4. Decode SenderID - senderIDsize bytes
	curr = next
	next += int(senderIDsize)
	senderIDb := buf[curr:next]
	msg.SenderID = string(senderIDb)

	// 5. Decode V - lenSum bytes
	curr = next
	next += int(lenSum)
	Vb := buf[curr:next]
	msg.V = util.Build(Vb, int(tsxSize))
}

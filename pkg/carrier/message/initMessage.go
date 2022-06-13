package message

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/OerllydSaethwr/carrier/pkg/util"
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
		//log.Error().Msgf(err.Error()) //TODO Don't panic
		panic(err)
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
	buf2 := make([]byte, 0)

	// 1. Put message type as a single byte. InitMessage has type 0
	buf2 = append(buf2, byte(0))

	// 2. Put size of SenderID as 8 bytes - uint64
	senderIDb := []byte(msg.SenderID)
	senderIDsize := util.MarshalUInt64(uint64(len(senderIDb)))
	//println("i " + strconv.Itoa(len(senderIDb)))
	buf2 = append(buf2, senderIDsize...)

	// 3. Put size of V as 8 bytes - uint64
	var lenSumCounter int
	for _, e := range msg.V {
		lenSumCounter += len(e)
	}
	lenSum := util.MarshalUInt64(uint64(lenSumCounter))
	//println("i " + strconv.Itoa(lenSumCounter))
	buf2 = append(buf2, lenSum...)

	// 4. Put byte representation of SenderID
	buf2 = append(buf2, senderIDb...)

	// 5. Put flattened byte array of V
	// Tsx is currently fixed so each element of V should have the same size. This will be crucial when decoding the struct.
	// @Critical C3 - if you change this condition, marshalling-unmarshalling will break
	for _, e := range msg.V {
		buf2 = append(buf2, e...)
	}

	// 0. Calculate the size of the whole struct, encode it as 8 bytes - uint64 and put it at the beginning of the whole buffer
	buf := make([]byte, 0)
	lenBuf2 := uint64(len(buf2))
	//println("i " + strconv.Itoa(int(lenBuf2)))
	buf = append(buf, util.MarshalUInt64(lenBuf2)...)
	buf = append(buf, buf2...)

	return buf
}

// BinaryUnmarshal should only be called by message.BinaryUnmarshal
func (msg *InitMessage) BinaryUnmarshal(buf []byte) {
	// 2. Decode SenderIDsize - 8 bytes
	senderIDsize := util.UnmarshalUInt64(buf[:8])

	// 3. Decode size of V - 8 bytes
	lenSum := util.UnmarshalUInt64(buf[8:16])

	// 4. Decode SenderID - senderIDsize bytes
	senderIDb := buf[16 : 16+senderIDsize]
	msg.SenderID = string(senderIDb)

	// 5. Decode V - lenSum bytes
	Vb := buf[16+senderIDsize : 16+senderIDsize+lenSum]
	msg.V = util.Build(Vb)
}

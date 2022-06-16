package carrier

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
	"io"
	"net"
	"sync"
)

type BinaryEncoder struct {
	conn net.Conn
	lock *sync.RWMutex
}

func NewBinaryEncoder(conn net.Conn) *BinaryEncoder {
	return &BinaryEncoder{
		conn: conn,
		lock: &sync.RWMutex{},
	}
}

func (be *BinaryEncoder) Encode(e any) error {

	var err error
	var toSend []byte

	switch data := e.(type) {
	case message.Message:
		toSend = data.BinaryMarshal()
	case *SuperBlockSummary:
		toSend = encodeSuperBlockSummary(data)
	case []byte:
		toSend = data
	default:
		panic("not implemented")
	}

	be.lock.Lock()
	_, err = be.conn.Write(util.Frame(toSend))
	be.lock.Unlock()

	return err
}

func encodeSuperBlockSummary(P *SuperBlockSummary) []byte {
	buf := make([]byte, 0)

	// 1. Length of map
	lnm := util.MarshalUInt32(uint32(len(*P)))
	buf = append(buf, lnm...)

	// 2. Iterate through map
	for h, sgnl := range *P {

		// 2.1. Size of Hash
		// 2.2. Hash
		Hb := []byte(h)
		Hsize := util.MarshalUInt32(uint32(len(Hb)))
		buf = append(buf, Hsize...)
		buf = append(buf, Hb...)

		// 2.3. Length of list
		lnl := util.MarshalUInt32(uint32(len(sgnl)))
		buf = append(buf, lnl...)

		// 2.4. Iterate through list
		for _, composite := range sgnl {

			// 2.4.1 Size of Signature
			// 2.4.2 Signature
			Sb := []byte(composite.S)
			Ssize := util.MarshalUInt32(uint32(len(Sb)))
			buf = append(buf, Ssize...)
			buf = append(buf, Sb...)

			// 2.4.3 Size of SenderID
			// 2.4.4 SenderID
			senderIDb := []byte(composite.SenderID)
			senderIDsize := util.MarshalUInt32(uint32(len(senderIDb)))
			buf = append(buf, senderIDsize...)
			buf = append(buf, senderIDb...)
		}
	}

	return buf
}

type BinaryDecoder struct {
	conn net.Conn
	lock *sync.RWMutex
}

func NewBinaryDecoder(conn net.Conn) *BinaryDecoder {
	return &BinaryDecoder{
		conn: conn,
		lock: &sync.RWMutex{},
	}
}

func (bd *BinaryDecoder) Decode(e any) error {
	var err error
	buf := make([]byte, 4)

	bd.lock.Lock()

	_, err = io.ReadFull(bd.conn, buf)
	if err != nil {
		log.Error().Msgf(err.Error())
		panic(err)
	}
	ls := util.UnmarshalUInt32(buf)

	buf2 := make([]byte, ls)
	_, err = io.ReadFull(bd.conn, buf2)
	if err != nil {
		log.Error().Msgf(err.Error())
		panic(err)
	}

	bd.lock.Unlock()

	switch data := e.(type) {

	/* I had to do some tricks here to avoid having to register types. The interface message.Message is always a pointer,
	but Go wouldn't let me dereference it, so I had to create a pointer to the interface so that I can change the
	value stored at the underlying address without returning anything, to respect the Decoder interface */
	case *message.Message:
		//TODO I'm not sure this pointer magic works
		_, m := message.BinaryUnmarshal(buf2)
		*data = m
	case *SuperBlockSummary:
		*data = decodeSuperBlockSummary(buf2)
	case []byte:
		panic("not implemented")
	default:
		panic("not implemented")
	}

	return err
}

// Very inconsistent when compared with message.BinaryUnmarshal
func decodeSuperBlockSummary(buf []byte) SuperBlockSummary {
	sbs := map[string][]util.Signature{}

	// 1. Length of map
	curr := uint32(0)
	next := uint32(4)
	lnm := util.UnmarshalUInt32(buf[curr:next])

	// 2. Map
	for i := uint32(0); i < lnm; i++ {

		// 2.1 Hash size
		// 2.2 Hash
		curr = next
		next += 4
		Hsize := util.UnmarshalUInt32(buf[curr:next])

		curr = next
		next = next + Hsize
		Hb := buf[curr:next]
		H := string(Hb)

		sbs[H] = make([]util.Signature, 0)

		// 2.3 Length of list
		curr = next
		next += 4
		lnl := util.UnmarshalUInt32(buf[curr:next])

		// 2.4 Iterate list
		for j := uint32(0); j < lnl; j++ {

			curr = next
			next += 4
			Ssize := util.UnmarshalUInt32(buf[curr:next])

			curr = next
			next += Ssize
			Sb := buf[curr:next]

			curr = next
			next += 4
			senderIDsize := util.UnmarshalUInt32(buf[curr:next])

			curr = next
			next += senderIDsize
			senderIDb := buf[curr:next]

			sbs[H] = append(sbs[H], util.Signature{
				S:        string(Sb),
				SenderID: string(senderIDb),
			})
		}
	}

	return sbs
}

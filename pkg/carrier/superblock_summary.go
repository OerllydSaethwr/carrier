package carrier

import "github.com/OerllydSaethwr/carrier/pkg/util"

func encodeSuperBlockSummary(P *SuperBlockSummary) []byte {
	buf := make([]byte, 0)

	// 1. ID - 4 bytes
	idb := util.MarshalUInt32(P.id)
	buf = append(buf, idb...)

	// 2. Length of map
	lnm := util.MarshalUInt32(uint32(len(P.payload)))
	buf = append(buf, lnm...)

	// 3. Iterate through map
	for h, sgnl := range P.payload {

		// 3.1. Size of Hash
		// 3.2. Hash
		Hb := []byte(h)
		Hsize := util.MarshalUInt32(uint32(len(Hb)))
		buf = append(buf, Hsize...)
		buf = append(buf, Hb...)

		// 3.3. Length of list
		lnl := util.MarshalUInt32(uint32(len(sgnl)))
		buf = append(buf, lnl...)

		// 3.4. Iterate through list
		for _, composite := range sgnl {

			// 3.4.1 Size of Signature
			// 3.4.2 Signature
			Sb := []byte(composite.S)
			Ssize := util.MarshalUInt32(uint32(len(Sb)))
			buf = append(buf, Ssize...)
			buf = append(buf, Sb...)

			// 3.4.3 Size of SenderID
			// 3.4.4 SenderID
			senderIDb := []byte(composite.SenderID)
			senderIDsize := util.MarshalUInt32(uint32(len(senderIDb)))
			buf = append(buf, senderIDsize...)
			buf = append(buf, senderIDb...)
		}
	}

	return buf
}

// Very inconsistent when compared with message.BinaryUnmarshal
func decodeSuperBlockSummary(buf []byte) SuperBlockSummary {
	sbs := SuperBlockSummary{
		id:      0,
		payload: map[string][]util.Signature{},
	}

	// 1. ID - 4 bytes
	curr := uint32(0)
	next := uint32(4)
	sbs.id = util.UnmarshalUInt32(buf[curr:next])

	// 2. Length of map
	curr = next
	next += 4
	lnm := util.UnmarshalUInt32(buf[curr:next])

	// 3. Map
	for i := uint32(0); i < lnm; i++ {

		// 3.1 Hash size
		// 3.2 Hash
		curr = next
		next += 4
		Hsize := util.UnmarshalUInt32(buf[curr:next])

		curr = next
		next = next + Hsize
		Hb := buf[curr:next]
		H := string(Hb)

		sbs.payload[H] = make([]util.Signature, 0)

		// 3.3 Length of list
		curr = next
		next += 4
		lnl := util.UnmarshalUInt32(buf[curr:next])

		// 3.4 Iterate list
		for j := uint32(0); j < lnl; j++ {

			// 3.4.1
			curr = next
			next += 4
			Ssize := util.UnmarshalUInt32(buf[curr:next])

			// 3.4.2
			curr = next
			next += Ssize
			Sb := buf[curr:next]

			// 3.4.3
			curr = next
			next += 4
			senderIDsize := util.UnmarshalUInt32(buf[curr:next])

			// 3.4.4
			curr = next
			next += senderIDsize
			senderIDb := buf[curr:next]

			sbs.payload[H] = append(sbs.payload[H], util.Signature{
				S:        string(Sb),
				SenderID: string(senderIDb),
			})
		}
	}

	return sbs
}

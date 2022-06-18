package carrier

import "github.com/OerllydSaethwr/carrier/pkg/util"

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

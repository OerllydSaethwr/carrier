package codec

import (
	"fmt"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/superblock"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"io"
	"net"
	"sync"
)

type BinaryDecoder struct {
	Conn net.Conn
	Lock *sync.RWMutex
}

// NewBinaryDecoder Usually we will be initiated with a conn that's alive, but just as a failsafe
func NewBinaryDecoder(conn net.Conn) *BinaryDecoder {
	return &BinaryDecoder{
		Conn: conn,
		Lock: &sync.RWMutex{},
	}
}

// Decode takes a pointer!
func (bd *BinaryDecoder) Decode(e any) error {
	var err error

	// Read the first 4 bytes, decode them into a uint32 - this tells us the size of the message
	buf := make([]byte, 4)

	// Synchronize access to the connection
	bd.Lock.Lock()

	_, err = io.ReadFull(bd.Conn, buf)
	if err != nil {
		bd.Lock.Unlock()
		return err
	}
	ls := util.UnmarshalUInt32(buf)

	// Read ls number of bytes
	buf2 := make([]byte, ls)
	_, err = io.ReadFull(bd.Conn, buf2)
	if err != nil {
		bd.Lock.Unlock()
		return err
	}

	// Release lock as soon as we're done, instead of deferring. Decoding can be an expensive operation.
	bd.Lock.Unlock()

	switch data := e.(type) {

	/* I had to do some tricks here to avoid having to register types. The interface message.Message is always a pointer,
	but Go wouldn't let me dereference it, so I had to create a pointer to the interface so that I can change the
	value stored at the underlying address without returning anything, to respect the Decoder interface */
	case *message.Message:
		//TODO I'm not sure this pointer magic works
		_, m := message.BinaryUnmarshal(buf2)
		*data = m
	case *superblock.SuperBlockSummary:
		*data = superblock.DecodeSuperBlockSummary(buf2)
	case *[]byte:
		*data = buf2
	default:
		err = fmt.Errorf("decoding this type is not supported")
	}

	return err
}

package codec

import (
	"fmt"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/superblock"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"net"
	"sync"
)

type BinaryEncoder struct {
	Conn net.Conn
	lock *sync.RWMutex
}

func NewBinaryEncoder(conn net.Conn) *BinaryEncoder {
	return &BinaryEncoder{
		Conn: conn,
		lock: &sync.RWMutex{},
	}
}

func (be *BinaryEncoder) Encode(e any) error {

	var err error
	var toSend []byte

	switch data := e.(type) {
	case message.Message:
		toSend = data.BinaryMarshal()
	case *superblock.SuperBlockSummary:
		toSend = superblock.EncodeSuperBlockSummary(data)
	case []byte:
		toSend = data
	default:
		err = fmt.Errorf("encoding this type is not supported")
	}

	be.lock.Lock()
	_, err = be.Conn.Write(util.Frame(toSend))
	be.lock.Unlock()

	return err
}

package carrier

import (
	"fmt"
	"github.com/OerllydSaethwr/carrier/pkg/carrier/message"
	"github.com/OerllydSaethwr/carrier/pkg/util"
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
		err = fmt.Errorf("encoding this type is not supported")
	}

	be.lock.Lock()
	_, err = be.conn.Write(util.Frame(toSend))
	be.lock.Unlock()

	return err
}

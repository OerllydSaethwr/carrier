package codec

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier/codec"
	"net"
	"reflect"
	"sync"
	"testing"
)

func TestBinaryDecoder_Decode(t *testing.T) {
	type fields struct {
		conn net.Conn
		lock *sync.RWMutex
	}
	type args struct {
		e any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bd := &codec.BinaryDecoder{
				Conn: tt.fields.conn,
				Lock: tt.fields.lock,
			}
			if err := bd.Decode(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBinaryEncoder_Encode(t *testing.T) {
	type fields struct {
		Conn net.Conn
		lock *sync.RWMutex
	}
	type args struct {
		e any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			be := &codec.BinaryEncoder{
				Conn: tt.fields.Conn,
				Lock: tt.fields.lock,
			}
			if err := be.Encode(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewBinaryDecoder(t *testing.T) {
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name string
		args args
		want *codec.BinaryDecoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := codec.NewBinaryDecoder(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinaryDecoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBinaryEncoder(t *testing.T) {
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name string
		args args
		want *codec.BinaryEncoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := codec.NewBinaryEncoder(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinaryEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

package superblock

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier/superblock"
	"reflect"
	"testing"
)

func TestDecodeSuperBlockSummary(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name string
		args args
		want superblock.SuperBlockSummary
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superblock.DecodeSuperBlockSummary(tt.args.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeSuperBlockSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeSuperBlockSummary(t *testing.T) {
	type args struct {
		P *superblock.SuperBlockSummary
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superblock.EncodeSuperBlockSummary(tt.args.P); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeSuperBlockSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}

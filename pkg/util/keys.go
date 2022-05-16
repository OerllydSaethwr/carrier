package util

import (
	"encoding/hex"
	"encoding/json"
	"github.com/rs/xid"
	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/pairing"
	"go.dedis.ch/kyber/v4/sign/bdn"
	"io/ioutil"
	"os"
)

type KeyPair struct {
	Name string
	Sk   kyber.Scalar
	Pk   kyber.Point
}

type KeyPairString struct {
	Name string `json:"name"`
	Sk   string `json:"sk"`
	Pk   string `json:"pk"`
}

func NewKeyPair(sk kyber.Scalar, pk kyber.Point) *KeyPair {
	return &KeyPair{
		Name: xid.New().String(),
		Sk:   sk,
		Pk:   pk,
	}
}

func GenerateRandomKeypair() *KeyPair {
	suite := pairing.NewSuiteBn256()
	sk, pk := bdn.NewKeyPair(suite, suite.RandomStream())

	return NewKeyPair(sk, pk)
}

func (kp *KeyPair) Convert() (*KeyPairString, error) {
	kpstr := &KeyPairString{}
	kpstr.Name = xid.New().String()
	skstr, pkstr, err := EncodeBdnToString(kp.Sk, kp.Pk)
	if err != nil {
		return nil, err
	}
	kpstr.Sk = skstr
	kpstr.Pk = pkstr
	return kpstr, nil
}

func (kpstr *KeyPairString) Convert() (*KeyPair, error) {
	kp := &KeyPair{}
	kp.Name = xid.New().String()
	sk, pk, err := DecodeStringToBdn(kpstr.Sk, kpstr.Pk)
	if err != nil {
		return nil, err
	}
	kp.Sk = sk
	kp.Pk = pk
	return kp, nil
}

func EncodeBdnToString(sk kyber.Scalar, pk kyber.Point) (string, string, error) {
	skstr, err := EncodeBdnToStringSk(sk)
	if err != nil {
		return "", "", err
	}
	pkstr, err := EncodeBdnToStringPk(pk)

	return skstr, pkstr, nil
}

func EncodeBdnToStringSk(sk kyber.Scalar) (string, error) {
	skm, err := sk.MarshalBinary()
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(skm), nil
}

func EncodeBdnToStringPk(pk kyber.Point) (string, error) {
	pkm, err := pk.MarshalBinary()
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(pkm), nil
}

func DecodeStringToBdn(sk, pk string) (kyber.Scalar, kyber.Point, error) {
	skk, err := DecodeStringToBdnSK(sk)
	if err != nil {
		return nil, nil, err
	}
	pkk, err := DecodeStringToBdnPK(pk)

	return skk, pkk, err
}

func DecodeStringToBdnSK(sk string) (kyber.Scalar, error) {
	skbuf, err := hex.DecodeString(sk)
	if err != nil {
		return nil, err
	}

	// This depends on the keys being generated on the G2 curve. Might be a good idea to parametrize this.
	// CRITICAL C1
	skk := pairing.NewSuiteBn256().G2().Scalar()
	err = skk.UnmarshalBinary(skbuf)
	if err != nil {
		return nil, err
	}

	return skk, nil
}

func DecodeStringToBdnPK(pk string) (kyber.Point, error) {
	pkbuf, err := hex.DecodeString(pk)
	if err != nil {
		return nil, err
	}

	// CRITICAL C1
	pkk := pairing.NewSuiteBn256().G2().Point()
	err = pkk.UnmarshalBinary(pkbuf)
	if err != nil {
		return nil, err
	}

	return pkk, nil
}

func ReadKeypairFile(f string) (*KeyPair, error) {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}

	kpstr := &KeyPairString{}
	err = json.Unmarshal(data, kpstr)
	if err != nil {
		return nil, err
	}

	kp, err := kpstr.Convert()
	if err != nil {
		return nil, err
	}

	return kp, nil
}

func WriteKeypairFile(f string, kp *KeyPair) error {
	keyFile, err := os.Create(f)
	if err != nil {
		return err
	}

	kpstr, err := kp.Convert()
	if err != nil {
		return err
	}
	data, err := json.Marshal(kpstr)
	if err != nil {
		return err
	}

	_, err = keyFile.Write(data)
	if err != nil {
		return err
	}

	return keyFile.Close()
}

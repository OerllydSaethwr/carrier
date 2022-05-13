package util

import (
	"encoding/hex"
	"github.com/rs/zerolog/log"
	"go.dedis.ch/kyber/v4/pairing"
	"go.dedis.ch/kyber/v4/sign/bdn"
)

func NewKeyPair() (string, string) {
	suite := pairing.NewSuiteBn256()
	privateKey, publicKey := bdn.NewKeyPair(suite, suite.RandomStream())

	skm, err := privateKey.MarshalBinary()
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	pkm, err := publicKey.MarshalBinary()
	if err != nil {
		log.Error().Msgf(err.Error())
	}

	skstr := hex.EncodeToString(skm)
	pkstr := hex.EncodeToString(pkm)

	return skstr, pkstr
}

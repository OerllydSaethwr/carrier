package util

import (
	"encoding/json"
	"go.dedis.ch/kyber/v4"
	"io/ioutil"
)

func ReadCarriersFile(filename string) (map[string]kyber.Point, error) {
	// TODO move this to util

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	carriersJSON := &CarriersJSON{}
	err = json.Unmarshal(data, carriersJSON)
	if err != nil {
		return nil, err
	}

	carriers := map[string]kyber.Point{}
	for _, c := range *carriersJSON {
		pkk, err := DecodeStringToBdnPK(c.Pk)
		if err != nil {
			return nil, err
		}
		carriers[c.Address] = pkk
	}

	return carriers, nil
}

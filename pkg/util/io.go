package util

import (
	"encoding/json"
	"io/ioutil"
)

func ReadCarriersFile(filename string) (map[string]string, error) {
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

	carriers := map[string]string{}
	for _, c := range *carriersJSON {
		//pkk, err := DecodeStringToBdnPK(c.Pk)
		if err != nil {
			return nil, err
		}
		carriers[c.Pk] = c.Address
	}

	return carriers, nil
}

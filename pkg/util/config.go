package util

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ID         string        `json:"id"`
	Addresses  Addresses     `json:"addresses"`
	Keys       KeyPairString `json:"keys"`
	Neighbours []Neighbour   `json:"neighbours"`
	Settings   Settings      `json:"settings"`
}

type Addresses struct {
	Client   string `json:"client"`
	Carrier  string `json:"carrier"`
	Front    string `json:"front"`
	Decision string `json:"decision"`
}

type KeyPairString struct {
	Sk string `json:"sk"`
	Pk string `json:"pk"`
}

type Neighbour struct {
	ID      string `json:"id"`
	Address string `json:"address"`
	PK      string `json:"pk"`
}

type Params struct {
	Hosts    []string `json:"hosts"`
	Fronts   []string `json:"fronts"`
	Settings Settings `json:"settings"`
}

type Settings struct {
	TsxSize int `json:"tsx-size"`
	Rate    int `json:"rate"`
	Nodes   int `json:"nodes"`

	DecisionPort int `json:"decision-port"`
	ClientPort   int `json:"client-port"`
	CarrierPort  int `json:"carrier-port"`

	InitThreshold int    `json:"init-threshold"`
	ForwardMode   int    `json:"forward-mode"`
	LogLevel      string `json:"log-level"`

	CarrierConnRetryDelay int `json:"carrier-conn-retry-delay"`
	CarrierConnMaxRetry   int `json:"carrier-conn-max-retry"`
	NodeConnRetryDelay    int `json:"node-conn-retry-delay"`
	NodeConnMaxRetry      int `json:"node-conn-max-retry"`

	LocalBasePort  int `json:"local-base-port"`
	LocalFrontPort int `json:"local-front-port"`
}

func NewSettings() Settings {
	return Settings{
		TsxSize:               128,
		Rate:                  100000,
		Nodes:                 4,
		CarrierConnRetryDelay: 1000,
		CarrierConnMaxRetry:   0,
		NodeConnRetryDelay:    1000,
		NodeConnMaxRetry:      0,
		InitThreshold:         100,
		LocalBasePort:         6000,
		LocalFrontPort:        9000,
		DecisionPort:          3000,
		CarrierPort:           4000,
		ClientPort:            5000,
		ForwardMode:           0,
		LogLevel:              "info",
	}
}

func LoadConfig(file string) (*Config, error) {
	// Read config file
	rawdata, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(rawdata, config)

	return config, err
}

func LoadParams(file string) (*Params, error) {
	rawdata, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	params := &Params{Settings: NewSettings()}
	err = json.Unmarshal(rawdata, params)

	return params, err
}

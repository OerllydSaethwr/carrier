package util

type Config struct {
	ID         string        `json:"id"`
	Addresses  Addresses     `json:"addresses"`
	Keys       KeyPairString `json:"keys"`
	Neighbours []Neighbour   `json:"neighbours"`
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

type Hosts struct {
	Hosts  []string `json:"hosts"`
	Fronts []string `json:"fronts"`
}

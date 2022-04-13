package main

import (
	"gitlab.epfl.ch/valaczka/carrier/pkg/carrier"
	"time"
)

func main() {
	c := carrier.NewCarrier()
	c.Start()

	for {
		time.Sleep(time.Second * 100000)
	}
}

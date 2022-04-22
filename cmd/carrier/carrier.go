package main

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier"
	"os"
	"strconv"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	front := os.Args[1]
	port, _ := strconv.Atoi(os.Args[2])

	c := carrier.NewCarrier(wg, front, port)
	c.Start()

	wg.Wait()
}

package main

import (
	"bytes"
	"fmt"
)

func main() {
	//TODO deprecated
	//wg := &sync.WaitGroup{}
	//front := os.Args[1]
	//port, _ := strconv.Atoi(os.Args[2])
	//
	//c := carrier.NewCarrier(wg, front, port)
	//c.Start()
	//
	//wg.Wait()

	e := []byte("hello world")
	data := make([][]byte, 0)
	for i := 0; i < i; i++ {
		data = append(data, e)
	}

	r := bytes.Join(data, nil)

	fmt.Println(string(r))
}

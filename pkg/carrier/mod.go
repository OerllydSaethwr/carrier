package carrier

import (
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
	"sync"
)

type Carrier struct {
	listener Listener

	quit  chan bool
	front string
	wg    *sync.WaitGroup

	secret string
}

func NewCarrier(wg *sync.WaitGroup, front string, port int) *Carrier {
	p := Carrier{
		quit:   make(chan bool, 1),
		secret: xid.New().String(), //TODO pass in
		front:  front,

		listener: &TCPListener{
			quit: make(chan bool, 1),
			name: "l",
			port: port, //TODO random number, change when size of testbed is known
		},
		wg: wg,
	}

	return &p
}

/*	Start listening to client requests
	Forward client requests
*/
func (c *Carrier) Start() {
	c.wg.Add(1)
	c.listener.Start()
}

func (c *Carrier) Stop() {
	log.Trace().Msgf("Stopping Carrier")

	c.listener.Stop()
	c.wg.Done()
}

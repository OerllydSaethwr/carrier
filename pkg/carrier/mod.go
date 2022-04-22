package carrier

import (
	"fmt"
	"github.com/OerllydSaethwr/carrier/pkg/util"
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

func NewCarrier(wg *sync.WaitGroup, front, mempool string) *Carrier {
	hostf, portf, err := util.SplitHostPort(front)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	hostm, portm, err := util.SplitHostPort(mempool)
	if err != nil {
		log.Error().Msgf(err.Error())
	}

	fmt.Println(hostf, portf, hostm) //TODO remove

	p := Carrier{
		quit:   make(chan bool, 1),
		secret: xid.New().String(), //TODO pass in
		front:  front,

		listener: &TCPListener{
			quit: make(chan bool, 1),
			name: "l",
			port: portm,
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

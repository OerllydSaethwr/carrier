package carrier

import (
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
	"gitlab.epfl.ch/valaczka/carrier/pkg/util"
)

type addr string

type Carrier struct {
	listener Listener

	quit          chan bool
	targetMempool addr

	secret string
}

func NewCarrier() *Carrier {
	p := Carrier{
		quit:          make(chan bool, 1),
		secret:        xid.New().String(), //TODO pass in
		targetMempool: "0.0.0.0",

		listener: &TCPListener{
			quit: make(chan bool, 1),
			name: "l",
			port: util.BASE_PORT + 200, //TODO random number, change when size of testbed is known
		},
	}

	return &p
}

/*	Start listening to client requests
	Forward client requests
*/
func (c *Carrier) Start() {
	c.listener.Start()
}

func (c *Carrier) Stop() {
	log.Trace().Msgf("Stopping Carrier")

	c.listener.Stop()
}

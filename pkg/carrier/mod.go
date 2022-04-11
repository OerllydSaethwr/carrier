package carrier

import (
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

type addr string

type Carrier struct {
	clientListener Listener
	nodeListener   Listener

	quit          chan bool
	targetMempool addr

	secret string
}

//  TODO: just forward everything to an address
func NewProxy() *Carrier {
	p := Carrier{
		quit:          make(chan bool, 1),
		secret:        xid.New().String(), //TODO
		targetMempool: "0.0.0.0",
	}

	return &p
}

/*	Start listening to client requests
	Forward client requests
*/
func (c *Carrier) Start() {
	c.clientListener.Start()
	c.nodeListener.Start()
}

func (c *Carrier) Stop() {
	log.Trace().Msgf("Stopping Carrier")

	c.clientListener.Stop()
	c.nodeListener.Stop()
}

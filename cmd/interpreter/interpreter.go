package main

import (
	"bufio"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.epfl.ch/valaczka/carrier/pkg/carrier"
	"os"
)

func main() {
	log.Error().Msgf("Starting Carrier")
	globalLevel := zerolog.TraceLevel
	var err error

	if 2 <= len(os.Args) {
		globalLevel, err = zerolog.ParseLevel(os.Args[1])
		if err != nil {
			log.Error().Msgf(err.Error())
		}
	}
	log.Error().Msgf("Logging level is set to " + globalLevel.String())
	zerolog.SetGlobalLevel(globalLevel)

	stdin := bufio.NewScanner(os.Stdin)
	nodes := make([]*carrier.Carrier, 0)

MAINLOOP:
	for stdin.Scan() {
		cmd := stdin.Text()

		switch cmd {
		case "set loglevel":
			log.Error().Msgf("Enter loglevel")
			stdin.Scan()
			level, err := zerolog.ParseLevel(stdin.Text())
			if err != nil {
				log.Error().Msgf(err.Error())
			}
			zerolog.SetGlobalLevel(level)
			log.Error().Msgf("Log level set to " + level.String())
		case "start":
			c := carrier.NewCarrier()
			c.Start()
			nodes = append(nodes, c)
			log.Error().Msgf("Carrier started")
		case "stop":
			for _, node := range nodes {
				node.Stop()
			}

			log.Error().Msgf("All nodes stopped")
		case "exit":
			fallthrough
		case "quit":
			log.Error().Msgf("Exiting")
			break MAINLOOP
		default:
			log.Error().Msgf("Unknown command")
		}
	}
}

package main

import (
	"github.com/rs/zerolog/log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const testPath = "test/"

func main() {
	_, err := exec.Command("rm", "-rf", testPath+"log").Output()
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	//log.Info().Msgf(string(output))
	_, err = exec.Command("mkdir", testPath+"log").Output()
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	//log.Info().Msgf(string(output))

	nodes := 4
	address := "127.0.0.1"
	colon := ":"
	port := 9000

	carriersFile := testPath + ".carriers.json"

	front := make([]string, 0)
	client := make([]string, 0)
	carrier := make([]string, 0)

	for i := 0; i < nodes; i++ {
		front = append(front, address+colon+strconv.Itoa(port+i+nodes))
		client = append(client, address+colon+strconv.Itoa(port+i+3*nodes))
		carrier = append(carrier, address+colon+strconv.Itoa(port+i+4*nodes))
	}

	logs := make([]*os.File, 0)
	for i := 0; i < nodes; i++ {
		// Create carrier logs
		logfile, err := os.Create(testPath + "log/.carrier-" + strconv.Itoa(i) + ".log")
		if err != nil {
			log.Error().Msgf(err.Error())
			break
		}
		defer logfile.Close()
		logs = append(logs, logfile)

		// Launch client
		if i == 0 {
			err = exec.Command("go", "run", testPath+"client/client.go", client[i]).Start()
			if err != nil {
				log.Error().Msgf(err.Error())
			}
		}

		// Launch carrier
		cmd := exec.Command("go", "run", "cmd/cobra/carrier.go", client[i], carrier[i], front[i], carriersFile, testPath+".carrier-0.json")
		cmd.Stdout = logfile
		cmd.Stderr = logfile
		err = cmd.Start()
		if err != nil {
			log.Error().Msgf(err.Error())
			break
		}

		// Launch node
		err = exec.Command("go", "run", testPath+"node/node.go", front[i]).Start()
		if err != nil {
			log.Error().Msgf(err.Error())
		}
	}

	time.Sleep(time.Second * 5)
}

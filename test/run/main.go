package main

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier"
	"github.com/OerllydSaethwr/carrier/pkg/util"
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
		return
	}
	//log.Info().Msgf(string(output))
	_, err = exec.Command("mkdir", testPath+"log").Output()
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}
	//log.Info().Msgf(string(output))

	nodes := 4
	address := "127.0.0.1"
	colon := ":"
	port := 9000

	carriersFile := testPath + ".carriers.json"
	keyPairFile := testPath + ".carrier-0.json"

	front := make([]string, 0)
	client := make([]string, 0)
	carrierA := make([]string, 0)

	for i := 0; i < nodes; i++ {
		front = append(front, address+colon+strconv.Itoa(port+i+nodes))
		client = append(client, address+colon+strconv.Itoa(port+i+3*nodes))
		carrierA = append(carrierA, address+colon+strconv.Itoa(port+i+4*nodes))
	}

	for i := 0; i < nodes; i++ {
		// Create client logs
		clientLog, err := os.Create(testPath + "log/.client-" + strconv.Itoa(i) + ".log")
		if err != nil {
			log.Error().Msgf(err.Error())
			return
		}
		defer clientLog.Close()

		// Create carrier logs
		carrierLog, err := os.Create(testPath + "log/.carrier-" + strconv.Itoa(i) + ".log")
		if err != nil {
			log.Error().Msgf(err.Error())
			return
		}
		defer carrierLog.Close()

		// Create node logs
		nodeLog, err := os.Create(testPath + "log/.node-" + strconv.Itoa(i) + ".log")
		if err != nil {
			log.Error().Msgf(err.Error())
			return
		}
		defer nodeLog.Close()

		// Launch client
		if i == 0 {
			cmd := exec.Command("go", "run", testPath+"client/client.go", client[i])
			cmd.Stdout = clientLog
			cmd.Stderr = clientLog
			err = cmd.Start()
			if err != nil {
				log.Error().Msgf(err.Error())
				return
			}
		}

		// Launch carrier
		if i != 0 {
			cmd := exec.Command("go", "run", "cmd/cobra/carrier.go", client[i], carrierA[i], front[i], carriersFile, testPath+".carrier-"+strconv.Itoa(i)+".json")
			cmd.Stdout = carrierLog
			cmd.Stderr = carrierLog
			err = cmd.Start()
			if err != nil {
				log.Error().Msgf(err.Error())
				return
			}
		}

		// Launch node
		cmd := exec.Command("go", "run", testPath+"node/node.go", front[i])
		cmd.Stdout = nodeLog
		cmd.Stderr = nodeLog
		err = cmd.Start()
		if err != nil {
			log.Error().Msgf(err.Error())
			return
		}
	}

	kp, err := util.ReadKeypairFile(keyPairFile)
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}

	carriers, err := util.ReadCarriersFile(carriersFile)
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}

	logfile, err := os.Create(testPath + "log/.carrier-0.log")
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}

	oldOut := os.Stdout
	oldErr := os.Stderr

	os.Stdout = logfile
	os.Stderr = logfile

	c := carrier.NewCarrier(client[0], carrierA[0], front[0], carriers, kp)
	wg := c.Start()

	wg.Wait()

	os.Stdout = oldOut
	os.Stderr = oldErr

	time.Sleep(time.Second * 5)
}

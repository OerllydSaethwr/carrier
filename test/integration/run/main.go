package main

import (
	"github.com/OerllydSaethwr/carrier/pkg/carrier"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	testPath := "test/"
	integrationPath := testPath + "integration/"
	filePath := testPath + "files/"
	scriptPath := "scripts/"
	configPath := filePath + "config/"

	nodes := 4
	frontPort := 9000
	tsxSize := 128
	rate := 5
	initThreshold := 10

	paramsFile := ".hosts-local.json"
	colon := ":"

	host := "127.0.0.1"
	portsPerCarrier := util.PortsPerCarrier

	zerolog.TimeFieldFormat = util.LogTimeFormat

	_, err := exec.Command("rm", "-rf", integrationPath+"log").Output()
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}
	//log.Info().Msgf(string(output))
	_, err = exec.Command("mkdir", integrationPath+"log").Output()
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}
	//log.Info().Msgf(string(output))

	cmd := exec.Command("python3", scriptPath+"generate-local-params.py", filePath+paramsFile, strconv.Itoa(nodes), strconv.Itoa(tsxSize), strconv.Itoa(rate), strconv.Itoa(initThreshold))
	_, err = cmd.Output()
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}

	params, err := util.LoadParams(filePath + paramsFile)
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logLevel, err := zerolog.ParseLevel(params.Settings.LogLevel)
	if err != nil {
		log.Warn().Msgf(err.Error())
	} else {
		zerolog.SetGlobalLevel(logLevel)
	}

	cmd = exec.Command("go", "run", "cmd/cobra/carrier.go", "generate", "config", filePath+paramsFile, configPath)
	_, err = cmd.Output()
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}

	front := make([]string, 0)
	client := make([]string, 0)
	decision := make([]string, 0)

	for i := 0; i < nodes; i++ {
		decision = append(decision, host+colon+strconv.Itoa(params.Settings.LocalBasePort+i*portsPerCarrier+1))
		client = append(client, host+colon+strconv.Itoa(params.Settings.LocalBasePort+i*portsPerCarrier+2))
		front = append(front, host+colon+strconv.Itoa(frontPort+i))
	}

	for i := 0; i < nodes; i++ {
		// Create client logs
		clientLog, err := os.Create(integrationPath + "log/.client-" + strconv.Itoa(i) + ".log")
		if err != nil {
			log.Error().Msgf(err.Error())
			return
		}
		defer clientLog.Close()

		// Create carrier logs
		carrierLog, err := os.Create(integrationPath + "log/.carrier-" + strconv.Itoa(i) + ".log")
		if err != nil {
			log.Error().Msgf(err.Error())
			return
		}
		defer carrierLog.Close()

		// Create node logs
		nodeLog, err := os.Create(integrationPath + "log/.node-" + strconv.Itoa(i) + ".log")
		if err != nil {
			log.Error().Msgf(err.Error())
			return
		}
		defer nodeLog.Close()

		// Launch client
		if i == 0 {
			cmd := exec.Command("go", "run", integrationPath+"client/client.go", client[i], strconv.Itoa(params.Settings.TsxSize), strconv.Itoa(params.Settings.Rate))
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
			cmd := exec.Command("go", "run", "cmd/cobra/carrier.go", configPath+".carrier-"+strconv.Itoa(i)+".json")
			cmd.Stdout = carrierLog
			cmd.Stderr = carrierLog
			err = cmd.Start()
			if err != nil {
				log.Error().Msgf(err.Error())
				return
			}
		}

		// Launch node
		cmd := exec.Command("go", "run", integrationPath+"node/node.go", front[i], decision[i])
		cmd.Stdout = nodeLog
		cmd.Stderr = nodeLog
		err = cmd.Start()
		if err != nil {
			log.Error().Msgf(err.Error())
			return
		}
	}

	logfile, err := os.Create(integrationPath + "log/.carrier-0.log")
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}

	oldOut := os.Stdout
	oldErr := os.Stderr

	os.Stdout = logfile
	os.Stderr = logfile

	config, err := util.LoadConfig(configPath + ".carrier-0.json")
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}

	c := carrier.NewCarrier(config)
	wg := c.Start()

	wg.Wait()

	os.Stdout = oldOut
	os.Stderr = oldErr

	time.Sleep(time.Second * 5)
}

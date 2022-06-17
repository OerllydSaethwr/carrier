/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strconv"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Generate configs",
	Long:  `Pass a hosts file, which we will read. Pass an output directory, where we will print all the new carrier configs.`,
	Args:  validateGenerateConfig,
	RunE:  runGenerateConfig,
}

func init() {
	generateCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func validateGenerateConfig(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: <hosts-file> <out-dir>")
	}

	return nil
}

func runGenerateConfig(cmd *cobra.Command, args []string) error {
	hostsFile := args[0]
	outDir := args[1]
	basePort := util.BasePort
	colon := ":"
	portsPerCarrier := util.PortsPerCarrier

	rawdata, err := ioutil.ReadFile(hostsFile)
	if err != nil {
		return err
	}

	hosts := &util.Hosts{}
	err = json.Unmarshal(rawdata, hosts)
	if err != nil {
		return err
	}

	configs := make([]util.Config, 0)
	neighbours := make([]util.Neighbour, 0)

	// We will build n=len(hosts.Hosts) configs
	for i, host := range hosts.Hosts {
		config := util.Config{}
		config.ID = strconv.Itoa(i)

		config.Addresses = util.Addresses{}
		config.Addresses.Front = hosts.Fronts[i]

		// If we're on localhost we need to shift ports for each carrier so that they are all unique
		shift := 0
		if isLocalHost(host) {
			shift = i * portsPerCarrier
			config.Addresses.Decision = host + colon + strconv.Itoa(basePort+1+shift)
			config.Addresses.Client = host + colon + strconv.Itoa(basePort+2+shift)
			config.Addresses.Carrier = host + colon + strconv.Itoa(basePort+3+shift)
		} else {
			config.Addresses.Decision = host + colon + strconv.Itoa(util.Decision)
			config.Addresses.Carrier = host + colon + strconv.Itoa(util.Carrier)
			config.Addresses.Client = host + colon + strconv.Itoa(util.Client)
		}

		kp, err := util.GenerateRandomKeypair()
		if err != nil {
			return err
		}
		config.Keys = *kp

		neighbour := util.Neighbour{
			ID:      config.ID,
			Address: config.Addresses.Carrier,
			PK:      config.Keys.Pk,
		}

		configs = append(configs, config)
		neighbours = append(neighbours, neighbour)
	}

	// TODO fix creation, clear if already exists
	//err = os.Mkdir(outDir, os.ModePerm)
	//if err != nil {
	//	return err
	//}

	err = os.Chdir(outDir)
	if err != nil {
		return err
	}

	for _, config := range configs {
		config.Neighbours = neighbours

		rawdata, err := json.Marshal(config)
		if err != nil {
			return err
		}

		f, err := os.Create(".carrier-" + config.ID + ".json")
		if err != nil {
			return err
		}

		_, err = f.Write(rawdata)
		if err != nil {
			return err
		}

		err = f.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func isLocalHost(host string) bool {
	return host == "localhost" || host == "127.0.0.1"
}

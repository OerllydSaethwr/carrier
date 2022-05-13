/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/OerllydSaethwr/carrier/pkg/carrier"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net"
	"os"
)

const (
	client2carrier  = 0
	carrier2carrier = 1
	front           = 2
	carriersFile    = 3
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra",
	Short: "Run carrier node",
	Long:  `nil`,
	Args:  validateCarrier,
	RunE:  runCarrier,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Error().Msgf(err.Error())
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func validateCarrier(cmd *cobra.Command, args []string) error {
	if len(args) < 4 {
		return fmt.Errorf("requires <client2carrier> <carrier2carrier> <front> <carriers_file>")
	}

	// Check IPs
	host, _, err := net.SplitHostPort(args[client2carrier])
	hostp := net.ParseIP(host)
	if err != nil || hostp == nil {
		return fmt.Errorf("<client2carrier> %s", err.Error())
	}
	host, _, err = net.SplitHostPort(args[carrier2carrier])
	hostp = net.ParseIP(host)
	if err != nil || hostp == nil {
		return fmt.Errorf("<carrier2carrier> %s", err.Error())
	}
	host, _, err = net.SplitHostPort(args[front])
	hostp = net.ParseIP(host)
	if err != nil || hostp == nil {
		return fmt.Errorf("<front> %s", err.Error())
	}

	// Check we can open carriersFile
	_, err = os.Stat(args[carriersFile])
	if err != nil {
		return fmt.Errorf("<carriers_file> %s", err.Error())
	}

	return nil
}

func runCarrier(cmd *cobra.Command, args []string) error {

	clientToCarrierAddr := args[client2carrier]
	carrierToCarrierAddr := args[carrier2carrier]
	frontAddr := args[front]

	carriersFile, err := os.Open(args[carriersFile])
	if err != nil {
		return err
	}

	byteValue, err := ioutil.ReadAll(carriersFile)
	var carriers carrier.CarrierAddrs
	err = json.Unmarshal(byteValue, &carriers)
	if err != nil {
		return err
	}

	c := carrier.NewCarrier(clientToCarrierAddr, carrierToCarrierAddr, frontAddr, carriers.CarrierAddrs)
	wg := c.Start()

	wg.Wait()

	return nil
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/OerllydSaethwr/carrier/pkg/carrier"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"net"
	"os"
)

const (
	client2carrier  = 0
	carrier2carrier = 1
	front           = 2
	decision        = 3
	carriersf       = 4
	keypairf        = 5
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
	if len(args) < 6 {
		return fmt.Errorf("requires <client2carrier> <carrier2carrier> <front> <decision> <carriers_file> <key_file>")
	}

	clientToCarrierAddr := args[client2carrier]
	carrierToCarrierAddr := args[carrier2carrier]
	frontAddr := args[front]
	decisionAddr := args[decision]
	carriersFile := args[carriersf]
	keyPairFile := args[keypairf]

	// Check IPs
	host, _, err := net.SplitHostPort(clientToCarrierAddr)
	hostp := net.ParseIP(host)
	if err != nil || hostp == nil {
		return fmt.Errorf("<client2carrier> %s", err.Error())
	}
	host, _, err = net.SplitHostPort(carrierToCarrierAddr)
	hostp = net.ParseIP(host)
	if err != nil || hostp == nil {
		return fmt.Errorf("<carrier2carrier> %s", err.Error())
	}
	host, _, err = net.SplitHostPort(frontAddr)
	hostp = net.ParseIP(host)
	if err != nil || hostp == nil {
		return fmt.Errorf("<front> %s", err.Error())
	}
	host, _, err = net.SplitHostPort(decisionAddr)
	hostp = net.ParseIP(host)
	if err != nil || hostp == nil {
		return fmt.Errorf("<decision> %s", err.Error())
	}

	// Check we can open carriersFile
	_, err = os.Stat(carriersFile)
	if err != nil {
		return fmt.Errorf("<carriers_file> %s", err.Error())
	}

	_, err = os.Stat(keyPairFile)
	if err != nil {
		return fmt.Errorf("<key_pair_file> %s", err.Error())
	}

	return nil
}

func runCarrier(cmd *cobra.Command, args []string) error {
	clientToCarrierAddr := args[client2carrier]
	carrierToCarrierAddr := args[carrier2carrier]
	frontAddr := args[front]
	decisionAddr := args[decision]
	carriersFile := args[carriersf]
	keyPairFile := args[keypairf]

	kp, err := util.ReadKeypairFile(keyPairFile)
	if err != nil {
		return err
	}

	carriers, err := util.ReadCarriersFile(carriersFile)
	if err != nil {
		return err
	}

	c := carrier.NewCarrier(clientToCarrierAddr, carrierToCarrierAddr, frontAddr, decisionAddr, carriers, kp)
	wg := c.Start()

	wg.Wait()

	return nil
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/OerllydSaethwr/carrier/pkg/carrier"
	"github.com/spf13/cobra"
	"net"
	"os"
	"sync"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application. For example:

			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	Args: validate,
	Run:  run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
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

func validate(cmd *cobra.Command, args []string) error {
	if len(args) < 3 {
		return errors.New("requires <client2carrier> <carrier2carrier> <front>")
	}
	host, _, err := net.SplitHostPort(args[0])
	hostp := net.ParseIP(host)
	if err != nil || hostp == nil {
		return fmt.Errorf("<client2carrier> %s", err)
	}
	hostp = net.ParseIP(host)
	if err != nil || hostp == nil {
		return fmt.Errorf("<carrier2carrier> %s", err)
	}
	host, _, err = net.SplitHostPort(args[1])
	hostp = net.ParseIP(host)
	if err != nil || hostp == nil {
		return fmt.Errorf("<front> %s", err)
	}
	return nil
}

func run(cmd *cobra.Command, args []string) {
	wg := &sync.WaitGroup{}
	clientToCarrierAddr := args[0]
	carrierToCarrierAddr := args[1]
	frontAddr := args[2]

	c := carrier.NewCarrier(wg, clientToCarrierAddr, carrierToCarrierAddr, frontAddr)
	c.Start()

	wg.Wait()
}

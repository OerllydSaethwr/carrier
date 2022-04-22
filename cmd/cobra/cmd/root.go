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
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires <target> <listener>")
		}
		host, _, err := net.SplitHostPort(args[0])
		hostp := net.ParseIP(host)
		if err != nil || hostp == nil {
			return fmt.Errorf("<front> %s", err)
		}
		host, _, err = net.SplitHostPort(args[1])
		hostp = net.ParseIP(host)
		if err != nil || hostp == nil {
			return fmt.Errorf("<mempool> %s", err)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		wg := &sync.WaitGroup{}
		front := args[0]
		mempool := args[1]

		c := carrier.NewCarrier(wg, front, mempool)
		c.Start()

		wg.Wait()
	},
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

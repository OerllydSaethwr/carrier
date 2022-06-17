/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/OerllydSaethwr/carrier/pkg/carrier"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra",
	Short: "Run carrier node",
	Long:  `Pass a config file. We will read the config file and start up a new carrier node.`,
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
	_, err := os.Stat(args[0])
	if err != nil {
		err = fmt.Errorf("<config_file> %s", err.Error())
	}

	return err
}

func runCarrier(cmd *cobra.Command, args []string) error {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	c, err := carrier.Load(args[0])
	if err != nil {
		return err
	}

	wg := c.Start()

	wg.Wait()

	return nil
}

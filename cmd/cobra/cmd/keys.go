/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/OerllydSaethwr/carrier/pkg/util"
	"github.com/rs/zerolog/log"
	"os"

	"github.com/spf13/cobra"
)

// keysCmd represents the generateKeys command
var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Generate key pair for signing",
	Long:  `Pass an output file where the newly generated keypair will be printed`,
	Args:  validateKeys,
	RunE:  runGenerateKeys,
}

func init() {
	generateCmd.AddCommand(keysCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// keysCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// keysCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func validateKeys(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("requires <output_file>")
	}

	// Check output file doesn't exist yet
	_, err := os.Stat(args[0])
	if err == nil {
		log.Warn().Msgf("output file already exists")
	}

	return nil
}

func runGenerateKeys(cmd *cobra.Command, args []string) error {
	kp, err := util.GenerateRandomKeypair()
	if err != nil {
		return err
	}
	return util.WriteKeypairFile(args[0], kp)
}

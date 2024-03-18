/*
Copyright © 2024 Joao Rodrigues jvfr96@gmail.com
*/

// Package cmd has all the existing command for this project
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "deck-api",
	Short: "Deck API is a simple REST API that handles card games deck handling",
	Long:  `Deck API is a simple REST API that handles card games deck handling`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

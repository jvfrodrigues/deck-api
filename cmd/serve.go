package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts up the Rest API server",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("serve called")
		err := start()
		return err
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

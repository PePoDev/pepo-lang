package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

const version string = "0.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the pepo-lang version information",
	Long:  "Show the pepo-lang version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Pepo Version: %s \n", version)
	},
}

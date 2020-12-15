package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "pepo",
	Long: "pepo-lang is a scripting language for holy shit performance and fucking syntax",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute init root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

package cmd

import "github.com/spf13/cobra"

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run api service",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

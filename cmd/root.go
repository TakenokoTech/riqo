package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "riqo",
	Short: "A CLI tool to assist with GitHub CLI commands",
	Long: `Riqo is a CLI tool designed to help users generate and manage GitHub CLI commands efficiently.
			It leverages local LLMs to provide intelligent suggestions and improve user productivity.`,
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

package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "notes",
	Short: "A simple CLI for managing notes",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

package cli

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "forge",
	Short: "Forge is a CLI tool for managing your projects",
	Long:  `Forge is a powerful command-line interface (CLI) tool that helps you manage and automate your projects with ease.`,
}

func Execute() error {
	return rootCmd.Execute()
}

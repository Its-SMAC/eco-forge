package cli

import (
	"forge/internal/project"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"new"},
	Short:   "Inicialize your project",
	Long:    `Use this command to create initial folders and files to use forge on your project`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return project.Init()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

package cmd

import (
	"project/common/cmd"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Short: "Project",
}

func Execute() {
	cmd.ExecuteRootCmd(&rootCmd)
}

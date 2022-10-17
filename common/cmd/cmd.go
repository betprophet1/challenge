package cmd

import (
	"fmt"
	"os"

	"project/common/rootcloser"

	"github.com/spf13/cobra"
)

func Initialize(initializers ...func()) {
	for _, initializer := range initializers {
		cobra.OnInitialize(initializer)
	}
}

func ExecuteRootCmd(cmd *cobra.Command) {
	defer rootcloser.Execute()
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

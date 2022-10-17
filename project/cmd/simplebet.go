package cmd

import (
	"project/common/cache/client"
	"project/common/database/orm"
	"project/project/services/simplebet"

	"github.com/spf13/cobra"
)

func simpleAppPreRun(cmd *cobra.Command, args []string) {
	orm.InitGorm()
	client.InitRedis()
}

var simplebetCmd = cobra.Command{
	Use:   "simplebet",
	Short: "Simple bet",
}

var simplebetHttpCmd = cobra.Command{
	Use:              "http",
	Short:            "Serve HTTP API service",
	PersistentPreRun: simpleAppPreRun,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetInt("port")
		simplebet.StartSimpleBet(host, port)
	},
}

var simplebetWorkerCmd = cobra.Command{
	Use:              "worker",
	Short:            "Start worker service",
	PersistentPreRun: simpleAppPreRun,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(&simplebetCmd)

	simplebetCmd.AddCommand(&simplebetHttpCmd)
	simplebetHttpCmd.Flags().String("host", "0.0.0.0", "Host to bind")
	simplebetHttpCmd.Flags().Int("port", 8080, "Port to bind")

	simplebetCmd.AddCommand(&simplebetWorkerCmd)
}

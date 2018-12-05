package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(confCmd)
}

var confCmd = &cobra.Command{
	Use:   "conf",
	Short: "Print the configuration",
	Long:  "Print the configuration of the Spider application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("core.count: %s\n", viper.GetString("core.count"))
		fmt.Printf("core.environment: %s\n", viper.GetString("core.environment"))
		fmt.Printf("core.seed: %s\n", viper.GetString("core.seed"))
		fmt.Printf("storage.adapter: %s\n", viper.GetString("storage.adapter"))
		fmt.Printf("storage.host: %s\n", viper.GetString("storage.host"))
		fmt.Printf("storage.port: %s\n", viper.GetString("storage.port"))
		fmt.Printf("storage.dbname: %s\n", viper.GetString("storage.dbname"))
		fmt.Printf("storage.user: %s\n", viper.GetString("storage.user"))
		fmt.Printf("storage.pass: %s\n", viper.GetString("storage.pass"))
	},
}

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
		fmt.Printf("core.environment: %s\n", viper.GetString("core.environment"))
		fmt.Printf("core.waitTime: %s\n", viper.GetString("core.waitTime"))
		fmt.Printf("db.adapter: %s\n", viper.GetString("db.adapter"))
		fmt.Printf("db.host: %s\n", viper.GetString("db.host"))
		fmt.Printf("db.port: %s\n", viper.GetString("db.port"))
		fmt.Printf("db.dbname: %s\n", viper.GetString("db.dbname"))
		fmt.Printf("db.user: %s\n", viper.GetString("db.user"))
		fmt.Printf("db.pass: %s\n", viper.GetString("db.pass"))
	},
}

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
		fmt.Printf("environment: %s\n", viper.GetString("core.environment"))
		files := viper.Get("core.files").([]interface{})
		fmt.Printf("files      : ")
		for i := 0; i < len(files); i++ {
			fmt.Printf("%s\n             ", files[i])
		}
	},
}

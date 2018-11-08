package cmd

import (
	"fmt"

	"github.com/donbstringham/cleaner/ver"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the version number of Spider application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Spider %s\n", ver.Version)
	},
}

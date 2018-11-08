package cmd

import (
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/spf13/jwalterweatherman"
)

var (
	cfgFile string
	Verbose bool
)

var RootCmd = &cobra.Command{
	Use:   "spider",
	Short: "Weber Spider",
	Long: `Weber Spider
        Final project for CS 4350.
	Here are few examples of CLI commands:
	$ spider help 		- Prints help.
	$ spider conf 		- Prints configuration information.
	$ spider version	- Prints current version.`,
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.spider.toml)")
	RootCmd.PersistentFlags().BoolVar(&Verbose, "verbose", false, "verbose output")
}

// initConfig reads in the config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.FATAL.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".spider")
	}
	viper.AutomaticEnv()
	viper.SetEnvPrefix("SPIDER")

	if err := viper.ReadInConfig(); err == nil {
		log.INFO.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// Execute adds all child commands to the root and sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.FATAL.Println(err)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thomaspoignant/api-scenario/pkg/log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)


var cfgFile string
var quiet bool
var noColor bool
var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "api-scenario",
	Short: "Scenario API testing from the command line.",
	Long: `API-scenario is a simple command line tool that allow you to execute easily a scenario to test your APIs.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.api-scenario.yaml)")
	runCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Run your scenario in quiet mode")
	runCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "Do not display color on the output")
	runCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Run your scenario in quiet mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".newApp" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".newApp")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// Disable logs
	if quiet {
		log.Logger.Level = logrus.ErrorLevel
	}

	if verbose {
		log.Logger.Level = logrus.TraceLevel
	}

	// Disable colors
	if noColor {
		log.Logger.Formatter = &log.OutputFormatter{
			DisableColors: true,
		}
	}

}

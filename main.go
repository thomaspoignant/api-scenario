package main

import (
	"github.com/spf13/viper"
	"github.com/thomaspoignant/api-scenario/cmd"
)

// VersionString is the current version of the cli. It is override im the make file during the build.
var version = "dev"
var commit = "none"
var date = "unknown"

func main() {
	viper.Set("version", version)
	viper.Set("commit", commit)
	viper.Set("date", date)
	cmd.Execute()
}

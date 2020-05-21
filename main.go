package main

import (
	"github.com/spf13/viper"
	"github.com/thomaspoignant/api-scenario/cmd"
)

// VersionString is the current version of the cli. It is override im the make file during the build.
var VersionString = "unset"

func main() {
	viper.Set("version", VersionString)
	cmd.Execute()
}

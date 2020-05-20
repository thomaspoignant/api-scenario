package main

import (
	"github.com/spf13/viper"
	"github.com/thomaspoignant/api-scenario/cmd"
)

var VersionString = "unset"

func main() {
	viper.Set("version", VersionString)
	cmd.Execute()
}
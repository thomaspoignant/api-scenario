package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thomaspoignant/api-scenario/pkg/config_helper"
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"github.com/thomaspoignant/api-scenario/pkg/model/context"
	"os"
	"strings"
)
var headers []string
var variables []string
var inputFile string
var token string
func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&inputFile, "scenario", "s", "", "Input file for the scenario.")
	runCmd.Flags().StringArrayVarP(&headers, "header", "H", []string{}, "Header you want to override (format should be \"header_name:value\").")
	runCmd.Flags().StringArrayVarP(&variables, "variable", "V", []string{}, "Value for a variable used in your scenario (format should be \"variable_name:value\").")
	runCmd.Flags().StringVarP(&token, "authorization-token", "t", "", "Authorization token send in the Authorization headers.")
	runCmd.MarkFlagRequired("scenario")
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute your scenario",
	Long:  `Execute your scenario`, //TODO: Change it for more details
	Run: func(cmd *cobra.Command, args []string) {
		// add variable to context
		addVariableToContext(variables)
		// format headers and add it to the config
		viper.Set("headers", formatHeadersForConfig(headers, token))

		// Parse the input file
		scenario, err := model.InitScenarioFromFile(inputFile)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}

		// run the scenario
		scenario.Run()
	},
}

func addVariableToContext(variables []string) {
	const separator  = ":"
	for _, variable := range variables {
		splitStr := strings.SplitN(variable, separator, 2)
		if len(splitStr) <= 1 {
			color.Red("Wrong format for parameter %s, it should be \"Key:value\", this parameter is ignored.", variable)
			break
		}
		context.GetContext().Add(strings.TrimSpace(splitStr[0]), strings.TrimSpace(splitStr[1]))
	}
}

func formatHeadersForConfig(headers []string, token string) map[string]string{
	const separator  = ":"
	var res = map[string] string{}
	for _, header := range headers {
		fmt.Println(header)
		splitStr := strings.SplitN(header, separator, 2)
		if len(splitStr) <= 1 {
			color.Red("Wrong format for parameter %s, it should be \"Key:value\", this parameter is ignored.", header)
			break
		}
		res[strings.TrimSpace(splitStr[0])] = strings.TrimSpace(splitStr[1])
	}

	// Authentication token
	if len(token) > 0 {
		res["Authorization"] = config_helper.FormatAuthorization(token)
	}
	return res
}
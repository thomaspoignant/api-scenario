package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"github.com/thomaspoignant/api-scenario/pkg/model/context"
	"github.com/thomaspoignant/api-scenario/pkg/util"
	"os"
	"strings"
)

var headers []string
var variables []string
var inputFile string
var token string

// init setup the flags used by the run command.
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
	Long:  `Execute your scenario`,
	Run: func(cmd *cobra.Command, args []string) {
		// add variable to context
		addVariableToContext(variables)

		// format headers and add it to the config
		viper.Set("headers", formatHeadersForConfig(headers, token))

		// Parse the input file
		scenario, err := model.InitScenarioFromFile(inputFile)
		if err != nil {
			logrus.Error(err)
			os.Exit(1)
		}

		// run the scenario
		res := scenario.Run()
		if !res.IsSuccess() {
			os.Exit(1)
		}
	},
}

// addVariableToContext is adding a variable to the context to replace wildcard strings
func addVariableToContext(variables []string) {
	const separator = ":"
	for _, variable := range variables {
		splitStr := strings.SplitN(variable, separator, 2)
		if len(splitStr) <= 1 {
			logrus.Errorf("Wrong format for parameter %s, it should be \"Key:value\", this parameter is ignored.", variable)
			break
		}
		context.GetContext().Add(strings.TrimSpace(splitStr[0]), strings.TrimSpace(splitStr[1]))
	}
}

// formatHeadersForConfig is formatting headers to put them in the config
func formatHeadersForConfig(headers []string, token string) map[string]string {
	const separator = ":"
	var res = map[string]string{}
	for _, header := range headers {
		splitStr := strings.SplitN(header, separator, 2)
		if len(splitStr) <= 1 {
			logrus.Errorf("Wrong format for parameter %s, it should be \"Key:value\", this parameter is ignored.", header)
			break
		}
		res[strings.TrimSpace(splitStr[0])] = strings.TrimSpace(splitStr[1])
	}

	// Authentication token
	if len(token) > 0 {
		res["Authorization"] = util.AddBearerPrefix(token)
	}
	return res
}

package cmd

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thomaspoignant/api-scenario/pkg/context"
	"github.com/thomaspoignant/api-scenario/pkg/controller"
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"github.com/thomaspoignant/api-scenario/pkg/util"
)

var headers []string
var variables []string
var inputFile string
var token string
var outputFile string
var outputFormat string

// init setup the flags used by the run command.
func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&inputFile, "scenario", "s", "", "Input file for the scenario.")
	runCmd.Flags().StringArrayVarP(&headers, "header", "H", []string{}, "Header you want to override (format should be \"header_name:value\").")
	runCmd.Flags().StringArrayVarP(&variables, "variable", "V", []string{}, "Value for a variable used in your scenario (format should be \"variable_name:value\").")
	runCmd.Flags().StringVarP(&token, "authorization-token", "t", "", "Authorization token send in the Authorization headers.")
	runCmd.Flags().StringVarP(&outputFile, "output-file", "f", "", "Output file where to save the result (use --output-format to specify if you want JSON or YAML output).")
	runCmd.Flags().StringVar(&outputFormat, "output-format", "JSON", "Format of the output file, available values are JSON and YAML (ignored if --output-file is not set).")
	if err := runCmd.MarkFlagRequired("scenario"); err != nil {
		panic(err)
	}
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
		check(err)

		// run the scenario
		ctrl, err := controller.InitializeScenarioController()
		check(err)

		res := ctrl.Run(scenario)
		saveResultInFile(res)
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
			continue
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
			continue
		}
		res[strings.TrimSpace(splitStr[0])] = strings.TrimSpace(splitStr[1])
	}

	// Authentication token
	if len(token) > 0 {
		res["Authorization"] = util.AddBearerPrefix(token)
	}
	return res
}

// save result in a file
func saveResultInFile(result model.ScenarioResult) {

	if len(outputFile) == 0 {
		return
	}

	var s []byte
	var err error
	// Marshal in the choosing format
	if strings.ToUpper(outputFormat) == "YAML" || strings.ToUpper(outputFormat) == "YML" {
		s, err = yaml.Marshal(result)
	} else {
		s, err = json.Marshal(result)
	}
	check(err)

	err = ioutil.WriteFile(outputFile, s, 0644)
	check(err)

	logrus.Infof("Output file %s saved", outputFile)
}

func check(err error) {
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
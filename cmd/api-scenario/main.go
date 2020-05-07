package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/thomaspoignant/api-scenario/pkg/config"
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"github.com/thomaspoignant/api-scenario/pkg/model/context"
	"io/ioutil"
	"log"
	"os"
	"strings"
)





// TODO :
// - ajouter la gestion de response size + headers

func main() {
	var headers config.CmdFlag
	flag.Var(&headers, "H", "Override headers.")
	var variables config.CmdFlag
	flag.Var(&variables, "V", "Value for variables used in your test file.")
	inputFile := flag.String("F", "", "Input file")
	flag.Parse()

	// Open file.
	file, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("Impossible to locate the file %s", *inputFile)
		os.Exit(1)
	}

	// Unmarshall file to launch the scenario.
	data := model.Scenario{}
	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Fatalf("Impossible to read file %s\nerror: %s", inputFile, err.Error())
		os.Exit(2)
	}

	// Add variables to context
	addToContext(variables)

	// Run the scenario.
	data.Run()
}


func addToContext(variables config.CmdFlag) {
	const separator  = ":"
	for _, variable := range variables {
		splitStr := strings.SplitN(variable, separator, 1)
		if len(splitStr) <= 1 {
			color.Red("Wrong format for parameter %s, it should be \"Key:value\", this parameter is ignored.", cmdParam)
			return
		}
		context.GetContext().Add(strings.TrimSpace(splitStr[0]), strings.TrimSpace(splitStr[1]))
	}
}

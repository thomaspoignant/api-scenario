package main

import (
	"encoding/json"
	"fmt"
	"github.com/thomaspoignant/api-scenario/pkg/config"
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"io/ioutil"
	"os"
)

// TODO :
// - ajouter la gestion de response size + headers

func main() {
	// Read command line options and put them in a config object
	_, err := config.InitConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %q", err)
		os.Exit(1)
	}

	file, _ := ioutil.ReadFile("testdata/test_1.json")
	data := model.Scenario{}
	_ = json.Unmarshal([]byte(file), &data)
	data.Run()
}

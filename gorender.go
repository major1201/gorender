package main

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/go-yaml/yaml"
	"github.com/major1201/goutils/logging"
	"github.com/urfave/cli"
	htmlTemplate "html/template"
	"os"
	"strings"
	textTemplate "text/template"
)

// AppVer means the project's version
const AppVer = "0.1.0"

var logger = logging.New("MAIN")

func parseArguments(c *cli.Context) map[string]interface{} {
	argMap := map[string]interface{}{}
	// 1. json
	if c.IsSet("json") {
		jsonFile, err := os.Open(c.String("json"))
		defer jsonFile.Close()
		if err != nil {
			logger.Fatal(err)
		}
		json.NewDecoder(jsonFile).Decode(&argMap)
	}
	// 2. yaml
	if c.IsSet("yaml") {
		yamlFile, err := os.Open(c.String("yaml"))
		defer yamlFile.Close()
		if err != nil {
			logger.Fatal(err)
		}
		yaml.NewDecoder(yamlFile).Decode(&argMap)
	}
	// 3. toml
	if c.IsSet("toml") {
		toml.DecodeFile(c.String("toml"), &argMap)
	}
	// 4. extra arguments overrides all above
	arguments := c.StringSlice("arguments")
	for _, argString := range arguments {
		arr := strings.SplitN(argString, "=", 2)
		argMap[arr[0]] = arr[1]
	}
	return argMap
}

func runApp(c *cli.Context) {
	// get arguments
	argMap := parseArguments(c)

	filename := c.Args().First()
	writer := os.Stdout
	if c.IsSet("output") {
		outputFile, err := os.Create(c.String("output"))
		defer outputFile.Close()
		if err != nil {
			logger.Fatal(err)
		}
		writer = outputFile
	}
	if c.Bool("in-place") {
		self, err := os.Create(filename)
		defer self.Close()
		if err != nil {
			logger.Fatal(err)
		}
		writer = self
	}
	if c.Bool("html") {
		t, err := htmlTemplate.ParseFiles(filename)
		if err != nil {
			logger.Fatal(err)
		}
		t.Execute(writer, argMap)
	} else {
		t, err := textTemplate.ParseFiles(filename)
		if err != nil {
			logger.Fatal(err)
		}
		t.Execute(writer, argMap)
	}
}

func main() {
	// init logging
	logging.AddStdout(0)

	// parse flags
	app := getApp()
	if err := app.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}

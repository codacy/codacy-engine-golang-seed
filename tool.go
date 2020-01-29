package codacytool

import (
	"encoding/json"
	"github.com/josemiguelmelo/gofile"
	logrus "github.com/sirupsen/logrus"
	"os"
)

// ToolImplementation interface to implement the tool
type ToolImplementation interface {
	Run(tool Tool, sourceDir string) ([]Issue, error)
}

// Tool represents a codacy tool
type Tool struct {
	Definition ToolDefinition
	config     Configuration
	Patterns   []Pattern
	Files      []string
}

func patternsFromConfig(toolName string, config Configuration) []Pattern {
	for _, t := range config.Tools {
		if t.Name == toolName {
			return t.Patterns
		}
	}
	return []Pattern{}
}

func newTool(toolDefinition ToolDefinition, config Configuration) Tool {
	return Tool{
		Definition: toolDefinition,
		config:     config,
		Files:      config.Files,
		Patterns:   patternsFromConfig(toolDefinition.Name, config),
	}
}

func defaultTool() Tool {
	toolDefinition, err := LoadToolDefinition(defaultDefinitionFile())
	if err != nil {
		panic(err)
	}
	config, err := ParseConfiguration(defaultConfigurationFile())
	if err != nil {
		logrus.Debug(defaultConfigurationFile() + " parsing error: " + err.Error())
	}

	return newTool(toolDefinition, config)
}

func appendResult(currentResultString, newResult string) string {
	if currentResultString != "" {
		currentResultString = currentResultString + "\n"
	}
	return currentResultString + newResult
}

func printResult(issues []Issue) {
	printResult := ""
	for _, i := range issues {
		iJSON, err := i.ToJSON()
		if err != nil {
			fileError := FileError{
				Filename: i.File,
				Message:  err.Error(),
			}

			fileErrorJSON, _ := fileError.ToJSON()
			printResult = appendResult(printResult, string(fileErrorJSON))
		} else {
			printResult = appendResult(printResult, string(iJSON))
		}
	}

	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&NoFormatter{})
	logrus.Info(printResult)
}

func startToolImplementation(impl ToolImplementation, sourceDir string) {
	tool := defaultTool()

	result, err := impl.Run(tool, sourceDir)
	if err != nil {
		logrus.Errorln(err.Error())
		os.Exit(1)
	}

	printResult(result)
}

// ToolDefinition is the configuration of the tool to run
type ToolDefinition struct {
	Name     string    `json:"name"`
	Version  string    `json:"version,omitempty"`
	Patterns []Pattern `json:"patterns"`
}

// ToJSON returns the json representation of the tool
func (i *ToolDefinition) ToJSON() ([]byte, error) {
	return json.Marshal(i)
}

// ToJSONBeautify returns the json representation of the tool
func (i *ToolDefinition) ToJSONBeautify() ([]byte, error) {
	return json.MarshalIndent(i, "", "  ")
}

// LoadToolDefinition loads tool information from documentation file
func LoadToolDefinition(fileLocation string) (ToolDefinition, error) {
	tool := ToolDefinition{}
	err := gofile.ParseJSONFile(fileLocation, &tool)
	return tool, err
}

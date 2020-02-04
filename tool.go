package codacytool

import (
	"encoding/json"
	logrus "github.com/sirupsen/logrus"
	"os"
	"strings"
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

func defaultTool(runConfig RunConfiguration) Tool {
	toolDefinition, err := LoadToolDefinition(defaultDefinitionFile(runConfig.toolConfigsBasePath))
	if err != nil {
		panic(err)
	}
	config, err := ParseConfiguration(defaultConfigurationFile(runConfig.toolConfigsBasePath))
	if err != nil {
		logrus.Debug(defaultConfigurationFile(runConfig.toolConfigsBasePath) + " parsing error: " + err.Error())
	}

	return newTool(toolDefinition, config)
}

func appendResult(currentResultString, newResult string) string {
	return currentResultString + newResult + "\n"
}

func resultAsString(issues []Issue) string {
	resultString := ""

	for _, i := range issues {
		iJSON, err := i.ToJSON()
		if err != nil {
			fileError := FileError{
				Filename: i.File,
				Message:  err.Error(),
			}

			fileErrorJSON, _ := fileError.ToJSON()
			resultString = appendResult(resultString, string(fileErrorJSON))
		} else {
			resultString = appendResult(resultString, string(iJSON))
		}
	}

	return strings.TrimSuffix(resultString, "\n")
}

func printResult(issues []Issue, err error) {
	if err != nil {
		logrus.Fatal(err.Error())
	}

	resultString := resultAsString(issues)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&NoFormatter{})
	logrus.Info(resultString)
}

func startToolImplementation(impl ToolImplementation, runConfiguration RunConfiguration) ([]Issue, error) {
	tool := defaultTool(runConfiguration)

	return impl.Run(tool, runConfiguration.sourceDir)
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

// LoadToolDefinition loads tool information from documentation file
func LoadToolDefinition(fileLocation string) (ToolDefinition, error) {
	tool := ToolDefinition{}
	jsonFileContent, err := readFile(fileLocation)
	if err != nil {
		return tool, err
	}
	err = json.Unmarshal(jsonFileContent, &tool)
	return tool, err
}

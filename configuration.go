package codacytool

import (
	"encoding/json"
)

// Configuration configuration provided to the tool to specify files and patterns to analyse
type Configuration struct {
	Files []string         `json:"files"`
	Tools []ToolDefinition `json:"tools"`
}

// ParseConfiguration parses configuration file
func ParseConfiguration(fileLocation string) (Configuration, error) {
	config := Configuration{}
	jsonFileContent, err := readFile(fileLocation)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(jsonFileContent, &config)
	return config, err
}

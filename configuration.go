package codacytool

import (
	"github.com/josemiguelmelo/gofile"
)

// Configuration configuration provided to the tool to specify files and patterns to analyse
type Configuration struct {
	Files []string `json:"files"`
	Tools []Tool   `json:"tools"`
}

// ParseConfiguration parses configuration file
func ParseConfiguration(fileLocation string) (Configuration, error) {
	config := Configuration{}
	err := gofile.ParseJSONFile(fileLocation, &config)
	return config, err
}

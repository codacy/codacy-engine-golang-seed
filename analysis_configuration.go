package codacytool

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

const defaultAnalysisConfigurationFile = ".codacyrc"

// AnalysisConfiguration represents the files to analyse and the tools to analyze them with, as obtained via the .codacyrc file.
type AnalysisConfiguration struct {
	Files *[]string         `json:"files"`
	Tools *[]ToolDefinition `json:"tools"`
}

// loadAnalysisConfiguration loads the analysis configuration from the default file.
// If the file does not exist, an empty AnalysisConfiguration is returned.
// Tools should know how to deal with the absence of these values.
// If the file exists but is invalid, the execution is aborted.
func loadAnalysisConfiguration(runConfiguration RunConfiguration) (AnalysisConfiguration, error) {
	fileLocation := filepath.Join(runConfiguration.ToolConfigurationDir, defaultAnalysisConfigurationFile)

	analysisConfiguration := AnalysisConfiguration{}

	fileContent, err := os.ReadFile(fileLocation)
	if err != nil {
		logrus.Infof("Failed to read analysis configuration file: %s\n%s", fileLocation, err.Error())
		return analysisConfiguration, nil
	}

	err = json.Unmarshal(fileContent, &analysisConfiguration)
	if err != nil {
		logrus.Infof("Failed to parse analysis configuration file content: %s\n%s", string(fileContent), err.Error())
	}
	return analysisConfiguration, err
}

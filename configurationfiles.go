package codacytool

import (
	"os"
)

var _toolFilesBasePath = "/"
var _defaultDefinitionFile = "docs/patterns.json"
var _defaultConfigurationFile = ".codacyrc"
var _basePathEnvVar = "TOOL_CONFIGS_BASEPATH"

func getBasePath() string {
	basePathFromEnv := os.Getenv(_basePathEnvVar)
	if basePathFromEnv != "" {
		return basePathFromEnv
	}
	return _toolFilesBasePath
}

func getPathToFile(file string) string {
	return getBasePath() + file
}

func defaultDefinitionFile() string {
	return getPathToFile(_defaultDefinitionFile)
}

func defaultConfigurationFile() string {
	return getPathToFile(_defaultConfigurationFile)
}

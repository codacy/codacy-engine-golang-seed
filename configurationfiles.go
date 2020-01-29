package codacytool

var _toolFilesBasePath = "/"
var _defaultDefinitionFile = "docs/patterns.json"
var _defaultConfigurationFile = ".codacyrc"
var _basePathEnvVar = "TOOL_CONFIGS_BASEPATH"

func getBasePath() string {
	var basePathFromEnv string
	if toolConfigsBasePathFlag != nil {
		basePathFromEnv = *toolConfigsBasePathFlag
	}
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

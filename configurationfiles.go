package codacytool

const (
	toolFilesBasePath         = "/"
	defaultDefinitionFileName = "docs/patterns.json"
	defaultConfigFileName     = ".codacyrc"
)

func getBasePath() string {
	var basePathFromEnv string
	if toolConfigsBasePathFlag != nil {
		basePathFromEnv = *toolConfigsBasePathFlag
	}
	if basePathFromEnv != "" {
		return basePathFromEnv
	}
	return toolFilesBasePath
}

func getPathToFile(file string) string {
	return getBasePath() + file
}

func defaultDefinitionFile() string {
	return getPathToFile(defaultDefinitionFileName)
}

func defaultConfigurationFile() string {
	return getPathToFile(defaultConfigFileName)
}

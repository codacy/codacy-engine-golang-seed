package codacytool

import (
	"path/filepath"
)

const (
	toolFilesBasePath         = "/"
	defaultDefinitionFileName = "docs/patterns.json"
	defaultConfigFileName     = ".codacyrc"
)

func getBasePath(toolConfigsBasePath string) string {
	if toolConfigsBasePath != "" {
		return toolConfigsBasePath
	}
	return toolFilesBasePath
}

func getPathToFile(toolConfigsBasePath string, file string) string {
	return filepath.Join(getBasePath(toolConfigsBasePath), file)
}

func defaultDefinitionFile(toolConfigsBasePath string) string {
	return getPathToFile(toolConfigsBasePath, defaultDefinitionFileName)
}

func defaultConfigurationFile(toolConfigsBasePath string) string {
	return getPathToFile(toolConfigsBasePath, defaultConfigFileName)
}

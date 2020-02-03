package codacytool

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBasePath(t *testing.T) {
	basePath := getBasePath()
	assert.Equal(t, toolFilesBasePath, basePath)
}

func TestGetPathToFile(t *testing.T) {
	file := "/file"
	expectedPath := getBasePath() + file
	path := getPathToFile(file)
	assert.Equal(t, expectedPath, path)
}

func TestDefaultDefinitionFile(t *testing.T) {
	expectedPath := getPathToFile(defaultDefinitionFileName)
	path := defaultDefinitionFile()
	assert.Equal(t, expectedPath, path)
}

func TestDefaultConfigurationFile(t *testing.T) {
	expectedPath := getPathToFile(defaultConfigFileName)
	path := defaultConfigurationFile()
	assert.Equal(t, expectedPath, path)
}

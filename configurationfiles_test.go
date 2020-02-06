package codacytool

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestGetBasePath(t *testing.T) {
	basePath := getBasePath(toolFilesBasePath)
	assert.Equal(t, toolFilesBasePath, basePath)
}

func TestGetPathToFile(t *testing.T) {
	file := "/file"
	expectedPath := filepath.Join(getBasePath(toolFilesBasePath), file)
	path := getPathToFile(toolFilesBasePath, file)
	assert.Equal(t, expectedPath, path)
}

func TestDefaultDefinitionFile(t *testing.T) {
	expectedPath := getPathToFile(toolFilesBasePath, defaultDefinitionFileName)
	path := defaultDefinitionFile(toolFilesBasePath)
	assert.Equal(t, expectedPath, path)
}

func TestDefaultConfigurationFile(t *testing.T) {
	expectedPath := getPathToFile(toolFilesBasePath, defaultConfigFileName)
	path := defaultConfigurationFile(toolFilesBasePath)
	assert.Equal(t, expectedPath, path)
}

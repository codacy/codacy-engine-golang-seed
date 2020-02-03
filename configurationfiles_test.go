package codacytool

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

const basePath = "/"

func TestGetBasePath(t *testing.T) {
	basePath := getBasePath(basePath)
	assert.Equal(t, toolFilesBasePath, basePath)
}

func TestGetPathToFile(t *testing.T) {
	file := "/file"
	expectedPath := filepath.Join(getBasePath(basePath), file)
	path := getPathToFile(basePath, file)
	assert.Equal(t, expectedPath, path)
}

func TestDefaultDefinitionFile(t *testing.T) {
	expectedPath := getPathToFile(basePath, defaultDefinitionFileName)
	path := defaultDefinitionFile(basePath)
	assert.Equal(t, expectedPath, path)
}

func TestDefaultConfigurationFile(t *testing.T) {
	expectedPath := getPathToFile(basePath, defaultConfigFileName)
	path := defaultConfigurationFile(basePath)
	assert.Equal(t, expectedPath, path)
}

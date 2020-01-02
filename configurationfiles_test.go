package codacytool

import (
	"testing"
)

func TestGetBasePath(t *testing.T) {
	basePath := getBasePath()
	if basePath != testsResourcesLocation {
		t.Errorf("Expected: %s; Got: %s", _toolFilesBasePath, basePath)
	}
}

func TestGetPathToFile(t *testing.T) {
	file := "/file"
	expectedPath := getBasePath() + file
	path := getPathToFile(file)
	if path != expectedPath {
		t.Errorf("Expected: %s; Got: %s", expectedPath, path)
	}
}

func TestDefaultDefinitionFile(t *testing.T) {
	expectedPath := getPathToFile(_defaultDefinitionFile)
	path := defaultDefinitionFile()
	if path != expectedPath {
		t.Errorf("Expected: %s; Got: %s", expectedPath, path)
	}
	return
}
func TestDefaultConfigurationFile(t *testing.T) {
	expectedPath := getPathToFile(_defaultConfigurationFile)
	path := defaultConfigurationFile()
	if path != expectedPath {
		t.Errorf("Expected: %s; Got: %s", expectedPath, path)
	}
	return
}

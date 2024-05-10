package codacytool

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToolExecution(t *testing.T) {
	// Arrange
	srcDir := "sourceDir"
	runConfig := RunConfiguration{
		ToolConfigurationDir: filepath.Join(testsResourcesLocation, "tool"),
		SourceDir:            srcDir,
	}

	expectedToolExecution := ToolExecution{
		ToolDefinition: ToolDefinition{
			Name:    "gorevive",
			Version: "1.0.0",
			Patterns: &[]Pattern{
				{
					ID:       "foo",
					Category: "ErrorProne",
					ScanType: "SAST",
					Level:    "Warning",
					Parameters: []PatternParameter{
						{
							Name:    "bar",
							Default: "nofunc",
						},
					},
				},
			},
		},
		Files: &[]string{
			"foo/bar/baz.go",
			"foo2/bar/baz.go",
		},
		SourceDir: runConfig.SourceDir,
		Patterns: &[]Pattern{
			{
				ID: "gorevive",
				Parameters: []PatternParameter{
					{
						Name:  "gorevive",
						Value: "vars",
					},
				},
			},
			{
				ID: "foo",
				Parameters: []PatternParameter{
					{
						Name:  "some name",
						Value: "some value",
					},
					{
						Name:    "bar",
						Default: "nofunc",
					},
				},
			},
		},
	}

	// Act
	toolExecution, err := newToolExecution(runConfig)

	// Assert
	if assert.NoError(t, err) {
		assert.Equal(t, expectedToolExecution, *toolExecution)
	}
}

func TestNewToolExecution_NoMatchingToolConfigured(t *testing.T) {
	// Arrange
	srcDir := "sourceDir"
	runConfig := RunConfiguration{
		ToolConfigurationDir: filepath.Join(testsResourcesLocation, "unmatched_tools"),
		SourceDir:            srcDir,
	}

	expectedToolExecution := ToolExecution{
		ToolDefinition: ToolDefinition{
			Name:    "trivy",
			Version: "1.0.0",
			Patterns: &[]Pattern{
				{
					ID:          "secret",
					Category:    "Security",
					Level:       "Critical",
					SubCategory: "Cryptography",
					ScanType:    "Secrets",
				},
			},
		},
		Files: &[]string{
			"foo/bar/baz.go",
			"foo2/bar/baz.go",
		},
		SourceDir: runConfig.SourceDir,
		Patterns:  nil,
	}

	// Act
	toolExecution, err := newToolExecution(runConfig)

	// Assert
	if assert.NoError(t, err) {
		assert.Equal(t, expectedToolExecution, *toolExecution)
	}
}

func TestNewToolExecution_NonExistingConfiguration(t *testing.T) {
	// Arrange
	srcDir := "sourceDir"
	runConfig := RunConfiguration{
		ToolConfigurationDir: filepath.Join(testsResourcesLocation, "non_existing_configuration"),
		SourceDir:            srcDir,
	}

	expectedToolExecution := ToolExecution{
		ToolDefinition: ToolDefinition{
			Name:    "trivy",
			Version: "1.0.0",
			Patterns: &[]Pattern{
				{
					ID:       "secret",
					Category: "Security",
					Level:    "Critical",
				},
			},
		},
		Files:     nil,
		SourceDir: runConfig.SourceDir,
		Patterns:  nil,
	}

	// Act
	toolExecution, err := newToolExecution(runConfig)

	// Assert
	if assert.NoError(t, err) {
		assert.Equal(t, expectedToolExecution, *toolExecution)
	}
}

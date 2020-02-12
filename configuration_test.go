package codacytool

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const configurationTestResources = "configuration"

func TestParseConfiguration(t *testing.T) {
	location := filepath.Join(testsResourcesLocation, configurationTestResources, ".codacyrc.valid")

	configuration, err := ParseConfiguration(location)

	assert.Nil(t, err, "Configuration file %s could not be found.", location)

	expected := 2
	got := len(configuration.Files)
	assert.Equal(t, expected, got)
}

func TestParseEmptyConfiguration(t *testing.T) {
	location := filepath.Join(testsResourcesLocation, configurationTestResources, ".codacyrc.empty")

	_, err := ParseConfiguration(location)

	assert.NotNil(t, err)
}

func TestParseInvalidConfiguration(t *testing.T) {
	location := filepath.Join(testsResourcesLocation, configurationTestResources, ".codacyrc.invalid")
	configuration, err := ParseConfiguration(location)

	assert.Nil(t, err)

	// Files is present
	expected := 2
	got := len(configuration.Files)
	assert.Equal(t, expected, got)

	// Tools has a typo, so it must not parse it and return 0 tools
	expected = 0
	got = len(configuration.Tools)
	assert.Equal(t, expected, got)
}

func TestParseNonExistingConfiguration(t *testing.T) {
	location := filepath.Join(testsResourcesLocation, configurationTestResources, ".codacyrc.notexists")

	_, err := ParseConfiguration(location)

	assert.NotNil(t, err)
}

func testingToolDefinition() Tool {
	definitionPatterns := []Pattern{
		{
			PatternID: "testing",
			Parameters: []PatternParameter{
				{
					Name:    "val",
					Default: "val",
				},
			},
		}, {
			PatternID: "testing_number",
			Parameters: []PatternParameter{
				{
					Name:    "val",
					Default: "6",
				},
			},
		}, {
			PatternID: "testing_default",
			Parameters: []PatternParameter{
				{
					Name:    "def",
					Default: "2",
				},
			},
		}, {
			PatternID: "no_param",
		}, {
			PatternID: "not_used",
		},
	}

	toolDefinition := ToolDefinition{
		Name:     "example",
		Patterns: definitionPatterns,
	}

	return Tool{Definition: toolDefinition}
}

func TestWithDefaultPatterns(t *testing.T) {
	tool := testingToolDefinition()

	tool.Patterns = []Pattern{
		{
			PatternID: "testing",
		},
		{
			PatternID: "testing_default",
			Parameters: []PatternParameter{
				{
					Name:  "def",
					Value: "3",
				},
			},
		}, {
			PatternID: "no_param",
		}, {
			PatternID: "testing_number",
		},
	}

	tool.config.Tools = []ToolDefinition{
		{
			Name:     "example",
			Patterns: tool.Patterns,
		},
	}

	patternsFromConfiguration := patternsFromConfig(tool.Definition.Name, tool.config)
	patterns := withDefaultParameters(tool.Definition, tool.config, patternsFromConfiguration)

	assert.Equal(t, len(tool.Patterns), len(patterns))

	expectedPatterns := []Pattern{
		{
			PatternID: "testing",
			Parameters: []PatternParameter{
				{
					Name:  "val",
					Value: "val",
				},
			},
		},
		{
			PatternID: "testing_default",
			Parameters: []PatternParameter{
				{
					Name:  "def",
					Value: "3",
				},
			},
		},
		{
			PatternID: "no_param",
		},
		{
			PatternID: "testing_number",
			Parameters: []PatternParameter{
				{
					Name:  "val",
					Value: 6.0,
				},
			},
		},
	}

	assert.Equal(t, expectedPatterns, patterns)
}

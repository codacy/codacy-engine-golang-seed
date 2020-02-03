package codacytool

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
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

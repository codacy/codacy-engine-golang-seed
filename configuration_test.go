package codacytool

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseConfiguration(t *testing.T) {
	location := testsResourcesLocation + "/.codacyrc"

	configuration, err := ParseConfiguration(location)

	assert.Nil(t, err, "Configuration file %s could not be found.", location)

	expected := 2
	got := len(configuration.Files)
	assert.Equal(t, expected, got)
}

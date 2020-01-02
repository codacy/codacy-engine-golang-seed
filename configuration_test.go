package codacytool

import (
	"testing"
)

func TestParseConfiguration(t *testing.T) {
	location := testsResourcesLocation + "/.codacyrc"

	configuration, err := ParseConfiguration(location)

	if err != nil {
		t.Errorf("Configuration file %s could not be found.", location)
	}
	expected := 2
	got := len(configuration.Files)
	if got != expected {
		t.Errorf("Expected number of files in configuration: %d. Got %d", expected, got)
	}
}

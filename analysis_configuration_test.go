package codacytool

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadAnalysisConfiguration(t *testing.T) {
	type testData struct {
		fileLocation          string
		expectedConfiguration AnalysisConfiguration
	}

	testSet := map[string]testData{
		"non-existent": {
			fileLocation:          "non-existent",
			expectedConfiguration: AnalysisConfiguration{},
		},
		"invalid": {
			fileLocation:          "configuration/invalid",
			expectedConfiguration: AnalysisConfiguration{},
		},
		"empty": {
			fileLocation: "configuration/empty",
			expectedConfiguration: AnalysisConfiguration{
				Files: &[]string{},
				Tools: &[]ToolDefinition{},
			},
		},
	}

	for testName, testData := range testSet {
		t.Run(testName, func(t *testing.T) {
			// Arrange
			runConfig := RunConfiguration{
				ToolConfigurationDir: filepath.Join(testsResourcesLocation, testData.fileLocation),
			}

			// Act
			config := loadAnalysisConfiguration(runConfig)

			// Assert
			assert.Equal(t, testData.expectedConfiguration, config)
		})
	}
}

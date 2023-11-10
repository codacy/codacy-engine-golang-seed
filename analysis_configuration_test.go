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
		expectedError         string
	}

	testSet := map[string]testData{
		"non-existent": {
			fileLocation:          "non-existent",
			expectedConfiguration: AnalysisConfiguration{},
			expectedError:         "",
		},
		"invalid": {
			fileLocation:          "configuration/invalid",
			expectedConfiguration: AnalysisConfiguration{},
			expectedError:         "invalid character 'i' looking for beginning of value",
		},
		"empty": {
			fileLocation: "configuration/empty",
			expectedConfiguration: AnalysisConfiguration{
				Files: &[]string{},
				Tools: &[]ToolDefinition{},
			},
			expectedError: "",
		},
	}

	for testName, testData := range testSet {
		t.Run(testName, func(t *testing.T) {
			// Arrange
			runConfig := RunConfiguration{
				ToolConfigurationDir: filepath.Join(testsResourcesLocation, testData.fileLocation),
			}

			// Act
			config, err := loadAnalysisConfiguration(runConfig)

			// Assert
			assert.Equal(t, testData.expectedConfiguration, config)
			if err != nil {
				assert.Equal(t, testData.expectedError, err.Error())
			}
		})
	}
}

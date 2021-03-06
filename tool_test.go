package codacytool

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"path/filepath"
	"strings"
	"testing"
)

type ToolTestSuite struct {
	suite.Suite
	runConfig RunConfiguration
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ToolTestSuite))
}

func (suite *ToolTestSuite) SetupTest() {
	suite.runConfig = RunConfiguration{
		toolConfigsBasePath: filepath.Join(testsResourcesLocation, "tool"),
		sourceDir:           "./",
	}
}

func testingTool(name, version string) (ToolDefinition, string) {
	patternObj, patternJSON := pattern()
	toolRepresentationAsJSON := fmt.Sprintf(`{"name":"%s","version":"%s","patterns":[%s]}`, name, version, patternJSON)
	if version == "" {
		toolRepresentationAsJSON = fmt.Sprintf(`{"name":"%s","patterns":[%s]}`, name, patternJSON)
	}
	return ToolDefinition{
		Name:     name,
		Version:  version,
		Patterns: []Pattern{patternObj},
	}, toolRepresentationAsJSON
}

func (suite *ToolTestSuite) TestToolToJSON() {
	name := "gorevive"
	version := "0.0.1"
	tool, toolJSONExpected := testingTool(name, version)
	toolAsJSON, err := tool.ToJSON()

	assert.Nil(suite.T(), err, "Failed converting to JSON")
	assert.Equal(suite.T(), toolJSONExpected, string(toolAsJSON), "Failed converting to JSON")
}

func (suite *ToolTestSuite) TestToolToJSONWithoutVersion() {
	name := "gorevive"
	tool, toolJSONExpected := testingTool(name, "")
	toolAsJSON, err := tool.ToJSON()

	assert.Nil(suite.T(), err, "Failed converting to JSON")
	assert.Equal(suite.T(), toolJSONExpected, string(toolAsJSON), "Failed converting to JSON")
}

func (suite *ToolTestSuite) TestLoadToolDefinition() {
	patternsFileLocation := defaultDefinitionFile(suite.runConfig.toolConfigsBasePath)
	tool, err := LoadToolDefinition(patternsFileLocation)

	assert.Nil(suite.T(), err, "Failed to load tool %s", patternsFileLocation)
	assert.Equal(suite.T(), "gorevive", tool.Name)

	numPatterns := len(tool.Patterns)
	expectedPatterns := 1

	assert.Equal(suite.T(), expectedPatterns, numPatterns)
}

func (suite *ToolTestSuite) TestPrintResults() {
	issue := testIssue()
	resultString := resultAsString([]Issue{
		issue,
		issue,
	})

	res := strings.TrimRight(string(resultString), "\n")

	expected, _ := issue.ToJSON()
	expectedAsString := string(expected) + "\n" + string(expected)
	assert.Equal(suite.T(), expectedAsString, res)
}

func (suite *ToolTestSuite) TestDefaultTool() {
	tool := defaultTool(suite.runConfig)

	patternsLen := len(tool.Patterns)
	expectedLen := 1
	assert.Equal(suite.T(), expectedLen, patternsLen)

	toolName := tool.Definition.Name
	assert.Equal(suite.T(), "gorevive", toolName)
}
func (suite *ToolTestSuite) TestPatternsFromConfig() {
	toolName := "gorevive"
	configFile := defaultConfigurationFile(suite.runConfig.toolConfigsBasePath)
	config, err := ParseConfiguration(configFile)
	assert.Nil(suite.T(), err, "Error parsing config file %s", configFile)

	patterns := patternsFromConfig(toolName, config)

	patternsLen := len(patterns)
	expectedLen := 1
	assert.Equal(suite.T(), expectedLen, patternsLen)

	expectedID := "gorevive"
	assert.Equal(suite.T(), expectedID, patterns[0].PatternID)
}

func (suite *ToolTestSuite) TestStartTool() {
	impl := ToolImplementationTest{}
	issue := testIssue()

	res, err := startToolImplementation(impl, suite.runConfig)
	resAsString := resultAsString(res)

	assert.Nil(suite.T(), err)

	expected, _ := issue.ToJSON()
	expectedAsString := string(expected)
	assert.Equal(suite.T(), expectedAsString, resAsString)
}

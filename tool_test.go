package codacytool

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"os"
	"strings"
	"testing"
)

type ToolTestSuite struct {
	suite.Suite
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ToolTestSuite))
}

func (suite *ToolTestSuite) SetupTest() {
	toolConfigsBasePathFlag = &testsResourcesLocation
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
	name := "govet"
	version := "0.0.1"
	tool, toolJSONExpected := testingTool(name, version)
	toolAsJSON, err := tool.ToJSON()

	assert.Nil(suite.T(), err, "Failed converting to JSON")
	assert.Equal(suite.T(), toolJSONExpected, string(toolAsJSON), "Failed converting to JSON")
}

func (suite *ToolTestSuite) TestToolToJSONWithoutVersion() {
	name := "govet"
	tool, toolJSONExpected := testingTool(name, "")
	toolAsJSON, err := tool.ToJSON()

	assert.Nil(suite.T(), err, "Failed converting to JSON")
	assert.Equal(suite.T(), toolJSONExpected, string(toolAsJSON), "Failed converting to JSON")
}

func (suite *ToolTestSuite) TestLoadToolDefinition() {
	patternsFileLocation := defaultDefinitionFile()
	tool, err := LoadToolDefinition(patternsFileLocation)

	assert.Nil(suite.T(), err, "Failed to load tool %s", patternsFileLocation)
	assert.Equal(suite.T(), "govet", tool.Name)

	numPatterns := len(tool.Patterns)
	expectedPatterns := 1

	assert.Equal(suite.T(), expectedPatterns, numPatterns)
}

func (suite *ToolTestSuite) TestPrintResults() {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	issue := testIssue()
	printResult([]Issue{
		issue,
		issue,
	})
	res := strings.TrimRight(buf.String(), "\n")
	expected, _ := issue.ToJSON()
	expectedAsString := string(expected) + "\n" + string(expected)
	assert.Equal(suite.T(), expectedAsString, res)
}

func (suite *ToolTestSuite) TestDefaultTool() {
	tool := defaultTool()

	patternsLen := len(tool.Patterns)
	expectedLen := 1
	assert.Equal(suite.T(), expectedLen, patternsLen)

	toolName := tool.Definition.Name
	assert.Equal(suite.T(), "govet", toolName)
}
func (suite *ToolTestSuite) TestPatternsFromConfig() {
	toolName := "govet"
	configFile := defaultConfigurationFile()
	config, err := ParseConfiguration(configFile)
	assert.Nil(suite.T(), err, "Error parsing config file %s", configFile)

	patterns := patternsFromConfig(toolName, config)

	patternsLen := len(patterns)
	expectedLen := 1
	assert.Equal(suite.T(), expectedLen, patternsLen)

	expectedID := "govet"
	assert.Equal(suite.T(), expectedID, patterns[0].PatternID)
}

type ToolImplementationTest struct{}

func (i ToolImplementationTest) Run(tool Tool, sourceDir string) ([]Issue, error) {
	issue := testIssue()
	return []Issue{issue}, nil
}

func (suite *ToolTestSuite) TestStartTool() {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	impl := ToolImplementationTest{}
	startToolImplementation(impl, "./")
	issue := testIssue()

	res := strings.TrimRight(buf.String(), "\n")

	expected, _ := issue.ToJSON()
	expectedAsString := string(expected)
	assert.Equal(suite.T(), res, expectedAsString)
}

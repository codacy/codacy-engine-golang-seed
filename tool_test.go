package codacytool

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func setup() {
	os.Setenv(_basePathEnvVar, testsResourcesLocation)
}
func shutdown() {
	os.Unsetenv(_basePathEnvVar)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
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

func TestToolToJSON(t *testing.T) {
	name := "govet"
	version := "0.0.1"
	tool, toolJSONExpected := testingTool(name, version)
	toolAsJSON, err := tool.ToJSON()

	if err != nil {
		t.Error("Failed converting to JSON")
	}

	if string(toolAsJSON) != toolJSONExpected {
		t.Errorf("Expected: %s ; Got: %s", toolJSONExpected, toolAsJSON)
	}
}

func TestToolToJSONWithoutVersion(t *testing.T) {
	name := "govet"
	tool, toolJSONExpected := testingTool(name, "")
	toolAsJSON, err := tool.ToJSON()

	if err != nil {
		t.Error("Failed converting to JSON")
	}

	if string(toolAsJSON) != toolJSONExpected {
		t.Errorf("Expected: %s ; Got: %s", toolJSONExpected, toolAsJSON)
	}
}

func TestLoadToolDefinition(t *testing.T) {
	patternsFileLocation := defaultDefinitionFile()
	tool, err := LoadToolDefinition(patternsFileLocation)

	if err != nil {
		t.Errorf("Failed to load tool %s", patternsFileLocation)
	}

	if tool.Name != "govet" {
		t.Errorf("Expected: %s; Got: %s", "govet", tool.Name)
	}

	numPatterns := len(tool.Patterns)
	expectedPatterns := 1
	if numPatterns != expectedPatterns {
		t.Errorf("Expected: %d; Got: %d", expectedPatterns, numPatterns)
	}
}

func TestPrintResults(t *testing.T) {
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
	if res != expectedAsString {
		t.Errorf("Expected: %s; Got: %s", expected, res)
	}
}

func TestDefaultTool(t *testing.T) {
	tool := defaultTool()

	patternsLen := len(tool.Patterns)
	expectedLen := 1
	if patternsLen != expectedLen {
		t.Errorf("Expected len: %d; Got: %d", expectedLen, patternsLen)
	}

	toolName := tool.Definition.Name
	if toolName != "govet" {
		t.Errorf("Expected len: %s; Got: %s", "govet", toolName)
	}
}
func TestPatternsFromConfig(t *testing.T) {
	toolName := "govet"
	configFile := defaultConfigurationFile()
	config, err := ParseConfiguration(configFile)
	if err != nil {
		t.Errorf("Error parsing config file %s", configFile)
	}
	patterns := patternsFromConfig(toolName, config)

	patternsLen := len(patterns)
	expectedLen := 1
	if patternsLen != expectedLen {
		t.Errorf("Expected len: %d; Got: %d", expectedLen, patternsLen)
	}
	expectedID := "govet"
	if patterns[0].PatternID != expectedID {
		t.Errorf("Expected PatternID: %s; Got: %s", expectedID, patterns[0].PatternID)
	}
}

type ToolImplementationTest struct{}

func (i ToolImplementationTest) Run(tool Tool, sourceDir string) ([]Issue, error) {
	issue := testIssue()
	return []Issue{issue}, nil
}

func TestStartTool(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	impl := ToolImplementationTest{}
	StartTool(impl, "./")
	issue := testIssue()

	res := strings.TrimRight(buf.String(), "\n")

	expected, _ := issue.ToJSON()
	expectedAsString := string(expected)
	if res != expectedAsString {
		t.Errorf("Expected: %s; Got: %s", expected, res)
	}
}

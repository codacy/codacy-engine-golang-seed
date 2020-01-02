package codacytool

import (
	"fmt"
	"testing"
)

func testingTool(name, version string) (Tool, string) {
	patternObj, patternJSON := pattern()
	toolRepresentationAsJSON := fmt.Sprintf(`{"name":"%s","version":"%s","patterns":[%s]}`, name, version, patternJSON)
	if version == "" {
		toolRepresentationAsJSON = fmt.Sprintf(`{"name":"%s","patterns":[%s]}`, name, patternJSON)
	}
	return Tool{
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

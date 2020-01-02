package codacytool

import (
	"fmt"
	"testing"
)

func patternDescription() (PatternDescription, string) {
	patternParam, patternParamJSON := patternParam()

	return PatternDescription{
		PatternID:   "foo",
		Title:       "bar",
		Description: "bar",
		Parameters:  []PatternParameter{patternParam},
		TimeToFix:   10,
	}, fmt.Sprintf(`{"patternId":"foo","title":"bar","description":"bar","parameters":[%s],"timeToFix":10}`, patternParamJSON)
}

func TestPatternDescriptionToJSON(t *testing.T) {
	patternDescr, patternDescrJSONExpected := patternDescription()
	patternDescrAsJSON, err := patternDescr.ToJSON()

	if err != nil {
		t.Error("Failed converting to JSON")
	}

	if string(patternDescrAsJSON) != patternDescrJSONExpected {
		t.Errorf("Expected: %s ; Got: %s", patternDescrJSONExpected, patternDescrAsJSON)
	}
}

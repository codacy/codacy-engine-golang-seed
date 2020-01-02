package codacytool

import (
	"fmt"
	"testing"
)

func patternParam() (PatternParameter, string) {
	return PatternParameter{
		Name:  "param",
		Value: "bar",
	}, `{"name":"param","value":"bar"}`
}

func pattern() (Pattern, string) {
	patternParam, patternParamJSON := patternParam()

	return Pattern{
		PatternID:  "foo",
		Category:   "bar",
		Parameters: []PatternParameter{patternParam},
	}, fmt.Sprintf(`{"patternId":"foo","category":"bar","parameters":[%s]}`, patternParamJSON)
}

func patternWithLevel() (Pattern, string) {
	patternParam, patternParamJSON := patternParam()

	return Pattern{
		PatternID:  "foo",
		Category:   "bar",
		Level:      "warn",
		Parameters: []PatternParameter{patternParam},
	}, fmt.Sprintf(`{"patternId":"foo","category":"bar","level":"warn","parameters":[%s]}`, patternParamJSON)
}

func patternWithoutParams() (Pattern, string) {
	return Pattern{
		PatternID:  "foo",
		Category:   "bar",
		Level:      "warn",
		Parameters: []PatternParameter{},
	}, fmt.Sprintf(`{"patternId":"foo","category":"bar","level":"warn"}`)
}

func TestPatternToJSON(t *testing.T) {
	pattern, patternRepresentation := patternWithoutParams()
	patternAsJSON, err := pattern.ToJSON()

	if err != nil {
		t.Error("Failed converting to JSON")
	}

	if string(patternAsJSON) != patternRepresentation {
		t.Errorf("Expected: %s ; Got: %s", patternRepresentation, patternAsJSON)
	}

}

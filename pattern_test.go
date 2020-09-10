package codacytool

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
	}, fmt.Sprintf(`{"patternId":"foo","category":"bar","parameters":[%s],"enabled":false}`, patternParamJSON)
}

func patternWithLevel() (Pattern, string) {
	patternParam, patternParamJSON := patternParam()

	return Pattern{
		PatternID:  "foo",
		Category:   "bar",
		Level:      "warn",
		Parameters: []PatternParameter{patternParam},
	}, fmt.Sprintf(`{"patternId":"foo","category":"bar","level":"warn","parameters":[%s],"enabled":false}`, patternParamJSON)
}

func patternWithoutParams() (Pattern, string) {
	return Pattern{
		PatternID:  "foo",
		Category:   "bar",
		Level:      "warn",
		Parameters: []PatternParameter{},
	}, fmt.Sprintf(`{"patternId":"foo","category":"bar","level":"warn","enabled":false}`)
}

func TestPatternToJSON(t *testing.T) {
	pattern, patternRepresentation := patternWithoutParams()
	patternAsJSON, err := pattern.ToJSON()

	assert.Nil(t, err, "Failed converting to JSON")

	assert.Equal(t, patternRepresentation, string(patternAsJSON))
}

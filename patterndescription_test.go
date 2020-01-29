package codacytool

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

	assert.Nil(t, err, "Failed converting to JSON")

	assert.Equal(t, patternDescrJSONExpected, string(patternDescrAsJSON))
}

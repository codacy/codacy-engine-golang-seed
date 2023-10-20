package codacytool

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func TestResultsToJSON(t *testing.T) {
	// Arrange
	issue := Issue{
		File:      "file",
		Line:      5,
		Message:   "message",
		PatternID: "pattern ID",
	}
	fileError := FileError{
		File:    "file-error",
		Message: "file-error",
	}
	badResult := BadResult{}

	expectedJSONResults := []string{
		`{"filename":"file","line":5,"message":"message","patternId":"pattern ID"}`,
		`{"filename":"file-error","message":"file-error"}`,
	}

	// Act
	jsonResults := Results{issue, fileError, badResult}.ToJSON()

	// Assert
	// Since a JSON object does not have order, we can't simply assert by doing `assert.ElementsMatch`.
	// To guarantee that the results are in the same order for proper comparison we can simply sort them by length.
	slices.SortFunc(jsonResults, func(a, b string) bool { return len(a) > len(b) })
	slices.SortFunc(expectedJSONResults, func(a, b string) bool { return len(a) > len(b) })

	for i, jsonResult := range jsonResults {
		assert.JSONEq(t, expectedJSONResults[i], jsonResult)
	}
}

type BadResult struct{}

func (r BadResult) ToJSON() ([]byte, error) {
	return nil, assert.AnError
}

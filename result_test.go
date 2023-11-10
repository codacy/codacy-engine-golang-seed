package codacytool

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
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
	sort.Slice(jsonResults, func(a, b int) bool { return len(jsonResults[a]) > len(jsonResults[b]) })
	sort.Slice(expectedJSONResults, func(a, b int) bool { return len(expectedJSONResults[a]) > len(expectedJSONResults[b]) })

	for i, jsonResult := range jsonResults {
		assert.JSONEq(t, expectedJSONResults[i], jsonResult)
	}
}

func TestResultsGetFile(t *testing.T) {
	// Arrange
	issue := Issue{File: "issue-file"}
	fileError := FileError{File: "file-error"}

	// Act
	issueFile := issue.GetFile()
	fileErrorFile := fileError.GetFile()

	// Assert
	assert.Equal(t, "issue-file", issueFile)
	assert.Equal(t, "file-error", fileErrorFile)
}

type BadResult struct{}

func (r BadResult) ToJSON() ([]byte, error) {
	return nil, assert.AnError
}
func (r BadResult) GetFile() string {
	return "not used"
}

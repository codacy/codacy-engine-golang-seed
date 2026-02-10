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
		SourceID:  "CVE-2025-11111",
	}
	fileError := FileError{
		File:    "file-error",
		Message: "file-error",
	}
	sbom := SBOM{
		BomFormat:   CycloneDXJSON,
		SpecVersion: "1.6",
		Sbom:        `{"bomFormat":"CycloneDX","specVersion":"1.6","metadata"...}`,
	}
	badResult := BadResult{}

	expectedJSONResults := []string{
		`{"filename":"file","line":5,"message":"message","patternId":"pattern ID", "sourceId":"CVE-2025-11111"}`,
		`{"filename":"file-error","message":"file-error"}`,
		`{"bomFormat":"CycloneDXJSON","specVersion":"1.6","sbom":"{\"bomFormat\":\"CycloneDX\",\"specVersion\":\"1.6\",\"metadata\"...}"}`,
	}

	// Act
	jsonResults := Results{issue, fileError, sbom, badResult}.ToJSON()

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
	sbom := SBOM{}

	// Act
	issueFile := issue.GetFile()
	fileErrorFile := fileError.GetFile()
	sbomFile := sbom.GetFile()

	// Assert
	assert.Equal(t, "issue-file", issueFile)
	assert.Equal(t, "file-error", fileErrorFile)
	assert.Empty(t, sbomFile)
}

type BadResult struct{}

func (r BadResult) ToJSON() ([]byte, error) {
	return nil, assert.AnError
}
func (r BadResult) GetFile() string {
	return "not used"
}

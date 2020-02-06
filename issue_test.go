package codacytool

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func testIssue() Issue {
	file := "src/foo.go"
	line := 5
	message := "Foo bar"
	patternID := "foo"
	return Issue{
		File:      file,
		Line:      line,
		Message:   message,
		PatternID: patternID,
	}
}
func TestIssueToJSON(t *testing.T) {
	issue := testIssue()

	issueAsJSON, err := issue.ToJSON()
	if err != nil {
		t.Error("Failed converting to JSON")
	}
	expectedJSON := fmt.Sprintf(`{"patternId":"%s","filename":"%s","line":%d,"message":"%s"}`, issue.PatternID, issue.File, issue.Line, issue.Message)

	assert.Equal(t, expectedJSON, string(issueAsJSON))
}

package codacytool

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

func testIssueWithSuggestion() Issue {
	file := "src/foo.go"
	line := 5
	message := "Foo bar"
	patternID := "foo"
	suggestion := "foo := 42"
	return Issue{
		File:       file,
		Line:       line,
		Message:    message,
		PatternID:  patternID,
		Suggestion: suggestion,
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

func TestIssueWithSuggestionToJSON(t *testing.T) {
	issue := testIssueWithSuggestion()

	issueAsJSON, _ := issue.ToJSON()
	expectedJSON := fmt.Sprintf(`{"patternId":"%s","filename":"%s","line":%d,"message":"%s","suggestion":"%s"}`, issue.PatternID, issue.File, issue.Line, issue.Message, issue.Suggestion)

	assert.Equal(t, expectedJSON, string(issueAsJSON))
}

package codacytool

import (
	"fmt"
	"testing"
)

func TestIssueToJSON(t *testing.T) {
	file := "src/foo.go"
	line := "5"
	message := "Foo bar"
	patternID := "foo"
	issue := Issue{
		File:      file,
		Line:      line,
		Message:   message,
		PatternID: patternID,
	}

	issueAsJSON, err := issue.ToJSON()
	if err != nil {
		t.Error("Failed converting to JSON")
	}
	expectedJSON := fmt.Sprintf(`{"patternId":"%s","file":"%s","line":"%s","message":"%s"}`, patternID, file, line, message)
	if string(issueAsJSON) != expectedJSON {
		t.Errorf("Expected: %s; Got: %s", expectedJSON, issueAsJSON)
	}
}

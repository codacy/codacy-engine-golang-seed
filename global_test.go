package codacytool

import (
	"time"
)

const (
	testsResourcesLocation = "./tests/"
)

// ToolImplementationTest mock tool implementation
type ToolImplementationTest struct{}

// Run mock tool for testing run implementation
func (i ToolImplementationTest) Run(tool Tool, sourceDir string) ([]Issue, error) {
	time.Sleep(5 * time.Millisecond)
	issue := testIssue()
	return []Issue{issue}, nil
}

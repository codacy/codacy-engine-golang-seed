package codacytool

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
	"time"
)

func TestRunWithTimeout(t *testing.T) {
	impl := ToolImplementationTest{}
	runConfiguration := RunConfiguration{
		toolConfigsBasePath: filepath.Join(testsResourcesLocation, "tool"),
		sourceDir:           "./",
		timeoutDuration:     2 * time.Second,
	}
	expectedIssues := []Issue{
		testIssue(),
	}

	result := runWithTimeout(impl, runConfiguration)

	assert.Equal(t, expectedIssues, result)
}

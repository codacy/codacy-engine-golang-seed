package codacytool

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const testsResourcesLocation = "./tests/"

func TestRunWithTimeout_NonExistingToolDefinition(t *testing.T) {
	// Arrange
	tool := SingleIssueTool{}
	runConfiguration := RunConfiguration{
		ToolConfigurationDir: "non-existing",
		SourceDir:            "./",
		Debug:                true,
		Timeout:              2 * time.Second,
	}

	// Act
	result, code := runWithTimeout(tool, runConfiguration)

	// Assert
	assert.Equal(t, 1, code)
	assert.Nil(t, result)
}

func TestRunWithTimeout_NonJSONToolDefinition(t *testing.T) {
	// Arrange
	tool := SingleIssueTool{}
	runConfiguration := RunConfiguration{
		ToolConfigurationDir: filepath.Join(testsResourcesLocation, "invalid_tool"),
		SourceDir:            "./",
		Debug:                true,
		Timeout:              2 * time.Second,
	}

	// Act
	result, code := runWithTimeout(tool, runConfiguration)

	// Assert
	assert.Equal(t, 1, code)
	assert.Nil(t, result)
}

func TestRunWithTimeout_ToolError(t *testing.T) {
	// Arrange
	tool := ErrorTool{e: assert.AnError}
	runConfiguration := RunConfiguration{
		ToolConfigurationDir: filepath.Join(testsResourcesLocation, "tool"),
		SourceDir:            "./",
		Debug:                true,
		Timeout:              2 * time.Second,
	}

	// Act
	result, code := runWithTimeout(tool, runConfiguration)

	// Assert
	assert.Equal(t, 1, code)
	assert.Nil(t, result)
}

func TestRunWithTimeout_Timeout(t *testing.T) {
	// Arrange
	tool := LongRunningTool{duration: 2 * time.Second}
	runConfiguration := RunConfiguration{
		ToolConfigurationDir: filepath.Join(testsResourcesLocation, "tool"),
		SourceDir:            "./",
		Debug:                true,
		Timeout:              tool.duration - 1*time.Second,
	}

	// Act
	result, code := runWithTimeout(tool, runConfiguration)

	// Assert
	assert.Equal(t, 2, code)
	assert.Nil(t, result)
}

func TestRunWithTimeout(t *testing.T) {
	issue := Issue{
		File:      "file",
		Line:      5,
		Message:   "message",
		PatternID: "pattern ID",
	}
	tool := SingleIssueTool{issue: issue}
	runConfiguration := RunConfiguration{
		ToolConfigurationDir: filepath.Join(testsResourcesLocation, "tool"),
		SourceDir:            "./",
		Debug:                true,
		Timeout:              1 * time.Second,
	}

	// Act
	result, code := runWithTimeout(tool, runConfiguration)

	// Assert
	assert.Equal(t, 0, code)
	assert.ElementsMatch(t, []Issue{issue}, result)
}

func TestGetTimeout(t *testing.T) {
	type testData struct {
		setEnvironment  func()
		expectedTimeout time.Duration
	}

	testSet := map[string]testData{
		"no environment variable": {
			setEnvironment:  func() {},
			expectedTimeout: defaultTimeout,
		},
		"invalid environment variable": {
			setEnvironment: func() {
				os.Setenv("TIMEOUT_SECONDS", "abc")
			},
			expectedTimeout: defaultTimeout,
		},
		"environment variable": {
			setEnvironment: func() {
				os.Setenv("TIMEOUT_SECONDS", "10")
			},
			expectedTimeout: 10 * time.Second,
		},
	}

	//
	for testName, testData := range testSet {
		t.Run(testName, func(t *testing.T) {
			// Arrange
			testData.setEnvironment()
			t.Cleanup(func() {
				os.Unsetenv("TIMEOUT_SECONDS")
			})

			// Act
			timeout := getTimeout()

			// Assert
			assert.Equal(t, testData.expectedTimeout, timeout)
		})
	}
}

func TestGetDebug(t *testing.T) {
	type testData struct {
		setEnvironment func()
		expectedDebug  bool
	}

	testSet := map[string]testData{
		"no environment variable": {
			setEnvironment: func() {},
			expectedDebug:  defaultDebug,
		},
		"invalid environment variable": {
			setEnvironment: func() {
				os.Setenv("DEBUG", "abc")
			},
			expectedDebug: defaultDebug,
		},
		"environment variable": {
			setEnvironment: func() {
				os.Setenv("DEBUG", "T")
			},
			expectedDebug: true,
		},
	}

	//
	for testName, testData := range testSet {
		t.Run(testName, func(t *testing.T) {
			// Arrange
			testData.setEnvironment()
			t.Cleanup(func() {
				os.Unsetenv("DEBUG")
			})

			// Act
			debug := getDebug()

			// Assert
			assert.Equal(t, testData.expectedDebug, debug)
		})
	}
}

type SingleIssueTool struct {
	issue Issue
}

func (t SingleIssueTool) Run(_ context.Context, _ ToolExecution) ([]Issue, error) {
	return []Issue{t.issue}, nil
}

type LongRunningTool struct {
	duration time.Duration
}

func (t LongRunningTool) Run(_ context.Context, _ ToolExecution) ([]Issue, error) {
	time.Sleep(t.duration)
	return []Issue{}, nil
}

type ErrorTool struct {
	e error
}

func (t ErrorTool) Run(_ context.Context, _ ToolExecution) ([]Issue, error) {
	return nil, t.e
}

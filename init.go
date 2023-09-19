package codacytool

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// RunConfiguration contains the process run configuration
type RunConfiguration struct {
	sourceDir           string
	toolConfigsBasePath string
	timeoutDuration     time.Duration
}

const (
	defaultTimeout = 15 * time.Minute
)

// StartTool receives the tool implementation as parameter and run the tool
func StartTool(impl ToolImplementation) {
	cmdLineConfig := parseFlags()

	toolRunRes := runWithTimeout(impl, cmdLineConfig)

	printResult(toolRunRes)
}

func runTool(impl ToolImplementation, runConfiguration RunConfiguration) []Issue {
	res, err := startToolImplementation(impl, runConfiguration)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	return res
}

func runWithTimeout(impl ToolImplementation, runConfiguration RunConfiguration) []Issue {
	runToolCall := func() []Issue {
		return runTool(impl, runConfiguration)
	}

	result, err := runToolCallWithTimeout(runToolCall, runConfiguration.timeoutDuration)
	if err != nil {
		logrus.Fatal(fmt.Sprintf("Timeout of %s seconds exceeded", runConfiguration.timeoutDuration))
	}

	return result
}

func getTimeoutDuration() time.Duration {
	timeoutSecondsEnvVar, exists := os.LookupEnv("TIMEOUT_SECONDS")
	if exists {
		return parseTimeoutSeconds(timeoutSecondsEnvVar)
	}
	return defaultTimeout
}

func parseFlags() RunConfiguration {
	sourceDirFlag := flag.String("sourceDir", "/src", "source to analyse folder")
	toolConfigLocationFlag := flag.String("toolConfigLocation", "/", "Location of tool configuration")

	flag.Parse()

	return RunConfiguration{
		sourceDir:           *sourceDirFlag,
		toolConfigsBasePath: *toolConfigLocationFlag,
		timeoutDuration:     getTimeoutDuration(),
	}
}

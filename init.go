package codacytool

import (
	"flag"
	"fmt"
	"os"
	"time"
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

	runWithTimeout(impl, cmdLineConfig)
}

func runWithTimeout(impl ToolImplementation, runConfiguration RunConfiguration) {
	runMethodWithTimeout(func() {
		startToolImplementation(impl, runConfiguration)
	}, func() {
		fmt.Fprintf(os.Stderr, "Timeout of %s seconds exceeded", runConfiguration.timeoutDuration)
		os.Exit(1)
	}, runConfiguration.timeoutDuration)
}

func parseFlags() RunConfiguration {
	cmdLineConfig := RunConfiguration{
		sourceDir:           *flag.String("sourceDir", "/src", "source to analyse folder"),
		toolConfigsBasePath: *flag.String("toolConfigLocation", "/", "Location of tool configuration"),
		timeoutDuration:     timeoutSeconds(),
	}

	flag.Parse()
	return cmdLineConfig
}

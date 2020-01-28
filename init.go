package codacytool

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	sourceDirFlag           *string
	toolConfigsBasePathFlag *string
)

const (
	defaultTimeout = 15 * time.Minute
)

// StartTool receives the tool implementation as parameter and run the tool
func StartTool(impl ToolImplementation) {
	parseFlags()

	os.Setenv("TOOL_CONFIGS_BASEPATH", *toolConfigsBasePathFlag)

	runWithTimeout(impl, timeoutSeconds())
}

func runWithTimeout(impl ToolImplementation, maxDuration time.Duration) {
	runMethodWithTimeout(func() {
		startToolImplementation(impl, *sourceDirFlag)
	}, func() {
		fmt.Fprintf(os.Stderr, "Timeout of %s seconds exceeded", maxDuration)
		os.Exit(1)
	}, maxDuration)
}

func parseFlags() {
	sourceDirFlag = flag.String("sourceDir", "/src", "source to analyse folder")
	toolConfigsBasePathFlag = flag.String("toolConfigLocation", "/", "Location of tool configuration")

	flag.Parse()
}

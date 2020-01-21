package codacytool

import (
	"flag"
	"os"
)

var (
	sourceDirFlag           *string
	toolConfigsBasePathFlag *string
)

// StartTool receives the tool implementation as parameter and run the tool
func StartTool(impl ToolImplementation) {
	parseFlags()

	os.Setenv("TOOL_CONFIGS_BASEPATH", *toolConfigsBasePathFlag)

	startToolImplementation(impl, *sourceDirFlag)
}

func parseFlags() {
	sourceDirFlag = flag.String("sourceDir", "/src", "source to analyse folder")
	toolConfigsBasePathFlag = flag.String("toolConfigLocation", "/", "Location of tool configuration")

	flag.Parse()
}

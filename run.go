package codacytool

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// RunConfiguration contains the process run configuration
type RunConfiguration struct {
	SourceDir            string
	ToolConfigurationDir string
	Timeout              time.Duration
	Debug                bool
}

const (
	defaultTimeout              = 15 * time.Minute
	defaultDebug                = false
	defaultSourceDir            = "/src"
	defaultToolConfigurationDir = "/"
)

// StartTool receives the tool implementation as parameter and runs the tool.
// Issues found will be printed to the standard output in a JSON format.
//
// Return codes are as follows:
//   - 0 - Tool executed successfully
//   - 1 - An unknown error occurred while running the tool
//   - 2 - Execution timeout
func StartTool(tool Tool) int {
	runConfiguration := parseRunConfiguration()

	if runConfiguration.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	issues, retCode := runWithTimeout(tool, runConfiguration)
	if retCode != 0 {
		return retCode
	}

	printResult(issues)
	return 0
}

func runWithTimeout(tool Tool, runConfiguration RunConfiguration) ([]Issue, int) {
	type IssuesAndRetCode struct {
		issues  []Issue
		retCode int
	}

	c := make(chan IssuesAndRetCode, 1)
	ctx, cancel := context.WithTimeout(context.Background(), runConfiguration.Timeout)
	defer cancel()

	go func() {
		toolExec, err := newToolExecution(runConfiguration)
		if err != nil {
			logrus.Errorf("Failed to create tool execution: %s", err.Error())

			c <- IssuesAndRetCode{issues: nil, retCode: 1}
			return
		}

		issues, err := tool.Run(ctx, *toolExec)
		if err != nil {
			logrus.Errorf("Failed to run the tool: %s", err.Error())

			c <- IssuesAndRetCode{issues: nil, retCode: 1}
			return
		}

		c <- IssuesAndRetCode{issues: issues, retCode: 0}
	}()

	select {
	case res := <-c:
		return res.issues, res.retCode
	case <-ctx.Done():
		logrus.Errorf("Failed to run the tool: Context deadline (%s) exceeded", runConfiguration.Timeout)
		return nil, 2
	}
}

func parseRunConfiguration() RunConfiguration {
	sourceDir := flag.String("sourceDir", defaultSourceDir, "Directory with the source files to analyse")
	toolConfigurationDir := flag.String("toolConfigLocation", defaultToolConfigurationDir, "Directory of the tool's configuration")

	flag.Parse()

	return RunConfiguration{
		SourceDir:            *sourceDir,
		ToolConfigurationDir: *toolConfigurationDir,
		Timeout:              getTimeout(),
		Debug:                getDebug(),
	}
}

func getTimeout() time.Duration {
	timeoutVar, exists := os.LookupEnv("TIMEOUT_SECONDS")
	if exists {
		seconds, err := strconv.Atoi(timeoutVar)
		if err == nil {
			return time.Duration(seconds) * time.Second
		}
	}
	return defaultTimeout
}

func getDebug() bool {
	debugVar, exists := os.LookupEnv("DEBUG")
	if exists {
		debug, err := strconv.ParseBool(debugVar)
		if err == nil {
			return debug
		}
	}
	return defaultDebug
}

func printResult(issues []Issue) {
	var resultList []string

	for _, i := range issues {
		iJSON, err := i.ToJSON()
		if err != nil {
			fileError := FileError{
				Filename: i.File,
				Message:  err.Error(),
			}

			fileErrorJSON, _ := fileError.ToJSON()
			resultList = append(resultList, string(fileErrorJSON))
		} else {
			resultList = append(resultList, string(iJSON))
		}
	}

	fmt.Print(strings.Join(resultList, "\n"))
}

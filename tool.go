package codacytool

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

// Tool is the interface each tool must implement.
type Tool interface {
	Run(ctx context.Context, toolExecution ToolExecution) ([]Issue, error)
}

// ToolExecution has the data for the execution of a tool.
type ToolExecution struct {
	// ToolDefinition is the metadata for the tool that will execute the analysis.
	ToolDefinition ToolDefinition
	// SourceDir is the directory of `Files`.
	SourceDir string
	// Files is an array of file paths, relative to `SourceDir`, to the files that will be analysed.
	// If undefined, the analysis should include all files inside `SourceDir`.
	// If empty, no analysis should be made.
	Files *[]string
	// Patterns is an array of the patterns (rules) to be used for analysis.
	// It can be a subset of the tool's supported patterns.
	// If undefined, the analysis should use the tool's configuration file, if available.
	// Otherwise, use the tool's default patterns.
	Patterns *[]Pattern
}

func newToolExecution(runConfig RunConfiguration) (*ToolExecution, error) {
	toolDefinition, err := loadToolDefinition(runConfig)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("Tool definition: %+v", toolDefinition)

	analysisConfig := loadAnalysisConfiguration(runConfig)
	logrus.Debugf("Analysis configuration: %+v", analysisConfig)

	var patterns *[]Pattern
	configuredTool, exists := lo.Find(*analysisConfig.Tools, func(item ToolDefinition) bool {
		return toolDefinition.Name == item.Name
	})
	// If there is no configured tool that matches the container's tool definition, patterns will be nil. The underlying tool will know what to do in that case.
	// Otherwise we guarantee that the configured patterns have all the required parameters as specified in the container's tool definition.
	if exists {
		p := patternsWithDefaultParameters(*toolDefinition, configuredTool)
		patterns = &p
	}

	toolExecution := ToolExecution{
		ToolDefinition: *toolDefinition,
		SourceDir:      runConfig.SourceDir,
		Files:          analysisConfig.Files,
		Patterns:       patterns,
	}
	logrus.Debugf("Tool execution: %+v", toolExecution)

	return &toolExecution, nil
}

// patternsWithDefaultParameters returns the patterns of `configuredTool` with any missing parameters added, as specified by `toolDefinition`.
func patternsWithDefaultParameters(toolDefinition, configuredTool ToolDefinition) []Pattern {

	var patterns []Pattern
	for _, configuredToolPattern := range configuredTool.Patterns {

		// Configured pattern exists in the tool definition patterns
		toolDefinitionPattern, exists := lo.Find(toolDefinition.Patterns, func(item Pattern) bool {
			return configuredToolPattern.ID == item.ID
		})

		// Patterns have different number of parameters
		if exists {
			for _, toolDefintionPatternParam := range toolDefinitionPattern.Parameters {

				// Add non-existing params to configured tool pattern params, as specified in the tool definition.
				paramExists := lo.SomeBy(configuredToolPattern.Parameters, func(item PatternParameter) bool {
					return toolDefintionPatternParam.Name == item.Name
				})
				if !paramExists {
					logrus.Debugf("Pattern parameter added to configured tool: %+v", toolDefintionPatternParam)
					configuredToolPattern.Parameters = append(
						configuredToolPattern.Parameters,
						toolDefintionPatternParam,
					)
				}
			}
		}
		patterns = append(patterns, configuredToolPattern)
	}

	return patterns
}

const defaultToolDefinitionFile = "docs/patterns.json"

// ToolDefinition is the metadata of the tool to run.
type ToolDefinition struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
	// Patterns contains all of the tool's supported patterns.
	Patterns []Pattern `json:"patterns"`
}

// loadToolDefinition loads tool information from the tool definition file.
func loadToolDefinition(runConfig RunConfiguration) (*ToolDefinition, error) {
	fileLocation := filepath.Join(runConfig.ToolConfigurationDir, defaultToolDefinitionFile)

	fileContent, err := os.ReadFile(fileLocation)
	if err != nil {
		return nil, fmt.Errorf("failed to read tool definition file: %s\n%w", fileLocation, err)
	}

	toolDefinition := ToolDefinition{}
	if err := json.Unmarshal(fileContent, &toolDefinition); err != nil {
		return nil, fmt.Errorf("failed to parse tool definition file: %s\n%w", string(fileContent), err)
	}
	return &toolDefinition, nil
}

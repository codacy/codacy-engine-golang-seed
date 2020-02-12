package codacytool

import (
	"encoding/json"
	"errors"
	"strconv"
)

// Configuration configuration provided to the tool to specify files and patterns to analyse
type Configuration struct {
	Files []string         `json:"files"`
	Tools []ToolDefinition `json:"tools"`
}

// ParseConfiguration parses configuration file
func ParseConfiguration(fileLocation string) (Configuration, error) {
	config := Configuration{}
	jsonFileContent, err := readFile(fileLocation)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(jsonFileContent, &config)
	return config, err
}

func toolConfiguration(toolName string, configuration Configuration) (ToolDefinition, error) {
	for _, t := range configuration.Tools {
		if t.Name == toolName {
			return t, nil
		}
	}
	return ToolDefinition{}, errors.New("Tool configuration not found")
}

func patternDefinition(patternID string, toolDefinition ToolDefinition) (Pattern, error) {
	for _, pattern := range toolDefinition.Patterns {
		if pattern.PatternID == patternID {
			return pattern, nil
		}
	}
	return Pattern{}, errors.New("Pattern definition not found")
}

func containsParam(param PatternParameter, parameters []PatternParameter) bool {
	for _, p := range parameters {
		if p.Name == param.Name {
			return true
		}
	}
	return false
}

func valueOfParam(param interface{}) interface{} {
	v, err := strconv.ParseFloat(param.(string), 64)
	if err == nil {
		return v
	}
	return param
}

func withDefaultParameters(toolDefinition ToolDefinition, configuration Configuration, patternsFromConfig []Pattern) []Pattern {
	toolConfig, err := toolConfiguration(toolDefinition.Name, configuration)
	if err != nil {
		return patternsFromConfig
	}

	var patterns []Pattern
	for _, p := range toolConfig.Patterns {
		patternDefinition, err := patternDefinition(p.PatternID, toolDefinition)
		if err == nil && len(patternDefinition.Parameters) != len(p.Parameters) {
			for _, param := range patternDefinition.Parameters {
				if !containsParam(param, p.Parameters) {
					p.Parameters = append(p.Parameters, PatternParameter{
						Name:  param.Name,
						Value: valueOfParam(param.Default),
					})
				}
			}
		}

		patterns = append(patterns, p)
	}

	return patterns
}

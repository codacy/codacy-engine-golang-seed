package codacytool

// Configuration configuration provided to the tool to specify files and patterns to analyse
type Configuration struct {
	Files []string         `json:"files"`
	Tools []ToolDefinition `json:"tools"`
}

// ParseConfiguration parses configuration file
func ParseConfiguration(fileLocation string) (Configuration, error) {
	config := Configuration{}
	err := parseJSONFile(fileLocation, &config)
	return config, err
}

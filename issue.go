package codacytool

import (
	"encoding/json"
)

// Issue is the output for each issue found by the tool
type Issue struct {
	PatternID  string `json:"patternId"`
	File       string `json:"filename"`
	Line       int    `json:"line"`
	Message    string `json:"message"`
	Suggestion string `json:"suggestion,omitempty"`
}

// ToJSON returns the json representation of the issue
func (i *Issue) ToJSON() ([]byte, error) {
	return json.Marshal(i)
}

// FileError error analysing a file
type FileError struct {
	Filename string `json:"filename"`
	Message  string `json:"message"`
}

// ToJSON returns the json representation of the issue
func (i *FileError) ToJSON() ([]byte, error) {
	return json.Marshal(i)
}

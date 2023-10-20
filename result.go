package codacytool

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

// Result encompasses all possible results: Issues and File Errors.
type Result interface {
	// ToJSON returns a JSON representation of the result.
	ToJSON() ([]byte, error)
	// GetFile returns the file for this result.
	GetFile() string
}

// Issue is the output for each issue found by the tool.
type Issue struct {
	PatternID  string `json:"patternId"`
	File       string `json:"filename"`
	Line       int    `json:"line"`
	Message    string `json:"message"`
	Suggestion string `json:"suggestion,omitempty"`
}

func (i Issue) ToJSON() ([]byte, error) {
	return json.Marshal(i)
}
func (i Issue) GetFile() string {
	return i.File
}

// FileError represents an error analysing a file.
// If this result is returned from an analysis, the referenced file is not considered to have been analysed.
type FileError struct {
	File    string `json:"filename"`
	Message string `json:"message"`
}

func (i FileError) ToJSON() ([]byte, error) {
	return json.Marshal(i)
}
func (i FileError) GetFile() string {
	return i.File
}

type Results []Result

func (r Results) ToJSON() []string {
	var jsonResults []string

	for _, result := range r {
		jsonResult, err := result.ToJSON()
		if err != nil {
			logrus.Errorf("Failed to convert Result to JSON: %+v\n%s", result, err.Error())
		} else {
			jsonResults = append(jsonResults, string(jsonResult))
		}
	}

	return jsonResults
}

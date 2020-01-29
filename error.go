package codacytool

import (
	"encoding/json"
)

// FileError error analysing a file
type FileError struct {
	Filename string `json:"filename"`
	Message  string `json:"message"`
}

// ToJSON returns the json representation of the issue
func (i *FileError) ToJSON() ([]byte, error) {
	return json.Marshal(i)
}

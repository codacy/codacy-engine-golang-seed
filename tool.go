package codacytool

import (
	"encoding/json"
)

// Tool is the configuration of the tool to run
type Tool struct {
	Name     string    `json:"name"`
	Version  string    `json:"version,omitempty"`
	Patterns []Pattern `json:"patterns"`
}

// ToJSON returns the json representation of the tool
func (i *Tool) ToJSON() ([]byte, error) {
	return json.Marshal(i)
}

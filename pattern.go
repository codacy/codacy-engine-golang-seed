package codacytool

import (
	"encoding/json"
)

// Pattern pattern/rule inspected by tool
type Pattern struct {
	PatternID   string             `json:"patternId"`
	Category    string             `json:"category,omitempty"`
	Level       string             `json:"level,omitempty"`
	Parameters  []PatternParameter `json:"parameters,omitempty"`
	SubCategory string             `json:"subcategory,omitempty"`
}

// ToJSON returns the json representation of the pattern
func (i *Pattern) ToJSON() ([]byte, error) {
	return json.Marshal(i)
}

// PatternParameter parameter to customize a pattern
type PatternParameter struct {
	Name        string      `json:"name"`
	Value       interface{} `json:"value,omitempty"`
	Default     interface{} `json:"default,omitempty"`
	Description string      `json:"description,omitempty"`
}

// ToJSON returns the json representation of the PatternParameter
func (i *PatternParameter) ToJSON() ([]byte, error) {
	return json.Marshal(i)
}

package codacytool

import (
	"encoding/json"
)

// PatternDescription description of a tool pattern
type PatternDescription struct {
	PatternID   string             `json:"patternId"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Parameters  []PatternParameter `json:"parameters,omitempty"`
	TimeToFix   int                `json:"timeToFix"`
}

// ToJSON returns the json representation of the pattern description
func (i *PatternDescription) ToJSON() ([]byte, error) {
	return json.Marshal(i)
}

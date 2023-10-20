package codacytool

// Pattern corresponds to a rule (pattern) inspected by the tool.
type Pattern struct {
	ID          string             `json:"patternId"`
	Category    string             `json:"category,omitempty"`
	SubCategory string             `json:"subcategory,omitempty"`
	Level       string             `json:"level,omitempty"`
	Parameters  []PatternParameter `json:"parameters,omitempty"`
	Enabled     bool               `json:"enabled"`
}

// PatternParameter represents a parameter that can be used to customize a pattern.
type PatternParameter struct {
	Name        string      `json:"name"`
	Value       interface{} `json:"value,omitempty"`
	Default     interface{} `json:"default,omitempty"`
	Description string      `json:"description,omitempty"`
}

// PatternDescription represents the exhaustive description of a pattern.
type PatternDescription struct {
	PatternID   string             `json:"patternId"`
	Title       string             `json:"title"`
	Description string             `json:"description,omitempty"`
	Parameters  []PatternParameter `json:"parameters,omitempty"`
	TimeToFix   int                `json:"timeToFix,omitempty"`
}

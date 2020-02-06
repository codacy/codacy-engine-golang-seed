package metrics

// MetricResult result of a metric tool
type MetricResult struct {
	Filename         string           `json:"filename"`
	Complexity       int              `json:"complexity"`
	Loc              int              `json:"loc"`
	CLoc             int              `json:"cloc"`
	NrMethods        int              `json:"nrMethods"`
	NrClasses        int              `json:"nrClasses"`
	LineComplexities []LineComplexity `json:"lineComplexities"`
}

// LineComplexity is the complexity on a line
type LineComplexity struct {
	Line  int `json:"line"`
	Value int `json:"value"`
}

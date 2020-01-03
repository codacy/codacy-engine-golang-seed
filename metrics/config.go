package metrics

// MetricConfig is the configuration to run the metric tool
type MetricConfig struct {
	Files    []string `json:"files"`
	Language string   `json:"language"`
}

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
	SourceID   string `json:"sourceId,omitempty"`
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

// An "enum" representing the supported BOM formats.
type BomFormat string

const (
	// [CycloneDX] specification in JSON format.
	//
	// [CycloneDX]: https://cyclonedx.org/
	CycloneDXJSON = BomFormat("CycloneDXJSON")
)

// SBOM - Software Bill of Materials
//
// A SBOM declares the inventory of components used to build a software artifact, including any open source and
// proprietary software components.
type SBOM struct {
	// The format of the SBOM. Currently only [CycloneDX] specification in JSON format is supported.
	//
	// [CycloneDX]: https://cyclonedx.org/
	BomFormat BomFormat `json:"bomFormat"`
	// The version of the SBOM format used to build this SBOM.
	SpecVersion string `json:"specVersion"`
	// The actual SBOM content. To be parsed by downstream consumers according to `bomFormat` and `specVersion`.
	Sbom string `json:"sbom"`
}

func (s SBOM) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}

// GetFile always returns an empty value since SBOM is for the whole project, not a single file.
func (s SBOM) GetFile() string {
	return ""
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

package codacytool

const (
	testsResourcesLocation = "./tests/"
)

// ToolImplementationTest mock tool implementation
type ToolImplementationTest struct{}

// Run mock tool for testing run implementation
func (i ToolImplementationTest) Run(tool Tool, sourceDir string) ([]Issue, error) {
	issue := testIssue()
	return []Issue{issue}, nil
}

package codacytool

import (
	"github.com/stretchr/testify/assert"

	"testing"
	"time"
)

func TestTimeoutFinish(t *testing.T) {
	expectedError := "Timeout exceeded"

	callToolTimeoutMock := func() []Issue {
		time.Sleep(5 * time.Second)
		return []Issue{}
	}

	_, err := runToolCallWithTimeout(callToolTimeoutMock, 200*time.Millisecond)

	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err.Error())
}

func TestTimeoutNotEnd(t *testing.T) {
	expectedSuccess := []Issue{
		testIssue(),
	}

	callToolNoTimeoutMock := func() []Issue {
		return expectedSuccess
	}

	result, err := runToolCallWithTimeout(callToolNoTimeoutMock, 200*time.Millisecond)

	assert.Nil(t, err)
	assert.Equal(t, expectedSuccess, result)
}

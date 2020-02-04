package codacytool

import (
	"bytes"
	logrus "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"testing"
	"time"
)

func waitForIo(buf bytes.Buffer) string {
	for {
		l, err := buf.ReadString('\n')
		if err != nil {
			return l
		}
	}
}

func prepareLogging() *bytes.Buffer {
	var buf bytes.Buffer
	logrus.SetFormatter(&NoFormatter{})
	logrus.SetOutput(&buf)
	return &buf
}

func TestTimeoutFinish(t *testing.T) {
	buf := prepareLogging()
	expectedError := "Timeout of 1 seconds exceeded"

	callToolTimeoutMock := func() ([]Issue, error) {
		time.Sleep(5 * time.Second)
		return nil, nil
	}

	handleTimeoutMock := func() {
		logrus.Info(expectedError)
	}

	runToolWithTimeout(callToolTimeoutMock, handleTimeoutMock, 200*time.Millisecond)

	result := waitForIo(*buf)
	assert.Equal(t, expectedError, result)
}

func TestTimeoutNotEnd(t *testing.T) {
	buf := prepareLogging()
	expectedSuccess := "Run success"

	callToolNoTimeoutMock := func() ([]Issue, error) {
		logrus.Info(expectedSuccess)
		return nil, nil
	}

	handleTimeoutMock := func() {
		logrus.Info("not expected error")
	}

	runToolWithTimeout(callToolNoTimeoutMock, handleTimeoutMock, 200*time.Millisecond)

	result := waitForIo(*buf)
	assert.Equal(t, expectedSuccess, result)
}

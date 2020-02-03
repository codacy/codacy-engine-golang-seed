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

	runMethodWithTimeout(func() {
		time.Sleep(5 * time.Second)
	}, func() {
		logrus.Info(expectedError)
	}, 200*time.Millisecond)

	result := waitForIo(*buf)
	assert.Equal(t, expectedError, result)
}

func TestTimeoutNotEnd(t *testing.T) {
	buf := prepareLogging()
	expectedSuccess := "Run success"

	runMethodWithTimeout(func() {
		logrus.Info(expectedSuccess)
	}, func() {
		logrus.Info("not expected error")
	}, 200*time.Millisecond)

	result := waitForIo(*buf)
	assert.Equal(t, expectedSuccess, result)
}

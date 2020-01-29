package codacytool

import (
	logrus "github.com/sirupsen/logrus"
)

// NoFormatter Logrus custom formatter
type NoFormatter struct {
}

// Format returns the log message
func (f *NoFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(entry.Message), nil
}

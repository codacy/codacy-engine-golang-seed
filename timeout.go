package codacytool

import (
	"os"
	"strconv"
	"time"
)

func timeoutSeconds() time.Duration {
	value, exists := os.LookupEnv("TIMEOUT_SECONDS")

	if !exists {
		return defaultTimeout
	}

	seconds, err := strconv.Atoi(value)
	if err != nil {
		return defaultTimeout
	}

	return time.Duration(seconds) * time.Second
}

func runMethodWithTimeout(method func(), timeoutExceeded func(), maxDuration time.Duration) {
	c1 := make(chan string, 1)
	go func() {
		method()
		c1 <- ""
	}()
	select {
	case <-c1:
	case <-time.After(maxDuration):
		timeoutExceeded()
	}
}

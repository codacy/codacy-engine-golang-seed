package codacytool

import (
	"context"
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
	ctx := context.Background()

	c1 := make(chan struct{})
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	go func() {
		method()
		c1 <- struct{}{}
	}()
	select {
	case <-c1:
	case <-ctx.Done():
		timeoutExceeded()
	}
}

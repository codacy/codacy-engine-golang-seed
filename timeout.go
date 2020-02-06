package codacytool

import (
	"context"
	"errors"
	"strconv"
	"time"
)

func parseTimeoutSeconds(value string) time.Duration {
	seconds, err := strconv.Atoi(value)
	if err != nil {
		return defaultTimeout
	}

	return time.Duration(seconds) * time.Second
}

func runToolCallWithTimeout(method func() []Issue, maxDuration time.Duration) ([]Issue, error) {
	ctx := context.Background()
	c1 := make(chan []Issue, 1)
	ctx, cancel := context.WithTimeout(ctx, maxDuration)
	defer cancel()

	go func() {
		c1 <- method()
	}()

	select {
	case res := <-c1:
		return res, nil
	case <-ctx.Done():
		return nil, errors.New("Timeout exceeded")
	}
}

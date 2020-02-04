package codacytool

import (
	"context"
	"os"
	"strconv"
	"time"
)

type runToolMethod func() ([]Issue, error)
type timeoutExceededMethod func()
type toolRunResult struct {
	results []Issue
	err     error
}

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

func runToolWithTimeout(method runToolMethod, timeoutExceeded timeoutExceededMethod, maxDuration time.Duration) {
	ctx := context.Background()
	c1 := make(chan toolRunResult, 1)
	ctx, cancel := context.WithTimeout(ctx, maxDuration)
	defer cancel()

	go callTool(method, c1)

	select {
	case runResult := <-c1:
		printResult(runResult.results, runResult.err)
	case <-ctx.Done():
		timeoutExceeded()
	}
}

func callTool(method runToolMethod, c1 chan toolRunResult) {
	result, err := method()
	c1 <- toolRunResult{
		results: result,
		err:     err,
	}
}

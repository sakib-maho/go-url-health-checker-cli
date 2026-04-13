package checker

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/sakib-maho/go-url-health-checker-cli/internal/model"
)

func CheckURLs(urls []string, timeout time.Duration, retries int) []model.Result {
	results := make([]model.Result, len(urls))
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for i, u := range urls {
		go func(idx int, url string) {
			defer wg.Done()
			results[idx] = checkSingle(url, timeout, retries)
		}(i, u)
	}

	wg.Wait()
	return results
}

func checkSingle(url string, timeout time.Duration, retries int) model.Result {
	client := &http.Client{Timeout: timeout}
	var finalErr string

	for attempt := 0; attempt <= retries; attempt++ {
		start := time.Now()
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			cancel()
			return model.Result{URL: url, Error: err.Error()}
		}

		resp, err := client.Do(req)
		cancel()
		if err != nil {
			finalErr = err.Error()
			continue
		}
		_ = resp.Body.Close()
		return model.Result{
			URL:       url,
			Status:    resp.StatusCode,
			LatencyMS: time.Since(start).Milliseconds(),
		}
	}

	return model.Result{
		URL:   url,
		Error: finalErr,
	}
}

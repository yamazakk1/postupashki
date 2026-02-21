package handlers

import (
	"context"
	"errors"
	"fmt"
	"io"
	"maps"
	"net"
	"net/http"
	"time"
)

type Output struct {
	Status  int
	Headers map[string][]string
	Body    string
	From    string
}

func DoGet(ctx context.Context, timeout time.Duration, url string) (*Output, error) {
	if url == "" {
		return nil, errors.New("empty url")
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ResponseHeaderTimeout: timeout,
		DisableKeepAlives:     true,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, context.DeadlineExceeded
		}
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	output := Output{
		Status:  resp.StatusCode,
		Headers: make(map[string][]string),
		From:    url,
	}

	maps.Copy(output.Headers, resp.Header)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}
	output.Body = string(bodyBytes)

	return &output, nil
}

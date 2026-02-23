package handlers

import (
	"context"
	"io"
	"net/http"
)

type Output struct {
	Status  int
	Headers map[string][]string
	Body    string
	From    string
	Err     error
}

func DoGetWithContext(ctx context.Context, url string, resChan chan Output) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	output := Output{
		From: url,
		Err:  nil,
	}
	if err != nil {
		output.Err = err
		resChan <- output
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		output.Err = err
		resChan <- output
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		output.Err = err
		resChan <- output
		return
	}

	output.Status = resp.StatusCode
	output.Headers = resp.Header
	output.Body = string(body)

	if ctx.Err() != nil {
		output.Err = ctx.Err()
	}
	resChan <- output
}

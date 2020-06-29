package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

// Client ...
type Client struct {
	EndPointURL *url.URL
	HTTPClient  *http.Client
	UserAgent   string

	Logger *log.Logger
}

func newClient(endpointURL string, httpClient *http.Client, userAgent string, logger *log.Logger) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(endpointURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %s", endpointURL)
	}

	var discardLogger = log.New(ioutil.Discard, "", log.LstdFlags)
	if logger == nil {
		logger = discardLogger
	}

	client := &Client{
		EndPointURL: parsedURL,
		HTTPClient:  httpClient,
		UserAgent:   userAgent,
		Logger:      discardLogger,
	}
	return client, nil
}

func (c *Client) newRequest(ctx context.Context, method, subPath string, queries, headers map[string]string, reqBody io.Reader) (*http.Request, error) {
	reqURL := *c.EndPointURL
	reqURL.Path = path.Join(c.EndPointURL.Path, subPath)

	if queries != nil {
		q := reqURL.Query()
		for k, v := range queries {
			q.Add(k, v)
		}
		reqURL.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, reqURL.String(), reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	req = req.WithContext(ctx)

	return req, nil
}

func (c *Client) doRequest(req *http.Request, respBody interface{}) (int, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || 300 <= resp.StatusCode {
		return resp.StatusCode, nil
	}

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return 0, err
	}

	return resp.StatusCode, nil

}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)

}

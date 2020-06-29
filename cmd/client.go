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

func (cli *Client) newRequest(ctx context.Context, method string, subPath string, body io.Reader) (*http.Request, error) {
	endpointURL := *cli.EndPointURL
	endpointURL.Path = path.Join(cli.EndPointURL.Path, subPath)

	req, err := http.NewRequest(method, endpointURL.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", cli.UserAgent)

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)

}

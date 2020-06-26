package cmd

import (
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

func newMockServer() (*http.ServeMux, *url.URL) {
	mux := http.NewServeMux()
	// server := httptest.NewServer(mux)
	// mockServerURL, _ := url.Parse(server.URL)
	mockServerURL, _ := url.Parse(viper.GetString("url"))
	return mux, mockServerURL
}

func newTestClient(mockServerURL *url.URL) *Client {
	endpointURL := mockServerURL.String()
	httpClient := &http.Client{}
	userAgent := "test client"
	cli, _ := newClient(endpointURL, httpClient, userAgent)
	return cli
}

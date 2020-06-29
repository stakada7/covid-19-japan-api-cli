package cmd

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func newTestClient(mockServerURL *url.URL) (*Client, error) {
	endpointURL := mockServerURL.String()
	httpClient := &http.Client{}
	userAgent := "test client"
	return newClient(endpointURL, httpClient, userAgent, nil)

}

func setup(t *testing.T, mockResponseHeaderFile, mockResponseBodyFile, expectedMethod, expectedRequestPath, expectedRawQuery string) (*Client, func()) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != expectedMethod {
			t.Fatalf("request method wrong. want=%s, got=%s", expectedMethod, req.Method)
		}
		if req.URL.Path != expectedRequestPath {
			t.Fatalf("request path wrong. want=%s, got=%s", expectedRequestPath, req.URL.Path)
		}
		if req.URL.RawQuery != expectedRawQuery {
			t.Fatalf("request query wrong. want=%s, got=%s", expectedRawQuery, req.URL.RawQuery)
		}

		headerBytes, err := ioutil.ReadFile(mockResponseHeaderFile)
		if err != nil {
			t.Fatalf("failed to read header '%s': %s", mockResponseHeaderFile, err.Error())
		}
		firstLine := strings.Split(string(headerBytes), "\n")[0]

		statusCode, err := strconv.Atoi(strings.Fields(firstLine)[1])
		if err != nil {
			t.Fatalf("failed to extract status code from header: %s", err.Error())
		}
		w.WriteHeader(statusCode)

		bodyBytes, err := ioutil.ReadFile(mockResponseBodyFile)
		if err != nil {
			t.Fatalf("failed to read body '%s': %s", mockResponseBodyFile, err.Error())
		}
		w.Write(bodyBytes)
	}))

	serverURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("failed to get mock server URL: %s", err.Error())
	}

	cli := &Client{
		EndPointURL: serverURL,
		HTTPClient:  server.Client(),
		UserAgent:   "test client",
		Logger:      nil,
	}

	teardown := func() {
		server.Close()
	}

	return cli, teardown
}

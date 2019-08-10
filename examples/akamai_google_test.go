package leotest

import (
	"net/http"
	"strings"
	"testinfra-go/pkg/akamai"
	"testing"
	"time"
)

// This file contains examples of how to use terratest-like engine to test an Akamai configuration.
// There is one test:
// - TestAkamaiYooxDefaultCheck: An example of how to read in requested HTTP/s URI verifing the response headers.

// Those are the settings of the Akamai HTTP client
var settings = &akamai.Settings{
	TLSHandshakeTimeout:   10 * time.Second,
	Connect:               10 * time.Second,
	ResponseHeaderTimeout: 10 * time.Second,
	Timeout:               10 * time.Second,
	SkipSSLVerify:         false,
}

// An example of how verify the Akamai response headers of a requested URL using the
// Advanced Request Headers aka Pragma.
func TestAkamaiGoogleDefaultCheck(t *testing.T) {
	// checkHeader function verify if the response header value matches with the provided value
	checkHeader := func(t *testing.T, got, want string) {
		t.Helper()
		if !strings.Contains(got, want) {
			t.Errorf("Values got %q want %q", got, want)
		}
	}
	// url is the URL that must be verified
	url := "https://www.google.com"
	// Create an Akamai HTTP client
	c := akamai.NewClient(settings)
	// Create a new HTTP request with method GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	// GET the resource setting the Pragma request headers
	resp, err := c.GetWithRequestE(req, map[string]interface{}{
		"Pragma": "no-cache,akamai-x-cache-on,akamai-x-cache-remote-on,akamai-x-check-cacheable,akamai-x-get-cache-key,akamai-x-get-ssl-client-session-id,akamai-x-get-true-cache-key,akamai-x-serial-no,akamai-x-get-request-id,X-Akamai-CacheTrack",
	})
	defer resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	// Verify the status code looking for 200 OK
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Test failed. Request to the URL %s failed with code %d", url, resp.StatusCode)
	}
	tests := []struct {
		name   string
		header string
		want   string
	}{
		{name: "Check if resource is cachable", header: "X-Check-Cacheable", want: "YES"},
		{name: "Check internal cache key value", header: "X-Cache-Key", want: "S/L/6666/8888/1h/www.google.com"},
		{name: "Check the CP code value", header: "X-Cache-Key", want: "8888"},
		{name: "Check the Cache TTL value", header: "X-Cache-Key", want: "1h"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkHeader(t, resp.Header.Get(tt.header), tt.want)
		})
	}
}

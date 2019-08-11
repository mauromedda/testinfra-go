package tests

import (
	"net/http"
	"testinfra-go/pkg/akamai"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// This file contains examples of how to use terratest-like engine to test an Akamai configuration.
// There is one test:
// - TestAkamaiCacheHeaderCheck: An example of how to read in requested HTTP/s URI verifing the response headers
//   X-Cache-Key and X-Check-Cacheable.
// - TestAkamaiCacheKey: An example of how to verify the CP code, TTL, Origin associated to a specific URI.
// - TestAkamaiCacheKeyValues: An example of how to verify the Akamai Internal Cache value using the XCacheKey struct.

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
func TestAkamaiCacheHeaderCheck(t *testing.T) {
	t.Parallel()
	// url is the URL that must be verified
	url := "https://www.mockorig.com/it/donna"
	// Create an Akamai HTTP client
	c := akamai.NewClient(settings)
	// Create a new HTTP request with method GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	// GET the resource setting the Pragma request headers
	resp, err := c.GetWithRequestE(req, map[string]string{
		"Pragma": "no-cache,akamai-x-cache-on,akamai-x-cache-remote-on,akamai-x-check-cacheable,akamai-x-get-cache-key,akamai-x-get-ssl-client-session-id,akamai-x-get-true-cache-key,akamai-x-serial-no,akamai-x-get-request-id,X-Akamai-CacheTrack",
	})
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

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
		{name: "Check internal cache key value", header: "X-Cache-Key", want: "S/L/8888/666666/3h/www.mockorig.com/it/donna"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, resp.Header.Get(tt.header))
		})
	}
}

// TestAkamaiCacheKey: An example of how to verify the CP code, TTL, Origin associated to a specific URI.
func TestAkamaiCacheKey(t *testing.T) {
	t.Parallel()
	// url is the URL that must be verified
	url := "https://www.mockorig.com/it/donna"
	// Create an Akamai HTTP client
	c := akamai.NewClient(settings)
	// Create a new HTTP request with method GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	// GET the resource setting the Pragma request headers
	resp, err := c.GetWithRequestE(req, map[string]string{
		"Pragma": "no-cache,akamai-x-cache-on,akamai-x-cache-remote-on,akamai-x-check-cacheable,akamai-x-get-cache-key,akamai-x-get-ssl-client-session-id,akamai-x-get-true-cache-key,akamai-x-serial-no,akamai-x-get-request-id,X-Akamai-CacheTrack",
	})
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	// Verify the status code looking for 200 OK
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Test failed. Request to the URL %s failed with code %d", url, resp.StatusCode)
	}

	// Expected CP Code value
	expectedCPCode := 666666

	// Expected TTL value
	expectedTTL := "3h"

	// Expected origin
	expectedOrigin := "www.mockorig.com/it/donna"

	// Map the header to the XCacheKey struct
	gotXCacheKey, err := akamai.XCacheKeyUnmarshal(resp.Header.Get("X-Cache-Key"))
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Verify CP value", func(t *testing.T) {
		assert.Equal(t, expectedCPCode, gotXCacheKey.GetCP())
	})

	t.Run("Verify Origin value", func(t *testing.T) {
		assert.Equal(t, expectedOrigin, gotXCacheKey.GetOrigin())
	})

	t.Run("Verify the TTL value", func(t *testing.T) {
		assert.Equal(t, expectedTTL, gotXCacheKey.GetTTL())

	})
}

// - TestAkamaiCacheKeyValues: An example of how to verify the Akamai Internal Cache value using the XCacheKey struct.
func TestAkamaiCacheKeyValues(t *testing.T) {
	t.Parallel()
	// url is the URL that must be verified
	url := "https://www.mockorig.com/it/donna"
	// Create an Akamai HTTP client
	c := akamai.NewClient(settings)
	// Create a new HTTP request with method GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	// GET the resource setting the Pragma request headers
	resp, err := c.GetWithRequestE(req, map[string]string{
		"Pragma": "no-cache,akamai-x-cache-on,akamai-x-cache-remote-on,akamai-x-check-cacheable,akamai-x-get-cache-key,akamai-x-get-ssl-client-session-id,akamai-x-get-true-cache-key,akamai-x-serial-no,akamai-x-get-request-id,X-Akamai-CacheTrack",
	})
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	// Verify the status code looking for 200 OK
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Test failed. Request to the URL %s failed with code %d", url, resp.StatusCode)
	}

	// Expected CP Code value
	expectedCPCode := 666666

	// Expected TTL value
	expectedTTL := "3h"

	// Expected origin
	expectedOrigin := "www.mockorig.com/it/donna"

	// expectedXCacheKeyHeader
	expectedXCacheKeyHeader := &akamai.XCacheKey{
		SecureDeliveryIndicator: "S",
		TypeCode:                "L",
		Serial:                  8888,
		CPCode:                  expectedCPCode,
		TTL:                     expectedTTL,
		FwdPath:                 expectedOrigin,
	}
	// Map the header to the XCacheKey struct
	gotXCacheKey, err := akamai.XCacheKeyUnmarshal(resp.Header.Get("X-Cache-Key"))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedXCacheKeyHeader, gotXCacheKey)
}

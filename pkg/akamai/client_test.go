package akamai

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWithRequestE(t *testing.T) {

	// url is the URL that must be verified
	url := "http://localhost:8080/it/donna"
	// Create an Akamai HTTP client
	c := NewClient(settings)
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

	assert.Equal(t, 200, resp.StatusCode)

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

func TestGetE(t *testing.T) {

	// url is the URL that must be verified
	url := "http://localhost:8080/it/donna"
	// Create an Akamai HTTP client
	c := NewClient(settings)
	// Create a new HTTP request with the simpliyf get method.
	resp, err := c.GetE(url)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)

	tests := []struct {
		name   string
		header string
		want   string
	}{
		{name: "Check if Akamai Header X-Check-Cacheable is absent", header: "X-Check-Cacheable", want: ""},
		{name: "Check internal cache key Header is absent", header: "X-Cache-Key", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, resp.Header.Get(tt.header))
		})
	}
}

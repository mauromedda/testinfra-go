package akamai

import (
	"crypto/tls"
	"net/http"
	"time"
)

// Settings contains an Akamai HTTP Cient configuration.
type Settings struct {
	TLSHandshakeTimeout   time.Duration
	Connect               time.Duration
	ResponseHeaderTimeout time.Duration
	Timeout               time.Duration
	SkipSSLVerify         bool
}

// Default returns a HTTP/s Client settings initialized with sensible defaults.
func (s *Settings) Default() *Settings {
	return &Settings{
		TLSHandshakeTimeout:   10 * time.Second,
		Connect:               10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		Timeout:               10 * time.Second,
		SkipSSLVerify:         false,
	}
}

// Client is an HTTP client.
// It wraps net/http's client and add some methods for making HTTP request easier.
type Client struct {
	client *http.Client
}

// NewClient return a custom http client
func NewClient(s *Settings) *Client {
	return &Client{client: &http.Client{
		Timeout: s.Timeout,
		Transport: &http.Transport{
			TLSHandshakeTimeout:   s.TLSHandshakeTimeout,
			ResponseHeaderTimeout: s.ResponseHeaderTimeout,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: s.SkipSSLVerify,
			},
		},
	},
	}
}

// GetE issues a GET to the specified URL. It returns an http.Response for further processing and any error.
func (c *Client) GetE(url string) (*http.Response, error) {
	return c.client.Get(url)
}

// GetWithRequestE issues a GET to the specified URL. It returns an http.Response for further processing and any error.
func (c *Client) GetWithRequestE(req *http.Request, hearders map[string]string) (*http.Response, error) {
	for k, v := range hearders {
		req.Header.Set(k, v)
	}
	return c.client.Do(req)
}

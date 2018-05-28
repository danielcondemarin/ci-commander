package teamcitytest

import (
	"net/http"
)

type whenDo func(*http.Request) (*http.Response, error)

type whenRead func(p []byte) (int, error)

type whenClose func() error

// mockHTTPClient is a partial mock implementation of http Client
type mockHTTPClient struct {
	WhenDo whenDo
}

// NewMockHTTPClient is the factory method for creating instances of mockHTTPClient
func NewMockHTTPClient() *mockHTTPClient {
	c := &mockHTTPClient{}
	c.WhenDo = func(r *http.Request) (*http.Response, error) { return nil, nil }
	return c
}

// mockReadCloser is a mock implementation for the interface ReadCloser
type mockReadCloser struct {
	WhenRead  whenRead
	WhenClose whenClose
}

// NewMockReadCloser is the factory method for creating instances of mockReadCloser
func NewMockReadCloser() *mockReadCloser {
	c := &mockReadCloser{}
	c.WhenRead = func(p []byte) (int, error) { return 1, nil }
	c.WhenClose = func() error { return nil }
	return c
}

// Read is a proxy to WhenRead
func (rc mockReadCloser) Read(p []byte) (int, error) {
	return rc.WhenRead(p)
}

// Close is a proxy to WhenClose
func (rc mockReadCloser) Close() error {
	return rc.WhenClose()
}

// Do is a proxy to WhenDo
func (c mockHTTPClient) Do(r *http.Request) (*http.Response, error) {
	return c.WhenDo(r)
}

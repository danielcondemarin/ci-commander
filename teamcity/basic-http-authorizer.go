package teamcity

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

// BasicHTTPAuthorizer is the basic HTTP Authorization
type BasicHTTPAuthorizer struct {
	username string
	password string
}

// Authorize implements basic HTTP Authorize by adding the Authorization HTTP Header to the request with the credentials provided in the basic form
func (a BasicHTTPAuthorizer) Authorize(r *http.Request) error {
	base64EncodedCredentials := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", a.username, a.password)))
	uriEncodedAuth := fmt.Sprintf("Basic %s", base64EncodedCredentials)
	r.Header.Add("Authorization", uriEncodedAuth)
	return nil
}

// NewBasicHTTPAuthorizer is the factory method for instances of BasicHTTPAuthorizer
func NewBasicHTTPAuthorizer(username string, password string) *BasicHTTPAuthorizer {
	return &BasicHTTPAuthorizer{username: username, password: password}
}

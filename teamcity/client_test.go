package teamcity

import (
	"encoding/base64"
	"net/http"
	"strings"
	"testing"

	"github.com/danielcondemarin/go-ci-commander/teamcity/teamcitytest"
)

func TestBasicAuthorization(t *testing.T) {
	httpClient := teamcitytest.NewMockHTTPClient()
	auth := NewBasicHTTPAuthorizer("gouser", "gopass")
	client := &client{url: "http://teamcityserver", client: httpClient, authorizer: auth}

	httpClient.WhenDo = func(r *http.Request) (*http.Response, error) {
		authHeaderVal := r.Header.Get("Authorization")
		encCreds := strings.Split(authHeaderVal, " ")[1]

		authHeaderDec, err := base64.StdEncoding.DecodeString(encCreds)

		if err != nil {
			t.Error(err)
		}

		if string(authHeaderDec) != "gouser:gopass" {
			t.Error("Expected Authorization Header to contain correct credentials")
		}

		body := teamcitytest.NewMockReadCloser()

		return &http.Response{Body: body}, nil
	}

	client.makeGETRequest("/foo/bar", nil)
}

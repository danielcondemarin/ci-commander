package teamcity

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/danielcondemarin/go-ci-commander/teamcity/teamcitytest"
)

func TestGetProjects(t *testing.T) {
	getProjectsPath := "/app/rest/projects"

	httpClient := teamcitytest.NewMockHTTPClient()
	auth := NewBasicHTTPAuthorizer("gouser", "gopass")
	client := &client{url: "http://teamcity", client: httpClient, authorizer: auth}

	body := teamcitytest.NewMockReadCloser()

	body.WhenRead = func(p []byte) (int, error) {
		projects := []project{project{ID: "1"}}
		tbRespStrB, _ := json.Marshal(projectsResponse{Projects: projects})
		copy(p, tbRespStrB)
		return len(p), nil
	}

	httpClient.WhenDo = func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != getProjectsPath {
			t.Errorf("Expected request URL to be %s but was %s", getProjectsPath, r.URL.Path)
		}

		return &http.Response{Body: body}, nil
	}

	res, err := client.GetProjects()

	if err != nil {
		t.Error(err)
	}

	if len(res.Projects) != 1 {
		t.Error("Expected 1 project in the result")
	}

}

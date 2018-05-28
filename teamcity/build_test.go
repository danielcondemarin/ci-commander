package teamcity

import (
	"encoding/json"
	"net/http"
	"testing"

	teamcitytest "github.com/danielcondemarin/go-ci-commander/teamcity/teamcitytest"
)

const (
	TestQueuedID = 1
)

func TestNewBuildRequest(t *testing.T) {
	buildID := "gopherbuild"
	bReq := NewBuildRequest(buildID)

	if bReq.BuildType.ID != buildID {
		t.Error("Wrong buildTypeID found on TriggerBuildRequest object")
	}
}

func TestTriggerBuild(t *testing.T) {
	triggerBuildPath := "/app/rest/buildQueue"
	httpClient := teamcitytest.NewMockHTTPClient()
	auth := NewBasicHTTPAuthorizer("gouser", "gopass")

	client := &client{url: "http://teamcityserver", client: httpClient, authorizer: auth}

	body := teamcitytest.NewMockReadCloser()

	body.WhenRead = func(p []byte) (int, error) {
		tbRespStrB, _ := json.Marshal(triggerBuildResponse{ID: TestQueuedID})
		copy(p, tbRespStrB)
		return len(p), nil
	}

	httpClient.WhenDo = func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != triggerBuildPath {
			t.Errorf("Expected request URL to be %s but was %s", triggerBuildPath, r.URL.Path)
		}

		return &http.Response{Body: body}, nil
	}

	bReq := NewBuildRequest("TestBuild")
	res, err := client.TriggerBuild(bReq)

	if err != nil {
		t.Error(err)
	}

	if res.ID != 1 {
		t.Errorf("Expected TriggerBuildResponse ID %d", TestQueuedID)
	}
}

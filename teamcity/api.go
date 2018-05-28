package teamcity

import (
	"net/http"
)

// Authorizer is the interface for different types of Authorization (ex. BasicHTTPAuthorizer)
type Authorizer interface {
	Authorize(*http.Request) error
}

// Client is the main API Interface
type Client interface {
	TriggerBuild(triggerBuildRequest) (triggerBuildResponse, error)
	GetProjects() (projectsResponse, error)
}

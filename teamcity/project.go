package teamcity

// ProjectsResponse is the response returned from GetProjects
type projectsResponse struct {
	Projects []project `json:"project"`
}

type project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Href        string `json:"href"`
	WebURL      string `json:"webUrl"`
}

func (c client) GetProjects() (projectsResponse, error) {
	var result projectsResponse
	err := c.makeGETRequest("/app/rest/projects", &result)

	if err != nil {
		return projectsResponse{}, err
	}

	return result, nil
}

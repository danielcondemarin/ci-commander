package teamcity

type triggerBuildRequest struct {
	BuildType buildType `json:"buildType"`
}

type triggerBuildResponse struct {
	ID int `json:"id"`
}

type buildType struct {
	ID string `json:"id"`
}

// NewBuildRequest is the factory method forcreating instances of TriggerBuildRequest
func NewBuildRequest(id string) triggerBuildRequest {
	return triggerBuildRequest{BuildType: buildType{ID: id}}
}

func (c client) TriggerBuild(r triggerBuildRequest) (triggerBuildResponse, error) {
	var result triggerBuildResponse
	err := c.makePOSTRequest("/app/rest/buildQueue", r, &result)

	if err != nil {
		return triggerBuildResponse{}, err
	}

	return result, nil
}

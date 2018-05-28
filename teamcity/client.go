package teamcity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type client struct {
	url        string
	client     httpClient
	authorizer Authorizer
}

// NewClient is the factory for the Client interface
func NewClient(url string, auth Authorizer) Client {
	c := client{url: url, client: &http.Client{}, authorizer: auth}
	return Client(c)
}

func addCommonRequestHeaders(r *http.Request) {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
}

func (c client) makeGETRequest(relativeURL string, result interface{}) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.url, relativeURL), nil)

	if err != nil {
		return err
	}

	addCommonRequestHeaders(req)
	c.authorizer.Authorize(req)

	resp, err := c.client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)

	return dec.Decode(result)
}

func (c client) makePOSTRequest(relativeURL string, body interface{}, result interface{}) error {
	requestBody, err := json.Marshal(body)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.url, relativeURL), bytes.NewBuffer(requestBody))

	if err != nil {
		return err
	}

	addCommonRequestHeaders(req)
	c.authorizer.Authorize(req)

	resp, err := c.client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)

	return dec.Decode(result)
}

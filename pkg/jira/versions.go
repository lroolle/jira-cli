package jira

import (
	"context"
	"encoding/json"
	"net/http"
)

// Version represents a Jira version.
type Version struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Released    bool   `json:"released"`
	ReleaseDate string `json:"releaseDate"`
}

// GetProjectVersions fetches the versions for a given project.
func (c *Client) GetProjectVersions(projectKey string) ([]*Version, error) {
	res, err := c.GetV2(context.Background(), "/project/"+projectKey+"/versions", nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, ErrEmptyResponse
	}
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != http.StatusOK {
		return nil, formatUnexpectedResponse(res)
	}

	var out []*Version
	err = json.NewDecoder(res.Body).Decode(&out)

	return out, err
}

package awx

import (
	"bytes"
	"encoding/json"
)

// GroupService implements awx Groups apis.
type GroupService struct {
	client *Client
}

// ListGroupsResponse represents `ListGroups` endpoint response.
type ListGroupsResponse struct {
	Pagination
	Results []*Group `json:"results"`
}

// ListGroups shows list of awx Groups.
func (u *GroupService) ListGroups(params map[string]string) ([]*Group, *ListGroupsResponse, error) {
	result := new(ListGroupsResponse)
	endpoint := "/api/v2/Groups/"
	resp, err := u.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateGroup creates an awx Group.
func (u *GroupService) CreateGroup(data map[string]interface{}, params map[string]string) (*Group, error) {
	result := new(Group)
	endpoint := "/api/v2/Groups/"
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Add check if Group exists and return proper error

	resp, err := u.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"training-plan/internal/api/transport/request"
)

type APIClient struct {
	baseURL string
	client  *http.Client
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (c *APIClient) CreateUser(username string) error {
	body := bytes.NewBufferString(fmt.Sprintf(`{"username":"%s"}`, username))

	resp, err := c.client.Post(c.baseURL+"/users/create", "application/json", body)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *APIClient) GetUser(username string) (*http.Response, error) {
	return c.client.Get(fmt.Sprintf("%s/users/%s", c.baseURL, username))
}

func (c *APIClient) CreateRunningProfile(req request.CreateRunningProfileRequest) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := c.client.Post(c.baseURL+"/running-profiles/create", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

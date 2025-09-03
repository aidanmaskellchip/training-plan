package client

import (
	"bytes"
	"fmt"
	"net/http"
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
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create user: %s", resp.Status)
	}

	return nil
}

func (c *APIClient) GetUser(username string) (*http.Response, error) {
	return c.client.Get(fmt.Sprintf("%s/users/%s", c.baseURL, username))
}

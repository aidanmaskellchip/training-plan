package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

type OpenAIClient struct {
	baseURL string
	apiKey  string
	model   string
	client  http.Client
}

func NewOpenAIClient(baseURL, key string) *OpenAIClient {
	return &OpenAIClient{
		baseURL: baseURL,
		apiKey:  key,
		client: http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (oc OpenAIClient) Chat(model, prompt string) (response string, err error) {
	request := ChatRequest{
		Model: model,
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}
	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("Marshal: %v\n", err)
	}

	chatURL, err := url.JoinPath(oc.baseURL, "/chat/completions")
	if err != nil {
		return "", fmt.Errorf("JoinPath: %w", err)
	}

	req, err := http.NewRequest("POST", chatURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("NewRequest: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+oc.apiKey)

	resp, err := oc.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Do: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ReadAll: %w", err)
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return "", fmt.Errorf("Unmarshal: %w", err)
	}

	return chatResp.Choices[0].Message.Content, nil
}

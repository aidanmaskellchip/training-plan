package tools

import (
	"context"
	"fmt"
	"github.com/pontus-devoteam/agent-sdk-go/pkg/tool"
	"training-plan/internal/ai/client"
)

func NewCreateUserTool(apiClient *client.APIClient) *tool.FunctionTool {
	return tool.NewFunctionTool(
		"create_user",
		"Create a new user",
		func(ctx context.Context, params map[string]interface{}) (interface{}, error) {
			username := params["username"].(string)

			err := apiClient.CreateUser(username)
			if err != nil {
				return nil, fmt.Errorf("failed to create user: %w", err)
			}

			return fmt.Sprintf("User: %s created successfully.", username), nil
		},
	).WithSchema(loadSchemaFromFile("./api_specs/create_user_schema.yaml", "/v1/users/create", "post"))
}

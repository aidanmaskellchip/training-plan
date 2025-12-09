package tools

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func loadSchemaFromFile(filepath, path, method string) map[string]interface{} {
	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(fmt.Errorf("failed to read schema file: %w", err))
	}

	var openAPISpec map[string]interface{}
	if err := yaml.Unmarshal(data, &openAPISpec); err != nil {
		panic(fmt.Errorf("failed to parse schema file: %w", err))
	}

	// Navigate through the OpenAPI structure safely
	paths, ok := openAPISpec["paths"].(map[string]interface{})
	if !ok {
		panic(fmt.Errorf("invalid OpenAPI spec: missing paths"))
	}

	targetPath, ok := paths[path].(map[string]interface{})
	if !ok {
		panic(fmt.Errorf("invalid OpenAPI spec: missing path %s", path))
	}

	methodSpec, ok := targetPath[method].(map[string]interface{})
	if !ok {
		panic(fmt.Errorf("invalid OpenAPI spec: missing %s method for path %s", method, path))
	}

	requestBody, ok := methodSpec["requestBody"].(map[string]interface{})
	if !ok {
		panic(fmt.Errorf("invalid OpenAPI spec: missing requestBody for %s %s", method, path))
	}

	content, ok := requestBody["content"].(map[string]interface{})
	if !ok {
		panic(fmt.Errorf("invalid OpenAPI spec: missing content"))
	}

	applicationJSON, ok := content["application/json"].(map[string]interface{})
	if !ok {
		panic(fmt.Errorf("invalid OpenAPI spec: missing application/json"))
	}

	schema, ok := applicationJSON["schema"].(map[string]interface{})
	if !ok {
		panic(fmt.Errorf("invalid OpenAPI spec: missing schema"))
	}

	return schema
}

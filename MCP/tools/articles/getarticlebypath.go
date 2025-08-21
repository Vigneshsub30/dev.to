package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/forem-api-v1/mcp-server/config"
	"github.com/forem-api-v1/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetarticlebypathHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		usernameVal, ok := args["username"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: username"), nil
		}
		username, ok := usernameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: username"), nil
		}
		slugVal, ok := args["slug"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: slug"), nil
		}
		slug, ok := slugVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: slug"), nil
		}
		url := fmt.Sprintf("%s/api/articles/%s/%s", cfg.BaseURL, username, slug)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("api-key", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGetarticlebypathTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_api_articles_username_slug",
		mcp.WithDescription("Published article by path"),
		mcp.WithString("username", mcp.Required(), mcp.Description("")),
		mcp.WithString("slug", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetarticlebypathHandler(cfg),
	}
}

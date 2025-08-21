package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/forem-api-v1/mcp-server/config"
	"github.com/forem-api-v1/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Put_api_pages_idHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Page
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/api/pages/%s", cfg.BaseURL, id)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.Page
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

func CreatePut_api_pages_idTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_api_pages_id",
		mcp.WithDescription("update details for a page"),
		mcp.WithNumber("id", mcp.Required(), mcp.Description("The ID of the page.")),
		mcp.WithString("title", mcp.Required(), mcp.Description("Input parameter: Title of the page")),
		mcp.WithString("body_json", mcp.Description("Input parameter: For JSON pages, the JSON body")),
		mcp.WithString("body_markdown", mcp.Description("Input parameter: The text (in markdown) of the ad (required)")),
		mcp.WithString("description", mcp.Required(), mcp.Description("Input parameter: For internal use, helps similar pages from one another")),
		mcp.WithBoolean("is_top_level_path", mcp.Description("Input parameter: If true, the page is available at '/{slug}' instead of '/page/{slug}', use with caution")),
		mcp.WithString("slug", mcp.Required(), mcp.Description("Input parameter: Used to link to this page in URLs, must be unique and URL-safe")),
		mcp.WithObject("social_image", mcp.Description("")),
		mcp.WithString("template", mcp.Required(), mcp.Description("Input parameter: Controls what kind of layout the page is rendered in")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Put_api_pages_idHandler(cfg),
	}
}

package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/forem-api-v1/mcp-server/config"
	"github.com/forem-api-v1/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetarticlesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		if val, ok := args["per_page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("per_page=%v", val))
		}
		if val, ok := args["tag"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tag=%v", val))
		}
		if val, ok := args["tags"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tags=%v", val))
		}
		if val, ok := args["tags_exclude"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tags_exclude=%v", val))
		}
		if val, ok := args["username"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("username=%v", val))
		}
		if val, ok := args["state"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("state=%v", val))
		}
		if val, ok := args["top"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("top=%v", val))
		}
		if val, ok := args["collection_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("collection_id=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/api/articles%s", cfg.BaseURL, queryString)
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
		var result []ArticleIndex
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

func CreateGetarticlesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_api_articles",
		mcp.WithDescription("Published articles"),
		mcp.WithNumber("page", mcp.Description("Pagination page")),
		mcp.WithNumber("per_page", mcp.Description("Page size (the number of items to return per page). The default maximum value can be overridden by \"API_PER_PAGE_MAX\" environment variable.")),
		mcp.WithString("tag", mcp.Description("Using this parameter will retrieve articles that contain the requested tag. Articles\nwill be ordered by descending popularity.This parameter can be used in conjuction with `top`.")),
		mcp.WithString("tags", mcp.Description("Using this parameter will retrieve articles with any of the comma-separated tags.\nArticles will be ordered by descending popularity.")),
		mcp.WithString("tags_exclude", mcp.Description("Using this parameter will retrieve articles that do _not_ contain _any_\nof comma-separated tags. Articles will be ordered by descending popularity.")),
		mcp.WithString("username", mcp.Description("Using this parameter will retrieve articles belonging\n            to a User or Organization ordered by descending publication date.\n            If `state=all` the number of items returned will be `1000` instead of the default `30`.\n            This parameter can be used in conjuction with `state`.")),
		mcp.WithString("state", mcp.Description("Using this parameter will allow the client to check which articles are fresh or rising.\n            If `state=fresh` the server will return fresh articles.\n            If `state=rising` the server will return rising articles.\n            This param can be used in conjuction with `username`, only if set to `all`.")),
		mcp.WithNumber("top", mcp.Description("Using this parameter will allow the client to return the most popular articles\nin the last `N` days.\n`top` indicates the number of days since publication of the articles returned.\nThis param can be used in conjuction with `tag`.")),
		mcp.WithNumber("collection_id", mcp.Description("Adding this will allow the client to return the list of articles\nbelonging to the requested collection, ordered by ascending publication date.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetarticlesHandler(cfg),
	}
}

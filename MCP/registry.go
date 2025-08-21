package main

import (
	"github.com/forem-api-v1/mcp-server/config"
	"github.com/forem-api-v1/mcp-server/models"
	tools_users "github.com/forem-api-v1/mcp-server/tools/users"
	tools_profile_images "github.com/forem-api-v1/mcp-server/tools/profile_images"
	tools_articles "github.com/forem-api-v1/mcp-server/tools/articles"
	tools_reactions "github.com/forem-api-v1/mcp-server/tools/reactions"
	tools_readinglist "github.com/forem-api-v1/mcp-server/tools/readinglist"
	tools_organizations "github.com/forem-api-v1/mcp-server/tools/organizations"
	tools_display_ads "github.com/forem-api-v1/mcp-server/tools/display_ads"
	tools_followers "github.com/forem-api-v1/mcp-server/tools/followers"
	tools_followed_tags "github.com/forem-api-v1/mcp-server/tools/followed_tags"
	tools_pages "github.com/forem-api-v1/mcp-server/tools/pages"
	tools_comments "github.com/forem-api-v1/mcp-server/tools/comments"
	tools_tags "github.com/forem-api-v1/mcp-server/tools/tags"
	tools_podcast_episodes "github.com/forem-api-v1/mcp-server/tools/podcast_episodes"
	tools_videos "github.com/forem-api-v1/mcp-server/tools/videos"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_users.CreateUnpublishuserTool(cfg),
		tools_profile_images.CreateGetprofileimageTool(cfg),
		tools_users.CreateGetusermeTool(cfg),
		tools_articles.CreateGetuserallarticlesTool(cfg),
		tools_reactions.CreatePost_api_reactions_toggleTool(cfg),
		tools_readinglist.CreateGetreadinglistTool(cfg),
		tools_organizations.CreateGetorgusersTool(cfg),
		tools_users.CreateSuspenduserTool(cfg),
		tools_display_ads.CreateGet_api_display_ads_idTool(cfg),
		tools_display_ads.CreatePut_api_display_ads_idTool(cfg),
		tools_followers.CreateGetfollowersTool(cfg),
		tools_followed_tags.CreateGetfollowedtagsTool(cfg),
		tools_articles.CreateUnpublisharticleTool(cfg),
		tools_articles.CreateGetuserarticlesTool(cfg),
		tools_articles.CreateGetuserpublishedarticlesTool(cfg),
		tools_pages.CreateDelete_api_pages_idTool(cfg),
		tools_pages.CreateGet_api_pages_idTool(cfg),
		tools_pages.CreatePut_api_pages_idTool(cfg),
		tools_users.CreateGetuserTool(cfg),
		tools_articles.CreateCreatearticleTool(cfg),
		tools_articles.CreateGetarticlesTool(cfg),
		tools_organizations.CreateGetorganizationTool(cfg),
		tools_articles.CreateGetlatestarticlesTool(cfg),
		tools_pages.CreateGet_api_pagesTool(cfg),
		tools_pages.CreatePost_api_pagesTool(cfg),
		tools_articles.CreateUpdatearticleTool(cfg),
		tools_articles.CreateGetarticlebyidTool(cfg),
		tools_articles.CreateGetarticlebypathTool(cfg),
		tools_comments.CreateGetcommentbyidTool(cfg),
		tools_tags.CreateGettagsTool(cfg),
		tools_podcast_episodes.CreateGetpodcastepisodesTool(cfg),
		tools_display_ads.CreatePut_api_display_ads_id_unpublishTool(cfg),
		tools_videos.CreateVideosTool(cfg),
		tools_users.CreatePostadminuserscreateTool(cfg),
		tools_display_ads.CreateGet_api_display_adsTool(cfg),
		tools_display_ads.CreatePost_api_display_adsTool(cfg),
		tools_comments.CreateGetcommentsbyarticleidTool(cfg),
		tools_reactions.CreatePost_api_reactionsTool(cfg),
		tools_articles.CreateGetuserunpublishedarticlesTool(cfg),
		tools_organizations.CreateGetorgarticlesTool(cfg),
	}
}

package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Comment represents the Comment schema from the OpenAPI specification
type Comment struct {
	Image_url string `json:"image_url,omitempty"` // Podcast image url
	Type_of string `json:"type_of,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Id_code string `json:"id_code,omitempty"`
}

// DisplayAd represents the DisplayAd schema from the OpenAPI specification
type DisplayAd struct {
	Approved bool `json:"approved,omitempty"` // Ad must be both published and approved to be in rotation
	Creator_id int `json:"creator_id,omitempty"` // Identifies the user who created the ad.
	Display_to string `json:"display_to,omitempty"` // Potentially limits visitors to whom the ad is visible
	Id int `json:"id,omitempty"` // The ID of the Display Ad
	Name string `json:"name"` // For internal use, helps distinguish ads from one another
	Organization_id int `json:"organization_id,omitempty"` // Identifies the organization to which the ad belongs
	Published bool `json:"published,omitempty"` // Ad must be both published and approved to be in rotation
	Article_exclude_ids string `json:"article_exclude_ids,omitempty"` // Articles this ad should *not* appear on (blank means no articles are disallowed, and this ad can appear next to any/all articles). Comma-separated list of integer Article IDs
	Placement_area string `json:"placement_area"` // Identifies which area of site layout the ad can appear in
	Body_markdown string `json:"body_markdown"` // The text (in markdown) of the ad (required)
	Tag_list string `json:"tag_list,omitempty"` // Tags on which this ad can be displayed (blank is all/any tags)
	Type_of string `json:"type_of,omitempty"` // Types of the billboards: in_house (created by admins), community (created by an entity, appears on entity's content), external ( created by an entity, or a non-entity, can appear everywhere)
}

// SharedPodcast represents the SharedPodcast schema from the OpenAPI specification
type SharedPodcast struct {
	Slug string `json:"slug,omitempty"`
	Title string `json:"title,omitempty"`
	Image_url string `json:"image_url,omitempty"` // Podcast image url
}

// UserInviteParam represents the UserInviteParam schema from the OpenAPI specification
type UserInviteParam struct {
	Name string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// ProfileImage represents the ProfileImage schema from the OpenAPI specification
type ProfileImage struct {
	Image_of string `json:"image_of,omitempty"` // Determines the type of the profile image owner (user or organization)
	Profile_image string `json:"profile_image,omitempty"` // Profile image (640x640)
	Profile_image_90 string `json:"profile_image_90,omitempty"` // Profile image (90x90)
	Type_of string `json:"type_of,omitempty"` // Return profile_image
}

// SharedOrganization represents the SharedOrganization schema from the OpenAPI specification
type SharedOrganization struct {
	Profile_image_90 string `json:"profile_image_90,omitempty"` // Profile image (90x90)
	Slug string `json:"slug,omitempty"`
	Username string `json:"username,omitempty"`
	Name string `json:"name,omitempty"`
	Profile_image string `json:"profile_image,omitempty"` // Profile image (640x640)
}

// SharedUser represents the SharedUser schema from the OpenAPI specification
type SharedUser struct {
	Username string `json:"username,omitempty"`
	Website_url string `json:"website_url,omitempty"`
	Github_username string `json:"github_username,omitempty"`
	Name string `json:"name,omitempty"`
	Profile_image string `json:"profile_image,omitempty"` // Profile image (640x640)
	Profile_image_90 string `json:"profile_image_90,omitempty"` // Profile image (90x90)
	Twitter_username string `json:"twitter_username,omitempty"`
}

// ArticleIndex represents the ArticleIndex schema from the OpenAPI specification
type ArticleIndex struct {
	Id int `json:"id"`
	Public_reactions_count int `json:"public_reactions_count"`
	Canonical_url string `json:"canonical_url"`
	Readable_publish_date string `json:"readable_publish_date"`
	Cover_image string `json:"cover_image"`
	Crossposted_at string `json:"crossposted_at"`
	Edited_at string `json:"edited_at"`
	Social_image string `json:"social_image"`
	Published_at string `json:"published_at"`
	Reading_time_minutes int `json:"reading_time_minutes"` // Reading time, in minutes
	Description string `json:"description"`
	Slug string `json:"slug"`
	Tags string `json:"tags"`
	User SharedUser `json:"user"` // The resource creator
	Created_at string `json:"created_at"`
	Flare_tag ArticleFlareTag `json:"flare_tag,omitempty"` // Flare tag of the article
	Title string `json:"title"`
	Last_comment_at string `json:"last_comment_at"`
	Positive_reactions_count int `json:"positive_reactions_count"`
	Url string `json:"url"`
	Organization SharedOrganization `json:"organization,omitempty"` // The organization the resource belongs to
	Published_timestamp string `json:"published_timestamp"` // Crossposting or published date time
	Tag_list []string `json:"tag_list"`
	Type_of string `json:"type_of"`
	Path string `json:"path"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Bg_color_hex string `json:"bg_color_hex,omitempty"`
	Id int64 `json:"id,omitempty"` // Tag id
	Name string `json:"name,omitempty"`
	Text_color_hex string `json:"text_color_hex,omitempty"`
}

// Article represents the Article schema from the OpenAPI specification
type Article struct {
	Article map[string]interface{} `json:"article,omitempty"`
}

// FollowedTag represents the FollowedTag schema from the OpenAPI specification
type FollowedTag struct {
	Points float32 `json:"points"`
	Id int64 `json:"id"` // Tag id
	Name string `json:"name"`
}

// ArticleFlareTag represents the ArticleFlareTag schema from the OpenAPI specification
type ArticleFlareTag struct {
	Bg_color_hex string `json:"bg_color_hex,omitempty"` // Background color (hexadecimal)
	Name string `json:"name,omitempty"`
	Text_color_hex string `json:"text_color_hex,omitempty"` // Text color (hexadecimal)
}

// Page represents the Page schema from the OpenAPI specification
type Page struct {
	Title string `json:"title"` // Title of the page
	Body_json string `json:"body_json,omitempty"` // For JSON pages, the JSON body
	Body_markdown string `json:"body_markdown,omitempty"` // The text (in markdown) of the ad (required)
	Description string `json:"description"` // For internal use, helps similar pages from one another
	Is_top_level_path bool `json:"is_top_level_path,omitempty"` // If true, the page is available at '/{slug}' instead of '/page/{slug}', use with caution
	Slug string `json:"slug"` // Used to link to this page in URLs, must be unique and URL-safe
	Social_image map[string]interface{} `json:"social_image,omitempty"`
	Template string `json:"template"` // Controls what kind of layout the page is rendered in
}

// VideoArticle represents the VideoArticle schema from the OpenAPI specification
type VideoArticle struct {
	Title string `json:"title,omitempty"`
	Video_duration_in_minutes string `json:"video_duration_in_minutes,omitempty"`
	Type_of string `json:"type_of,omitempty"`
	User map[string]interface{} `json:"user,omitempty"` // Author of the article
	User_id int64 `json:"user_id,omitempty"`
	Video_source_url string `json:"video_source_url,omitempty"`
	Cloudinary_video_url string `json:"cloudinary_video_url,omitempty"`
	Id int64 `json:"id,omitempty"`
	Path string `json:"path,omitempty"`
}

// Organization represents the Organization schema from the OpenAPI specification
type Organization struct {
	Tag_line string `json:"tag_line,omitempty"`
	Tech_stack string `json:"tech_stack,omitempty"`
	Type_of string `json:"type_of,omitempty"`
	Url string `json:"url,omitempty"`
	Username string `json:"username,omitempty"`
	Name string `json:"name,omitempty"`
	Joined_at string `json:"joined_at,omitempty"`
	Github_username string `json:"github_username,omitempty"`
	Location string `json:"location,omitempty"`
	Story string `json:"story,omitempty"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Summary string `json:"summary,omitempty"`
}

// PodcastEpisodeIndex represents the PodcastEpisodeIndex schema from the OpenAPI specification
type PodcastEpisodeIndex struct {
	Path string `json:"path"`
	Podcast SharedPodcast `json:"podcast"` // The podcast that the resource belongs to
	Title string `json:"title"`
	Type_of string `json:"type_of"`
	Class_name string `json:"class_name"`
	Id int `json:"id"`
	Image_url string `json:"image_url"` // Podcast episode image url or podcast image url
}

// User represents the User schema from the OpenAPI specification
type User struct {
	Name string `json:"name,omitempty"`
	Summary string `json:"summary,omitempty"`
	Website_url string `json:"website_url,omitempty"`
	Joined_at string `json:"joined_at,omitempty"`
	Profile_image string `json:"profile_image,omitempty"`
	Github_username string `json:"github_username,omitempty"`
	Location string `json:"location,omitempty"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Id int64 `json:"id,omitempty"`
	Type_of string `json:"type_of,omitempty"`
	Username string `json:"username,omitempty"`
}

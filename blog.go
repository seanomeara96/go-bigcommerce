package bigcommerce

import (
	"fmt"
	"strconv"
)

type Blog struct {
	ID                   int      `json:"id"`
	Title                string   `json:"title"`
	URL                  string   `json:"url"`
	PreviewURL           string   `json:"preview_url"`
	Body                 string   `json:"body"`
	Tags                 []string `json:"tags"`
	Summary              string   `json:"summary"`
	IsPublished          bool     `json:"is_published"`
	PublishedDate        Date     `json:"published_date"`
	PublishedDateISO8601 string   `json:"published_date_iso8601"`
	MetaDescription      string   `json:"meta_description"`
	MetaKeywords         string   `json:"meta_keywords"`
	Author               string   `json:"author"`
	ThumbnailPath        string   `json:"thumbnail_path"`
}

type UpdateBlogParams struct {
	Title           string   `json:"title,omitempty"`
	URL             string   `json:"url,omitempty"`
	Body            string   `json:"body,omitempty"`
	Tags            []string `json:"tags,omitempty"`
	IsPublished     bool     `json:"is_published,omitempty"`
	MetaDescription string   `json:"meta_description,omitempty"`
	MetaKeywords    string   `json:"meta_keywords,omitempty"`
	Author          string   `json:"author,omitempty"`
	ThumbnailPath   string   `json:"thumbnail_path,omitempty"`
	PublishedDate   string   `json:"published_date,omitempty"`
}

type Date struct {
	Date         string `json:"date"`
	TimezoneType int    `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}

// GetBlog retrieves a specific blog post by its ID.
//
// Parameters:
//   - id: The unique identifier of the blog post to retrieve.
//
// Returns:
//   - Blog: The retrieved blog post.
//   - error: An error if the request fails, or nil if successful.
//
// Example usage:
//
//	blog, err := client.V2.GetBlog(123)
//	if err != nil {
//	    log.Fatalf("Failed to get blog: %v", err)
//	}
//	fmt.Printf("Blog title: %s\n", blog.Title)
func (client *V2Client) GetBlog(id int) (Blog, error) {
	type ResponseObject struct {
		Data Blog     `json:"data"`
		Meta MetaData `json:"meta"`
	}

	var response ResponseObject

	path := client.constructURL("/blog/posts", strconv.Itoa(id))

	if err := client.Get(path, &response); err != nil {
		return response.Data, fmt.Errorf("failed to get blog with ID %d: %w", id, err)
	}

	return response.Data, nil
}

// UpdateBlog updates an existing blog post with the specified ID.
//
// Parameters:
//   - blogId: The unique identifier of the blog post to update.
//   - params: UpdateBlogParams containing the fields to be updated.
//
// Returns:
//   - Blog: The updated blog post.
//   - error: An error if the request fails, or nil if successful.
//
// Example usage:
//
//	params := UpdateBlogParams{
//	    Title: "Updated Blog Title",
//	    Body:  "This is the updated blog content.",
//	}
//	updatedBlog, err := client.V2.UpdateBlog(123, params)
//	if err != nil {
//	    log.Fatalf("Failed to update blog: %v", err)
//	}
//	fmt.Printf("Updated blog title: %s\n", updatedBlog.Title)
func (client *V2Client) UpdateBlog(blogId int, params UpdateBlogParams) (Blog, error) {
	type ResponseObject struct {
		Data Blog     `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	path := client.constructURL("/blog/posts", strconv.Itoa(blogId))

	if err := client.Put(path, params, &response.Data); err != nil {
		return response.Data, fmt.Errorf("failed to update blog with ID %d: %w", blogId, err)
	}

	return response.Data, nil
}

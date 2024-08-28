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

func (client *V2Client) UpdateBlog(blogId int, params UpdateBlogParams) (Blog, error) {
	type ResponseObject struct {
		Data Blog     `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	path := client.constructURL("/blog/posts", strconv.Itoa(blogId))

	if err := client.Put(path, params, &response); err != nil {
		return response.Data, fmt.Errorf("failed to update blog with ID %d: %w", blogId, err)
	}

	return response.Data, nil
}

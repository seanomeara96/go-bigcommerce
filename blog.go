package bigcommerce

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func (client *Client) GetBlog(id int) (Blog, error) {
	return client.blogOperation(http.MethodGet, id, nil)
}

func (client *Client) UpdateBlog(blogId int, params UpdateBlogParams) (Blog, error) {
	return client.blogOperation(http.MethodPut, blogId, params)
}

func (client *Client) blogOperation(method string, id int, params interface{}) (Blog, error) {
	var blog Blog
	path := client.BaseURL().JoinPath("/blog/posts", fmt.Sprint(id)).String()

	var resp *http.Response
	var err error

	if method == http.MethodPut {
		payloadBytes, err := json.Marshal(params)
		if err != nil {
			return blog, fmt.Errorf("error marshaling params: %w", err)
		}
		resp, err = client.Put(path, payloadBytes)
	} else {
		resp, err = client.Get(path)
	}

	if err != nil {
		return blog, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if err = expectStatusCode(200, resp); err != nil {
		return blog, err
	}

	if err = json.NewDecoder(resp.Body).Decode(&blog); err != nil {
		return blog, fmt.Errorf("error decoding response: %w", err)
	}

	return blog, nil
}

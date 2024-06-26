package bigcommerce

import (
	"encoding/json"
	"fmt"
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
	type ResponseObject struct {
		Data Blog     `json:"data"`
		Meta MetaData `json:"meta"`
	}

	var response ResponseObject

	path := client.BaseURL().JoinPath("/blog/posts", fmt.Sprint(id)).String()

	resp, err := client.Get(path)
	if err != nil {
		return response.Data, nil
	}
	defer resp.Body.Close()

	if err = expectStatusCode(200, resp); err != nil {
		return response.Data, err
	}

	// passing in response.Data here instead as it is V2
	if err = json.NewDecoder(resp.Body).Decode(&response.Data); err != nil {
		return response.Data, err
	}

	return response.Data, nil
}

func (client *Client) UpdateBlog(blogId int, params UpdateBlogParams) (Blog, error) {
	type ResponseObject struct {
		Data Blog     `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	path := client.BaseURL().JoinPath("/blog/posts", fmt.Sprint(blogId)).String()

	payloadBytes, err := json.Marshal(params)
	if err != nil {
		return response.Data, err
	}

	resp, err := client.Put(path, payloadBytes)
	if err != nil {
		return response.Data, err
	}
	defer resp.Body.Close()

	if err = expectStatusCode(200, resp); err != nil {
		return response.Data, err
	}

	// passing in response.Data here instead as it is V2
	if err = json.NewDecoder(resp.Body).Decode(&response.Data); err != nil {
		return response.Data, err
	}

	return response.Data, nil
}

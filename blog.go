package bigcommerce

import (
	"encoding/json"
	"fmt"
	"time"
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
	PublishedDate        string   `json:"published_date"`
	PublishedDateISO8601 string   `json:"published_date_iso8601"`
	MetaDescription      string   `json:"meta_description"`
	MetaKeywords         string   `json:"meta_keywords"`
	Author               string   `json:"author"`
	ThumbnailPath        string   `json:"thumbnail_path"`
}

func (ct *Blog) ParsePublishDate() (time.Time, error) {
	parsedTime, err := time.Parse(`"2006-01-02 15:04:05.000000"`, string(ct.PublishedDate))
	if err != nil {
		return nil, err
	}
	return parsedTime, nil
}

func (client *Client) GetBrand(id int) (Blog, error) {
	type ResponseObject struct {
		Data Blog     `json:"data"`
		Meta MetaData `json:"meta"`
	}

	var response ResponseObject

	path := client.BaseURL.JoinPath("/blog/posts", fmt.Sprint(id)).String()

	resp, err := client.Get(path)
	if err != nil {
		return response.Data, nil
	}
	defer resp.Body.Close()

	if err = expectStatusCode(200, resp); err != nil {
		return response.Data, err
	}

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return response.Data, err
	}

	return response.Data, nil
}

package bigcommerce

import (
	"errors"
)

type Redirect struct {
	ID       int              `json:"id"`
	SiteID   int              `json:"site_id"`
	FromPath string           `json:"from_path"`
	To       RedirectToObject `json:"to"`
	ToURL    string           `json:"to_url"`
}

func FromPaths(redirects []Redirect) []string {
	fromPaths := []string{}
	for i := range redirects {
		fromPaths = append(fromPaths, redirects[i].FromPath)
	}
	return fromPaths
}

type RedirectToObject struct {
	Type     string `json:"type"`
	EntityID int    `json:"entity_id"`
	URL      string `json:"url"`
}

func (client *V3Client) GetAllRedirects(params RedirectQueryParams) ([]Redirect, error) {
	redirects := []Redirect{}
	params.Page = 1
	params.Limit = 250
	for {
		res, err := client.GetRedirects(params)
		if err != nil {
			return nil, err
		}
		if len(res) < params.Limit {
			return redirects, nil
		}

		redirects = append(redirects, res...)

		params.Page++
	}
}

func (client *V3Client) GetRedirects(params RedirectQueryParams) ([]Redirect, error) {
	type ResponseObject struct {
		Data []Redirect `json:"data"`
		Meta MetaData   `json:"meta"`
	}
	var response ResponseObject

	getRedirectsURL, err := urlWithQueryParams(client.constructURL("/storefront/redirects"), params)
	if err != nil {
		return response.Data, err
	}

	if err := client.Get(getRedirectsURL, &response); err != nil {
		return response.Data, err
	}

	return response.Data, nil
}

type RedirectQueryParams struct {
	SiteID    int    `url:"site_id,omitempty"`
	IDs       []int  `url:"id,omitempty"`
	Limit     int    `url:"limit,omitempty"`
	Page      int    `url:"page,omitempty"`
	Sort      string `url:"sort,omitempty"`
	Direction string `url:"direction,omitempty"`
	Include   string `url:"include,omitempty"`
	Keyword   string `url:"keyword,omitempty"`
}

func validateRedirectUpsert(redirect RedirectUpsert) error {
	if redirect.FromPath == "" {
		return errors.New("from_path is required")
	}

	if redirect.SiteID <= 0 {
		return errors.New("site_id must be a positive integer")
	}

	if redirect.To.Type == "" {
		return errors.New("to.type is required")
	}

	if redirect.To.Type != "product" && redirect.To.Type != "brand" && redirect.To.Type != "category" &&
		redirect.To.Type != "page" && redirect.To.Type != "post" && redirect.To.Type != "url" {
		return errors.New("to.type has an invalid value")
	}

	if redirect.To.Type != "url" && redirect.To.EntityID <= 0 {
		return errors.New("to.entity_id must be a positive integer")
	}

	if redirect.To.Type == "url" && len(redirect.To.URL) > 2048 {
		return errors.New("to.url must be 2048 characters or less")
	}

	return nil
}

type RedirectUpsert struct {
	FromPath string         `json:"from_path"`
	SiteID   int            `json:"site_id"`
	To       RedirectTarget `json:"to"`
}

type RedirectTarget struct {
	Type     string `json:"type"`
	EntityID int    `json:"entity_id"`
	URL      string `json:"url"`
}

func (client *V3Client) UpsertRedirects(redirects []RedirectUpsert) ([]Redirect, error) {
	type ResponseObject struct {
		Data []Redirect `json:"data"`
		Meta MetaData   `json:"meta"`
	}
	var response ResponseObject

	for i := 0; i < len(redirects); i++ {
		err := validateRedirectUpsert(redirects[i])
		if err != nil {
			return response.Data, err
		}
	}

	path := client.constructURL("/storefront/redirects")

	err := client.Put(path, redirects, &response)
	if err != nil {
		return response.Data, err
	}

	return response.Data, nil
}

type DeleteRedirectsParams struct {
	ID     []int `url:"id:in,omitempty"`
	SiteID int   `url:"site_id,omitempty"`
}

func (client *V3Client) DeleteRedirect(params DeleteRedirectsParams) error {
	path, err := urlWithQueryParams(client.constructURL("/storefront/redirects"), params)
	if err != nil {
		return err
	}

	if err := client.Delete(path, nil); err != nil {
		return err
	}

	return nil
}

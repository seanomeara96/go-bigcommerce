package bigcommerce

import (
	"encoding/json"
	"fmt"
)

type Page struct {
	ID              int    `json:"id"`
	ChannelID       int    `json:"channel_id"`
	Name            string `json:"name" validate:"required,min=1,max=100"`
	IsVisible       bool   `json:"is_visible"`
	ParentID        int    `json:"parent_id"`
	SortOrder       int    `json:"sort_order"`
	Type            string `json:"type" validate:"required,oneof=page raw contact_form feed link blog"`
	IsHomepage      bool   `json:"is_homepage"`
	IsCustomersOnly bool   `json:"is_customers_only"`
	URL             string `json:"url"`
	MetaTitle       string `json:"meta_title"`
	MetaKeywords    string `json:"meta_keywords"`
	MetaDescription string `json:"meta_description"`
	SearchKeywords  string `json:"search_keywords"`
}

type GetPagesParams struct {
	ChannelID int    `url:"channel_id,omitempty"`
	ID        string `url:"id,in,omitempty"`
	Name      string `url:"name,omitempty"`
	NameLike  string `url:"name:like,omitempty"`
	Limit     int    `url:"limit,omitempty"`
	Page      int    `url:"page,omitempty"`
	Include   string `url:"include,omitempty"`
}

func (client *Client) GetPages(queryParams GetPagesParams) ([]Page, MetaData, error) {
	type ResponseObject struct {
		Data []Page   `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	queryString, err := paramString(queryParams)
	if err != nil {
		return response.Data, response.Meta, err
	}

	path := client.BaseURL.JoinPath("/content/pages").String() + queryString

	resp, err := client.Get(path)
	if err != nil {
		return response.Data, response.Meta, err
	}
	defer resp.Body.Close()

	err = expectStatusCode(200, resp)
	if err != nil {
		return response.Data, response.Meta, err
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response.Data, response.Meta, err
	}

	return response.Data, response.Meta, nil
}

func (client *Client) CreatePage(params CreatePageParams) (Page, error) {
	type ResponseObject struct {
		Data Page     `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	path := client.BaseURL.JoinPath("/content/pages").String()

	paramBytes, err := json.Marshal(params)
	if err != nil {
		return response.Data, err
	}

	resp, err := client.Post(path, paramBytes)
	if err != nil {
		return response.Data, err
	}
	defer resp.Body.Close()

	err = expectStatusCode(201, resp)
	if err != nil {
		err = expectStatusCode(207, resp)
		if err != nil {
			return response.Data, err
		}
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response.Data, err
	}

	return response.Data, nil

}

func (client *Client) DeletePage(pageID int) error {
	path := client.BaseURL.JoinPath("/content/pages", fmt.Sprint(pageID)).String()
	resp, err := client.Delete(path)
	if err != nil {
		return err
	}
	err = expectStatusCode(204, resp)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) GetPage(pageID int) (Page, error) {
	type ResponseObject struct {
		Data Page     `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	path := client.BaseURL.JoinPath("/content/pages", fmt.Sprint(pageID)).String()

	resp, err := client.Get(path)
	if err != nil {
		return response.Data, err
	}
	defer resp.Body.Close()

	err = expectStatusCode(200, resp)
	if err != nil {
		return response.Data, err
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response.Data, err
	}

	return response.Data, nil
}

func (client *Client) UpdatePage(pageID int, params UpdatePageParams) (Page, error) {
	type ResponseObject struct {
		Data Page     `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	path := client.BaseURL.JoinPath("/content/pages", fmt.Sprint(pageID)).String()

	paramBytes, err := json.Marshal(params)
	if err != nil {
		return response.Data, err
	}

	resp, err := client.Put(path, paramBytes)
	if err != nil {
		return response.Data, err
	}
	defer resp.Body.Close()

	err = expectStatusCode(200, resp)
	if err != nil {
		return response.Data, err
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response.Data, err
	}

	return response.Data, nil
}

type CreatePageParams struct {
	Email           string         `json:"email,omitempty" validate:"omitempty,max=255"`
	MetaTitle       string         `json:"meta_title,omitempty"`
	Body            string         `json:"body,omitempty"`
	Feed            string         `json:"feed,omitempty"`
	Link            string         `json:"link,omitempty"`
	ContactFields   []ContactField `json:"contact_fields,omitempty"`
	MetaKeywords    string         `json:"meta_keywords,omitempty"`
	MetaDescription string         `json:"meta_description,omitempty"`
	SearchKeywords  string         `json:"search_keywords,omitempty"`
	URL             string         `json:"url,omitempty"`
	ChannelID       int            `json:"channel_id,omitempty"`
	Name            string         `json:"name" validate:"required,min=1,max=100"`
	IsVisible       bool           `json:"is_visible,omitempty"`
	ParentID        int            `json:"parent_id,omitempty"`
	SortOrder       int            `json:"sort_order,omitempty"`
	Type            PageType       `json:"type" validate:"required,oneof=page raw contact_form feed link blog"`
	IsHomepage      bool           `json:"is_homepage,omitempty"`
	IsCustomersOnly bool           `json:"is_customers_only,omitempty"`
}

type UpdatePageParams struct {
	Name            string         `json:"name,omitempty"`
	IsVisible       bool           `json:"is_visible,omitempty"`
	ParentID        int            `json:"parent_id,omitempty"`
	SortOrder       int            `json:"sort_order,omitempty"`
	Type            PageType       `json:"type,omitempty"`
	IsHomepage      bool           `json:"is_homepage,omitempty"`
	IsCustomersOnly bool           `json:"is_customers_only,omitempty"`
	ID              int            `json:"id,omitempty"`
	Email           string         `json:"email,omitempty"`
	MetaTitle       string         `json:"meta_title,omitempty"`
	Body            string         `json:"body,omitempty"`
	Feed            string         `json:"feed,omitempty"`
	Link            string         `json:"link,omitempty"`
	ContactFields   []ContactField `json:"contact_fields,omitempty"`
	MetaKeywords    string         `json:"meta_keywords,omitempty"`
	MetaDescription string         `json:"meta_description,omitempty"`
	SearchKeywords  string         `json:"search_keywords,omitempty"`
	URL             string         `json:"url,omitempty"`
	ChannelID       int            `json:"channel_id,omitempty"`
}

type PageType string

const (
	BlogPage        PageType = "blog"
	ContactFormPage PageType = "contact_form"
	LinkPage        PageType = "link"
	UserDefinedPage PageType = "page"
	RawPage         PageType = "raw"
	RSSFeedPage     PageType = "rss_feed"
)

// AllowedPageTypes is a slice containing all allowed page types.
var AllowedPageTypes = []PageType{
	BlogPage,
	ContactFormPage,
	LinkPage,
	UserDefinedPage,
	RawPage,
	RSSFeedPage,
}

type ContactField string

const (
	FullnameField    ContactField = "fullname"
	PhoneField       ContactField = "phone"
	CompanyNameField ContactField = "companyname"
	OrderNoField     ContactField = "orderno"
	RMAField         ContactField = "rma"
)

// AllowedContactFields is a slice containing all allowed contact fields.
var AllowedContactFields = []ContactField{
	FullnameField,
	PhoneField,
	CompanyNameField,
	OrderNoField,
	RMAField,
}

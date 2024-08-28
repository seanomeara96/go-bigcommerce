package bigcommerce

import (
	"fmt"
	"strconv"
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

func (client *V3Client) GetPages(queryParams GetPagesParams) ([]Page, MetaData, error) {
	type ResponseObject struct {
		Data []Page   `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	path, err := urlWithQueryParams(client.constructURL("/content/pages"), queryParams)
	if err != nil {
		return nil, MetaData{}, fmt.Errorf("failed to construct URL for GetPages: %w", err)
	}

	if err := client.Get(path, &response); err != nil {
		return nil, MetaData{}, fmt.Errorf("failed to get pages: %w", err)
	}

	return response.Data, response.Meta, nil
}

func (client *V3Client) CreatePage(params CreatePageParams) (Page, error) {
	type ResponseObject struct {
		Data Page     `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	path := client.constructURL("/content/pages")

	if err := client.Post(path, params, &response); err != nil {
		return Page{}, fmt.Errorf("failed to create page: %w", err)
	}

	return response.Data, nil
}

func (client *V3Client) DeletePage(pageID int) error {
	path := client.constructURL("/content/pages", strconv.Itoa(pageID))

	if err := client.Delete(path, nil); err != nil {
		return fmt.Errorf("failed to delete page with ID %d: %w", pageID, err)
	}

	return nil
}

func (client *V3Client) GetPage(pageID int) (Page, error) {
	type ResponseObject struct {
		Data Page     `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	path := client.constructURL("/content/pages", strconv.Itoa(pageID))

	if err := client.Get(path, &response); err != nil {
		return Page{}, fmt.Errorf("failed to get page with ID %d: %w", pageID, err)
	}

	return response.Data, nil
}

func (client *V3Client) UpdatePage(pageID int, params UpdatePageParams) (Page, error) {
	type ResponseObject struct {
		Data Page     `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	path := client.constructURL("/content/pages", strconv.Itoa(pageID))

	if err := client.Put(path, params, &response); err != nil {
		return Page{}, fmt.Errorf("failed to update page with ID %d: %w", pageID, err)
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

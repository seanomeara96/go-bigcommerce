package bigcommerce

import (
	"fmt"
	"strconv"
)

type Banner struct {
	ID          int    `json:"id"`
	DateCreated string `json:"date_created"`
	Name        string `json:"name"`
	Content     string `json:"content"`
	Page        string `json:"page"`
	Location    string `json:"location"`
	DateType    string `json:"date_type"`
	DateFrom    string `json:"date_from,omitempty"`
	DateTo      string `json:"date_to,omitempty"`
	Visible     string `json:"visible"`
	ItemID      string `json:"item_id,omitempty"`
}
type GetBannersParams struct {
	MinID int `url:"min_id,omitempty"`
	MaxID int `url:"max_id,omitempty"`
	Page  int `url:"page,omitempty"`
	Limit int `url:"limit,omitempty"`
}

type CreateUpdateBannerParams struct {
	Name     string `json:"name" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Page     string `json:"page" binding:"required"`
	Location string `json:"location" binding:"required"`
	DateType string `json:"date_type" binding:"required"`
	DateFrom string `json:"date_from,omitempty"`
	DateTo   string `json:"date_to,omitempty"`
	Visible  string `json:"visible,omitempty"`
	ItemID   string `json:"item_id,omitempty"`
}

type ValidationErrors []string

func (ve ValidationErrors) Error() string {
	return fmt.Sprintf("validation failed: %v", []string(ve))
}

func validateBannerParams(params CreateUpdateBannerParams) error {
	var errors ValidationErrors

	if params.Name == "" {
		errors = append(errors, "Name is required")
	}

	if params.Content == "" {
		errors = append(errors, "Content is required")
	}

	if params.Page == "" {
		errors = append(errors, "Page is required")
	}

	if params.Location == "" {
		errors = append(errors, "Location is required")
	}

	if params.DateType == "" {
		errors = append(errors, "DateType is required")
	}

	if params.DateType == "custom" {
		if params.DateFrom == "" {
			errors = append(errors, "DateFrom is required when DateType is 'custom'")
		}
		if params.DateTo == "" {
			errors = append(errors, "DateTo is required when DateType is 'custom'")
		}
	}

	if params.Visible == "" {
		errors = append(errors, "Visible is required")
	}

	if params.ItemID == "" && (params.Page == "category_page" || params.Page == "brand_page") {
		errors = append(errors, "ItemID is required for category_page or brand_page")
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (client *Client) CreateBanner(params CreateUpdateBannerParams) (Banner, error) {
	type ResponseObject struct {
		Data Banner   `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject
	err := client.Version2Required()
	if err != nil {
		return response.Data, fmt.Errorf("CreateBanner: version 2 API required: %w", err)
	}

	err = validateBannerParams(params)
	if err != nil {
		return response.Data, fmt.Errorf("CreateBanner: invalid parameters: %w", err)
	}

	path := client.constructURL("banners")

	err = client.Post(path, params, &response)
	if err != nil {
		return response.Data, fmt.Errorf("CreateBanner: failed to create banner: %w", err)
	}

	return response.Data, nil
}

func (client *Client) UpdateBanner(bannerID int, params CreateUpdateBannerParams) (Banner, error) {
	type ResponseObject struct {
		Data Banner   `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	if err := client.Version2Required(); err != nil {
		return response.Data, fmt.Errorf("UpdateClient: version 2 API required: %w", err)
	}

	if err := validateBannerParams(params); err != nil {
		return response.Data, fmt.Errorf("UpdateClient: invalid parameters: %w", err)
	}

	path := client.constructURL("banners", strconv.Itoa(bannerID))

	if err := client.Put(path, params, &response); err != nil {
		return response.Data, fmt.Errorf("UpdateClient: failed to update banner %d: %w", bannerID, err)
	}

	return response.Data, nil
}

func (client *Client) GetBanners(params GetBannersParams) ([]Banner, MetaData, error) {
	type ResponseObject struct {
		Data []Banner `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	err := client.Version2Required()
	if err != nil {
		return response.Data, response.Meta, fmt.Errorf("GetBanners: version 2 API required: %w", err)
	}

	path, err := urlWithQueryParams(client.constructURL("banners"), params)
	if err != nil {
		return response.Data, response.Meta, fmt.Errorf("GetBanners: failed to construct URL with query params: %w", err)
	}

	if err := client.Get(path, &response); err != nil {
		return response.Data, response.Meta, fmt.Errorf("GetBanners: failed to retrieve banners: %w", err)
	}

	return response.Data, response.Meta, nil
}

func (client *Client) GetBanner(bannerID int) (Banner, error) {
	type ResponseObject struct {
		Data Banner   `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	err := client.Version2Required()
	if err != nil {
		return response.Data, fmt.Errorf("GetBanner: version 2 API required: %w", err)
	}

	path := client.constructURL("banners", strconv.Itoa(bannerID))

	if err := client.Get(path, &response); err != nil {
		return response.Data, fmt.Errorf("GetBanner: failed to retrieve banner %d: %w", bannerID, err)
	}

	return response.Data, nil
}

func (client *Client) DeleteBanner(bannerID int) error {
	err := client.Version2Required()
	if err != nil {
		return fmt.Errorf("DeleteBanner: version 2 API required: %w", err)
	}
	path := client.constructURL("banners", strconv.Itoa(bannerID))
	if err := client.Delete(path, nil); err != nil {
		return fmt.Errorf("DeleteBanner: failed to delete banner %d: %w", bannerID, err)
	}

	return nil
}

package bigcommerce

import (
	"strconv"
)

type ProductVideo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	SortOrder   int    `json:"sort_order"`
	Type        string `json:"type"`
	VideoID     string `json:"video_id"`
	ID          int    `json:"id"`
	ProductID   int    `json:"product_id"`
	Length      string `json:"length"`
}

func (client *V3Client) GetAllProductVideos(productID int, params GetAllProductVideosQueryParams) ([]ProductVideo, MetaData, error) {
	type ResponseObject struct {
		Data []ProductVideo `json:"data"`
		Meta MetaData       `json:"meta"`
	}
	var response ResponseObject

	getProductVideosPath, err := urlWithQueryParams(client.constructURL("catalog", "products", strconv.Itoa(productID), "videos"), params)
	if err != nil {
		return response.Data, response.Meta, err
	}

	if err := client.Get(getProductVideosPath, &response); err != nil {
		return response.Data, response.Meta, err
	}

	return response.Data, response.Meta, nil
}

/*func (client *Client) CreateProductVideo() {}
func (client *Client) GetProductVideo()    {}
func (client *Client) UpdateProductVideo() {}
func (client *Client) DeleteProductVideo() {}*/

type GetAllProductVideosQueryParams struct {
	IncludeFields string `url:"include_fields,omitempty"`
	ExcludeFields string `url:"exclude_fields,omitempty"`
	Page          int    `url:"page,omitempty"`
	Limit         int    `url:"limit,omitempty"`
}

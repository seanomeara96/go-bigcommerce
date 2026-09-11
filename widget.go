package bigcommerce

import "net/url"

// Widget is a configured instance of a WidgetTemplate.
type Widget struct {
	UUID                     string         `json:"uuid"`
	Name                     string         `json:"name"`
	Description              string         `json:"description"`
	WidgetTemplateUUID       string         `json:"widget_template_uuid"`
	VersionUUID              string         `json:"version_uuid"`
	CurrentVersionUUID       string         `json:"current_version_uuid"`
	WidgetConfiguration      map[string]any `json:"widget_configuration"`
	StorefrontAPIQueryParams map[string]any `json:"storefront_api_query_params"`
	WidgetTemplate           WidgetTemplate `json:"widget_template"`
	ChannelID                int            `json:"channel_id"`
	DateCreated              string         `json:"date_created"`
	DateModified             string         `json:"date_modified"`
}

type WidgetQueryParams struct {
	Page               int      `url:"page,omitempty"`
	Limit              int      `url:"limit,omitempty"`
	WidgetTemplateKind string   `url:"widget_template_kind,omitempty"`
	WidgetTemplateUUID string   `url:"widget_template_uuid,omitempty"`
	Name               string   `url:"name,omitempty"`
	Names              []string `url:"name:in,omitempty,comma"`
	ChannelIDs         []int    `url:"channel_id:in,omitempty,comma"`
	SiteIDs            []int    `url:"site_id:in,omitempty,comma"`
}

type CreateWidgetParams struct {
	Name                string         `json:"name"`
	WidgetTemplateUUID  string         `json:"widget_template_uuid"`
	Description         string         `json:"description,omitempty"`
	WidgetConfiguration map[string]any `json:"widget_configuration,omitempty"`
	ChannelID           int            `json:"channel_id,omitempty"`
}

type UpdateWidgetParams struct {
	Name                *string         `json:"name,omitempty"`
	WidgetTemplateUUID  *string         `json:"widget_template_uuid,omitempty"`
	Description         *string         `json:"description,omitempty"`
	WidgetConfiguration *map[string]any `json:"widget_configuration,omitempty"`
	ChannelID           *int            `json:"channel_id,omitempty"`
	Upgrade             *bool           `json:"upgrade,omitempty"`
}

func (client *V3Client) GetWidgets(params WidgetQueryParams) ([]Widget, MetaData, error) {
	var response struct {
		Data []Widget `json:"data"`
		Meta MetaData `json:"meta"`
	}

	path, err := widgetURLWithQuery(client.constructURL("content", "widgets"), params)
	if err != nil {
		return nil, MetaData{}, err
	}
	if err := client.Get(path, &response); err != nil {
		return response.Data, response.Meta, err
	}
	return response.Data, response.Meta, nil
}

func (client *V3Client) GetWidget(uuid string) (Widget, error) {
	return client.getWidgetAt(client.constructURL("content", "widgets", uuid))
}

func (client *V3Client) CreateWidget(params CreateWidgetParams) (Widget, error) {
	var response struct {
		Data Widget `json:"data"`
	}
	err := client.Post(client.constructURL("content", "widgets"), params, &response)
	return response.Data, err
}

func (client *V3Client) UpdateWidget(uuid string, params UpdateWidgetParams) (Widget, error) {
	var response struct {
		Data Widget `json:"data"`
	}
	err := client.Put(client.constructURL("content", "widgets", uuid), params, &response)
	return response.Data, err
}

func (client *V3Client) DeleteWidget(uuid string) error {
	return client.Delete(client.constructURL("content", "widgets", uuid), nil)
}

func (client *V3Client) getWidgetAt(path *url.URL) (Widget, error) {
	var response struct {
		Data Widget `json:"data"`
	}
	err := client.Get(path, &response)
	return response.Data, err
}

func widgetURLWithQuery(path *url.URL, params any) (*url.URL, error) {
	return urlWithQueryParams(path, params)
}

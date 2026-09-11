package bigcommerce

// WidgetTemplate defines the Handlebars markup and optional Page Builder UI
// schema shared by one or more widgets.
type WidgetTemplate struct {
	UUID               string           `json:"uuid"`
	CurrentVersionUUID string           `json:"current_version_uuid"`
	Name               string           `json:"name"`
	Template           string           `json:"template"`
	Schema             []map[string]any `json:"schema"`
	StorefrontAPIQuery string           `json:"storefront_api_query"`
	Kind               string           `json:"kind"`
	IconName           string           `json:"icon_name"`
	TemplateEngine     string           `json:"template_engine"`
	ClientRerender     bool             `json:"client_rerender"`
	ChannelID          int              `json:"channel_id"`
	DateCreated        string           `json:"date_created"`
	DateModified       string           `json:"date_modified"`
}

type WidgetTemplateQueryParams struct {
	Page               int    `url:"page,omitempty"`
	Limit              int    `url:"limit,omitempty"`
	WidgetTemplateKind string `url:"widget_template_kind,omitempty"`
	Name               string `url:"name,omitempty"`
	ChannelID          int    `url:"channel_id,omitempty"`
}

type CreateWidgetTemplateParams struct {
	Name               string           `json:"name"`
	Template           string           `json:"template"`
	Schema             []map[string]any `json:"schema,omitempty"`
	StorefrontAPIQuery string           `json:"storefront_api_query,omitempty"`
	ChannelID          int              `json:"channel_id,omitempty"`
}

type UpdateWidgetTemplateParams struct {
	Name               *string           `json:"name,omitempty"`
	Template           *string           `json:"template,omitempty"`
	Schema             *[]map[string]any `json:"schema,omitempty"`
	StorefrontAPIQuery *string           `json:"storefront_api_query,omitempty"`
	CreateNewVersion   *bool             `json:"create_new_version,omitempty"`
}

type RenderWidgetTemplateParams struct {
	WidgetConfiguration      map[string]any `json:"widget_configuration,omitempty"`
	StorefrontAPIQueryParams map[string]any `json:"storefront_api_query_params,omitempty"`
}

type RenderedWidgetTemplate struct {
	HTML string `json:"html"`
}

func (client *V3Client) GetWidgetTemplates(params WidgetTemplateQueryParams) ([]WidgetTemplate, MetaData, error) {
	var response struct {
		Data []WidgetTemplate `json:"data"`
		Meta MetaData         `json:"meta"`
	}
	path, err := widgetURLWithQuery(client.constructURL("content", "widget-templates"), params)
	if err != nil {
		return nil, MetaData{}, err
	}
	if err := client.Get(path, &response); err != nil {
		return response.Data, response.Meta, err
	}
	return response.Data, response.Meta, nil
}

func (client *V3Client) GetWidgetTemplate(uuid string) (WidgetTemplate, error) {
	var response struct {
		Data WidgetTemplate `json:"data"`
	}
	err := client.Get(client.constructURL("content", "widget-templates", uuid), &response)
	return response.Data, err
}

func (client *V3Client) CreateWidgetTemplate(params CreateWidgetTemplateParams) (WidgetTemplate, error) {
	var response struct {
		Data WidgetTemplate `json:"data"`
	}
	err := client.Post(client.constructURL("content", "widget-templates"), params, &response)
	return response.Data, err
}

func (client *V3Client) UpdateWidgetTemplate(uuid string, params UpdateWidgetTemplateParams) (WidgetTemplate, error) {
	var response struct {
		Data WidgetTemplate `json:"data"`
	}
	err := client.Put(client.constructURL("content", "widget-templates", uuid), params, &response)
	return response.Data, err
}

func (client *V3Client) RenderWidgetTemplate(uuid string, params RenderWidgetTemplateParams) (RenderedWidgetTemplate, error) {
	var response struct {
		Data RenderedWidgetTemplate `json:"data"`
	}
	err := client.Post(client.constructURL("content", "widget-templates", uuid, "render"), params, &response)
	return response.Data, err
}

func (client *V3Client) DeleteWidgetTemplate(uuid string) error {
	return client.Delete(client.constructURL("content", "widget-templates", uuid), nil)
}

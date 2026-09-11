package bigcommerce

type Placement struct {
	UUID         string `json:"uuid"`
	WidgetUUID   string `json:"widget_uuid"`
	Widget       Widget `json:"widget"`
	TemplateFile string `json:"template_file"`
	Region       string `json:"region"`
	EntityID     string `json:"entity_id"`
	Status       string `json:"status"`
	SortOrder    int    `json:"sort_order"`
	ChannelID    int    `json:"channel_id"`
	DateCreated  string `json:"date_created"`
	DateModified string `json:"date_modified"`
}

type PlacementQueryParams struct {
	Page               int    `url:"page,omitempty"`
	Limit              int    `url:"limit,omitempty"`
	WidgetTemplateKind string `url:"widget_template_kind,omitempty"`
	TemplateFile       string `url:"template_file,omitempty"`
	WidgetUUID         string `url:"widget_uuid,omitempty"`
	WidgetTemplateUUID string `url:"widget_template_uuid,omitempty"`
	ChannelIDs         []int  `url:"channel_id:in,omitempty,comma"`
	SiteIDs            []int  `url:"site_id:in,omitempty,comma"`
}

type CreatePlacementParams struct {
	WidgetUUID   string `json:"widget_uuid"`
	TemplateFile string `json:"template_file"`
	Status       string `json:"status"`
	Region       string `json:"region,omitempty"`
	EntityID     string `json:"entity_id,omitempty"`
	SortOrder    int    `json:"sort_order,omitempty"`
}

type UpdatePlacementParams struct {
	WidgetUUID   *string `json:"widget_uuid,omitempty"`
	TemplateFile *string `json:"template_file,omitempty"`
	Status       *string `json:"status,omitempty"`
	Region       *string `json:"region,omitempty"`
	EntityID     *string `json:"entity_id,omitempty"`
	SortOrder    *int    `json:"sort_order,omitempty"`
}

type ThemeRegion struct {
	Name string `json:"name"`
}

type ThemeRegionQueryParams struct {
	TemplateFile string `url:"template_file"`
	ChannelID    string `url:"channel_id,omitempty"`
}

func (client *V3Client) GetPlacements(params PlacementQueryParams) ([]Placement, MetaData, error) {
	var response struct {
		Data []Placement `json:"data"`
		Meta MetaData    `json:"meta"`
	}
	path, err := widgetURLWithQuery(client.constructURL("content", "placements"), params)
	if err != nil {
		return nil, MetaData{}, err
	}
	if err := client.Get(path, &response); err != nil {
		return response.Data, response.Meta, err
	}
	return response.Data, response.Meta, nil
}

func (client *V3Client) GetPlacement(uuid string) (Placement, error) {
	var response struct {
		Data Placement `json:"data"`
	}
	err := client.Get(client.constructURL("content", "placements", uuid), &response)
	return response.Data, err
}

func (client *V3Client) CreatePlacement(params CreatePlacementParams) (Placement, error) {
	var response struct {
		Data Placement `json:"data"`
	}
	err := client.Post(client.constructURL("content", "placements"), params, &response)
	return response.Data, err
}

func (client *V3Client) UpdatePlacement(uuid string, params UpdatePlacementParams) (Placement, error) {
	var response struct {
		Data Placement `json:"data"`
	}
	err := client.Put(client.constructURL("content", "placements", uuid), params, &response)
	return response.Data, err
}

func (client *V3Client) DeletePlacement(uuid string) error {
	return client.Delete(client.constructURL("content", "placements", uuid), nil)
}

func (client *V3Client) GetThemeRegions(params ThemeRegionQueryParams) ([]ThemeRegion, error) {
	var response struct {
		Data []ThemeRegion `json:"data"`
	}
	path, err := widgetURLWithQuery(client.constructURL("content", "regions"), params)
	if err != nil {
		return nil, err
	}
	err = client.Get(path, &response)
	return response.Data, err
}

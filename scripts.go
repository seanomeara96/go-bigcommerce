package bigcommerce

type Script struct {
	Name            string `json:"name"`
	UUID            string `json:"uuid"`
	DateCreated     string `json:"date_created"`
	DateModified    string `json:"date_modified"`
	Description     string `json:"description"`
	HTML            string `json:"html"`
	Src             string `json:"src"`
	AutoUninstall   bool   `json:"auto_uninstall"`
	LoadMethod      string `json:"load_method"`
	Location        string `json:"location"`
	Visibility      string `json:"visibility"`
	Kind            string `json:"kind"`
	APIClientID     string `json:"api_client_id"`
	ConsentCategory string `json:"consent_category"`
	Enabled         bool   `json:"enabled"`
	ChannelID       int    `json:"channel_id"`
}
type UpdateScriptParams struct {
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	HTML            string `json:"html,omitempty"`
	Src             string `json:"src,omitempty"`
	AutoUninstall   bool   `json:"auto_uninstall,omitempty"`
	LoadMethod      string `json:"load_method,omitempty"`
	Location        string `json:"location,omitempty"`
	Visibility      string `json:"visibility,omitempty"`
	Kind            string `json:"kind,omitempty"`
	APIClientID     string `json:"api_client_id,omitempty"`
	ConsentCategory string `json:"consent_category,omitempty"`
	Enabled         bool   `json:"enabled,omitempty"`
	ChannelID       int    `json:"channel_id,omitempty"`
}

type ScriptsQuery struct {
	Page       int      `json:"page,omitempty"`
	Limit      int      `json:"limit,omitempty"`
	Sort       string   `json:"sort,omitempty"`
	Direction  string   `json:"direction,omitempty"`
	ChannelIDs []string `json:"channel_id,omitempty"`
}

func (client *V3Client) GetScripts(params ScriptsQuery) ([]Script, MetaData, error) {
	type ResponseObject struct {
		Data []Script `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	path := client.constructURL("/content/scripts")

	err := client.Get(path, &response)
	if err != nil {
		return response.Data, response.Meta, err
	}

	return response.Data, response.Meta, nil
}

func (client *V3Client) GetAllScripts(limit int) ([]Script, error) {
	var scripts []Script
	page := 1

	for {
		p, _, err := client.GetScripts(ScriptsQuery{Limit: limit, Page: page})
		if err != nil {
			return scripts, err
		}

		for i := 0; i < len(p); i++ {
			scripts = append(scripts, p[i])
		}

		if len(p) < limit {
			break
		}

		page++
	}

	return scripts, nil
}

type CreateScriptParams struct {
	Name            string `json:"name" validate:"required,min=1,max=255"`
	Description     string `json:"description,omitempty"`
	HTML            string `json:"html,omitempty" validate:"max=65536"`
	Src             string `json:"src,omitempty"`
	AutoUninstall   bool   `json:"auto_uninstall,omitempty"`
	LoadMethod      string `json:"load_method,omitempty" validate:"omitempty,oneof=default async defer"`
	Location        string `json:"location,omitempty" validate:"omitempty,oneof=head footer"`
	Visibility      string `json:"visibility,omitempty" validate:"omitempty,oneof=storefront all_pages checkout order_confirmation"`
	Kind            string `json:"kind,omitempty" validate:"omitempty,oneof=src script_tag"`
	APIClientID     string `json:"api_client_id,omitempty"`
	ConsentCategory string `json:"consent_category,omitempty" validate:"omitempty,oneof=essential functional analytics targeting"`
	Enabled         bool   `json:"enabled,omitempty"`
	ChannelID       int    `json:"channel_id,omitempty"`
}

func StorefrontFooterHTMLScript(Name string, HTML string) CreateScriptParams {
	p := CreateScriptParams{}

	p.Name = Name
	p.HTML = HTML
	p.Kind = "script_tag"
	p.ConsentCategory = "essential"
	p.Enabled = true
	p.Visibility = "storefront"
	p.Location = "footer"
	p.LoadMethod = "default"

	return p
}

func (client *V3Client) CreateScript(params CreateScriptParams) (Script, error) {
	type ResponseObject struct {
		Data Script   `json:"data"`
		Meta MetaData `json:"meta"`
	}

	var response ResponseObject

	path := client.constructURL("/content/scripts")

	err := client.Post(path, params, &response)
	if err != nil {
		return response.Data, err
	}

	return response.Data, nil

}

func (client *V3Client) UpdateScript(uuid string, params UpdateScriptParams) (Script, error) {
	type ResponseObject struct {
		Data Script   `json:"data"`
		Meta MetaData `json:"meta"`
	}

	var response ResponseObject

	updateScriptURL := client.constructURL("content", "scripts", uuid)

	err := client.Put(updateScriptURL, params, &response)
	if err != nil {
		return response.Data, err
	}

	return response.Data, nil
}

package bigcommerce

import "encoding/json"

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

func (client *Client) UpdateScript(uuid string, params UpdateScriptParams) (Script, error) {
	type ResponseObject struct {
		Data Script   `json:"data"`
		Meta MetaData `json:"meta"`
	}

	var response ResponseObject

	updateScriptURL := client.BaseURL.JoinPath("/content/scripts/" + uuid).String()
	resp, err := client.Get(updateScriptURL)
	if err != nil {
		return response.Data, err
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

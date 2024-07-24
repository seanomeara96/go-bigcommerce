package bigcommerce

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

type CustomURL struct {
	URL          string `json:"url,omitempty"`
	IsCustomized bool   `json:"is_customized,omitempty"`
}

type MetaData struct {
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Total       int   `json:"total"`
	Count       int   `json:"count"`
	PerPage     int   `json:"per_page"`
	CurrentPage int   `json:"current_page"`
	TotalPages  int   `json:"total_pages"`
	Links       Links `json:"links"`
}

type Links struct {
	Current string `json:"current"`
}

type ErrorPayload struct {
	Status   int    `json:"status"`
	Title    string `json:"title"`
	Type     string `json:"type"`
	Instance string `json:"instance"`
}

func expectStatusCode(expectedStatusCode int, response *http.Response) error {
	if response.StatusCode != expectedStatusCode {
		var errorPayload ErrorPayload
		if err := json.NewDecoder(response.Body).Decode(&errorPayload); err != nil {
			bytes, err := io.ReadAll(response.Body)
			if err != nil {
				return fmt.Errorf(
					"expected status code %d, received code: %d. There was a problem decoding the error payload: %s",
					expectedStatusCode,
					response.StatusCode,
					string(bytes),
				)
			}

			return fmt.Errorf(
				"expected status code %d, received code: %d. Could not decode or read response body. status: %s",
				expectedStatusCode,
				response.StatusCode,
				response.Status,
			)

		}
		return fmt.Errorf(
			"bigcommerce responded with status: %d, type: %s, title: %s, instance: %s",
			errorPayload.Status,
			errorPayload.Type,
			errorPayload.Title,
			errorPayload.Instance,
		)
	}
	return nil
}

func paramString(params interface{}) (string, error) {
	queryParamValues, err := query.Values(params)
	if err != nil {
		return "", err
	}
	var queryParams string = queryParamValues.Encode()
	if len(queryParams) > 0 {
		queryParams = "?" + queryParams
	}
	return queryParams, nil
}

func (client *Client) Version2Required() error {
	if client.Version() != 2 {
		return fmt.Errorf("need to be using version 2 api for this function")
	}
	return nil
}

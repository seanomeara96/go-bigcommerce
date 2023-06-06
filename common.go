package bigcommerce

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CustomURL struct {
	URL          string `json:"url"`
	IsCustomized bool   `json:"is_customized"`
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
			return fmt.Errorf(
				"expected status code %d, received code: %d. There was a problem decoding the error payload",
				expectedStatusCode,
				response.StatusCode,
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

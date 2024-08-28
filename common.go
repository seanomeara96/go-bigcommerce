package bigcommerce

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

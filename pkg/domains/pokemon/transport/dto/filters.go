package dto

type FilterRequest struct {
	FilterKey   string
	FilterValue string
}

type Filters struct {
	Filters []FilterRequest
}

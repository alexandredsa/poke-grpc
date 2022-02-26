package dto

type FilterRequest struct {
	FilterKey   string
	FilterValue string
}

type Filters struct {
	Filters []FilterRequest
}

func (f Filters) Empty() bool {
	return f.Filters == nil || len(f.Filters) == 0
}

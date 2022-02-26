package dto

type FilterRequest struct {
	Key   string
	Type  string
	Value string
}

type Filters struct {
	Filters []FilterRequest
}

func (f Filters) Empty() bool {
	return f.Filters == nil || len(f.Filters) == 0
}

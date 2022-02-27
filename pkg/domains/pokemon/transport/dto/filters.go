package dto

import "strings"

type FilterType int

const (
	FilterExactType FilterType = iota
	FilterRegexType
)

func (e FilterType) String() string {
	return [...]string{
		"exact",
		"regex",
		"",
	}[e]
}

var (
	filterSupportedTypes = map[string]FilterType{
		"exact": FilterExactType,
		"regex": FilterRegexType,
		"":      FilterRegexType,
	}
)

func ParseFilterTypeFromString(filterType string) (FilterType, bool) {
	fType, ok := filterSupportedTypes[strings.ToLower(filterType)]
	return fType, ok
}

type FilterRequest struct {
	Key   string
	Type  FilterType
	Value string
}

type Filters struct {
	Filters []FilterRequest
}

func (f Filters) Empty() bool {
	return f.Filters == nil || len(f.Filters) == 0
}

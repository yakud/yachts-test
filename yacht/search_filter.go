package yacht

type SearchFilter struct {
	Field string
	Value string
}

func Filter(field, value string) SearchFilter {
	return SearchFilter{
		Field: field,
		Value: value,
	}
}

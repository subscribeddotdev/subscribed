package query

type PaginationParams struct {
	Page  *int
	Limit *int
}

type Paginated[D any] struct {
	Total       int64
	PerPage     int64
	CurrentPage int64
	TotalPages  int64
	Data        D
}

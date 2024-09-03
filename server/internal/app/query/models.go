package query

type PaginationParams struct {
	page  int
	limit int
}

func NewPaginationParams(page, limit *int) PaginationParams {
	p := 1
	l := 15
	if page != nil {
		p = *page
	}

	if limit != nil {
		l = *limit
	}

	return PaginationParams{
		page:  p,
		limit: l,
	}
}

func (p *PaginationParams) Page() int {
	return p.page
}

func (p *PaginationParams) Limit() int {
	return p.limit
}

type Paginated[D any] struct {
	Total       int
	PerPage     int
	CurrentPage int
	TotalPages  int
	Data        D
}

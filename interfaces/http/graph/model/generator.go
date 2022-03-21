package model

func newPaginationInfo(total, count, limit, offset int) *PaginationInfo {
	page := 1
	pageLength := 1
	if limit > 0 {
		page = (offset / limit) + 1
		pageLength = (total / limit) + 1
	}

	return &PaginationInfo{
		Page:             page,
		PaginationLength: pageLength,
		HasNextPage:      offset+limit < total,
		HasPreviousPage:  offset > 0,
		Count:            count,
		TotalCount:       total,
	}
}

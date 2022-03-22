package pagination

import (
	"math"
)

type Pagination struct {
	Total int64
	PerPage int32
	CurrentPage int32
	LastPage int32
	From int32
	To int32
}

var (
	PerPageDefault int32 = 20
)

func NewPagination(page, perPage int32, total int64) *Pagination {
	if perPage <= 0 {
		perPage = PerPageDefault
	}

	lastPage := int32(math.Ceil(float64(total) / float64(perPage)))
	if page <= 0 {
		page = 1
	}
	// total=0, lastPage = 0
	if lastPage <= 0 {
		lastPage = 1
	}
	if page > lastPage {
		page = lastPage
	}

	var offset = (page - 1) * perPage

	return &Pagination{
		Total: total,
		PerPage: perPage,
		CurrentPage: page,
		LastPage: lastPage,
		From: offset,
		To: offset + perPage,
	}
}

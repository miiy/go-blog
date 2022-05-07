package paginate

import (
	"gorm.io/gorm"
	"math"
)

func Paginate(page, pageSize, total int) (scope func (db *gorm.DB) *gorm.DB, totalPage int) {
	totalPage = int(math.Ceil(float64(total) / float64(pageSize)))

	return func (db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}, totalPage
}

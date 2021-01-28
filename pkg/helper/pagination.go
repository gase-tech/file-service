package helper

import "gorm.io/gorm"

type Param struct {
	DB      *gorm.DB
	Page    int
	Limit   int
	OrderBy []string
	ShowSQL bool
}

// Pagination struct
type Pagination struct {
	TotalRecord int64       `json:"totalElements"`
	TotalPage   int         `json:"totalPages"`
	Data        interface{} `json:"content"`
	Limit       int         `json:"size"`
	Page        int         `json:"number"`
	PrevPage    int         `json:"prev_page"`
	NextPage    int         `json:"next_page"`
}

func Paginate(page int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case size > 100:
			size = 100
		case size <= 0:
			size = 10
		}

		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}

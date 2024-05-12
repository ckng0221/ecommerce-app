package utils

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"gorm.io/gorm"
)

// Pagination
//
// page: default 1, page-size: default 100.
func Paginate(r *http.Request) (func(db *gorm.DB) *gorm.DB, error) {
	pageString := r.URL.Query().Get("page")
	pageSizeString := r.URL.Query().Get("page-size")
	var page int
	var pageSize int
	var error error

	if pageString != "" && pageSizeString != "" {
		page, error = strconv.Atoi(pageString)
		if error != nil {
			return nil, errors.New("invalid page")
		}
		pageSize, error = strconv.Atoi(pageSizeString)
		if error != nil {
			return nil, errors.New("invalid page size")
		}
	}

	if page <= 0 {
		page = 1
	}

	switch {
	case pageSize <= 0:
		pageSize = 100
	}

	offset := (page - 1) * pageSize
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}, nil
}

func EmptyScope(db *gorm.DB) *gorm.DB {
	return db
}

type DefaultModel struct {
	ID        uint            `gorm:"primarykey" json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

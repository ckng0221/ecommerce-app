package utils

import (
	"errors"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func Paginate(r *http.Request) (func(db *gorm.DB) *gorm.DB, error) {
	pageString := r.URL.Query().Get("page")
	pageSizeString := r.URL.Query().Get("page-size")
	page, error := strconv.Atoi(pageString)
	if error != nil {
		return nil, errors.New("invalid page")
	}
	pageSize, error := strconv.Atoi(pageSizeString)
	if error != nil {
		return nil, errors.New("invalid page size")
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

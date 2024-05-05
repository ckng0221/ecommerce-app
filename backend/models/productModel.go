package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string  `gorm:"type:varchar(255)" json:"name"`
	Description   string  `gorm:"type:varchar(255)" json:"description"`
	UnitPrice     float32 `gorm:"type:decimal(10,2)" json:"unit_price"`
	StockQuantity int     `gorm:"default:0" json:"stock_quantity"`
	IsActive      bool    `gorm:"default:true" json:"is_active"`
}

type ProductUpdate struct {
	Name          *string          `json:"name"`
	Description   *string          `json:"decription"`
	UnitPrice     *decimal.Decimal `json:"unit_price"`
	StockQuantity *int             `json:"stock_quantity"`
	IsActive      *bool            `json:"is_active"`
}

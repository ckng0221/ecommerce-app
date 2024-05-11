package models

import (
	"ecommerce-app/utils"

	"github.com/shopspring/decimal"
)

type Product struct {
	utils.DefaultModel
	Name          string  `gorm:"type:varchar(255)" json:"name"`
	Description   string  `gorm:"type:varchar(255)" json:"description"`
	UnitPrice     float32 `gorm:"type:decimal(10,2)" json:"unit_price"`
	Currency      string  `gorm:"type:char(3); default:myr" json:"currency"`
	StockQuantity int     `gorm:"default:0" json:"stock_quantity"`
	IsActive      bool    `gorm:"default:true" json:"is_active"`
}

type ProductUpdate struct {
	Name          *string          `json:"name"`
	Description   *string          `json:"decription"`
	UnitPrice     *decimal.Decimal `json:"unit_price"`
	Currency      *string          `json:"currency"`
	StockQuantity *int             `json:"stock_quantity"`
	IsActive      *bool            `json:"is_active"`
}

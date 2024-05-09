package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Quantity  int     `json:"quantity"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID"`
	UserID    uint    `json:"user_id"`
}

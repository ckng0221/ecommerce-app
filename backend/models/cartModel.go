package models

import "ecommerce-app/utils"

type Cart struct {
	utils.DefaultModel
	Quantity  int      `json:"quantity"`
	ProductID uint     `json:"product_id"`
	Product   *Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	UserID    uint     `json:"user_id"`
}

type CartUpdate struct {
	Quantity  *int     `json:"quantity"`
	ProductID *uint    `json:"product_id"`
	Product   *Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

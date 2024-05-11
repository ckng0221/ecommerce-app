package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	OrderID         uint   `json:"order_id"`
	Order           Order  `gorm:"foreignKey:OrderID"`
	StripeSessionID string `gorm:"type:char(66)"`
}

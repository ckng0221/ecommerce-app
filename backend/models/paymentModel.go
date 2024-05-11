package models

import "ecommerce-app/utils"

type Payment struct {
	utils.DefaultModel
	OrderID         uint   `json:"order_id"`
	Order           *Order `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	StripeSessionID string `gorm:"type:char(66)" json:"stripe_session_id"`
}

package models

import (
	"ecommerce-app/utils"
	"time"
)

type Payment struct {
	utils.DefaultModel
	OrderID         uint       `json:"order_id"`
	Order           *Order     `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	StripeSessionID string     `gorm:"type:char(66)" json:"stripe_session_id"`
	IsComplete      bool       `gorm:"default:false" json:"is_complete"`
	PaymentAt       *time.Time `json:"payment_at"`
}

type PaymentUpdate struct {
	OrderID         *uint      `json:"order_id"`
	StripeSessionID *string    `json:"stripe_session_id"`
	IsComplete      *bool      `json:"is_complete"`
	PaymentAt       *time.Time `json:"payment_at"`
}

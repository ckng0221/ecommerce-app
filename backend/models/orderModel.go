package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID      uint        `json:"user_id"`
	User        User        `gorm:"foreignKey:UserID"`
	AddressID   uint        `json:"address_id"`
	Address     Address     `gorm:"foreignKey:AddressID"`
	PaymentAt   time.Time   `json:"payment_at"`
	OrderStatus OrderStatus `gorm:"type:enum('to_pay', 'to_ship', 'to_receive', 'to_review', 'complete'); default:to_pay"`
}

type OrderUpdate struct {
	gorm.Model
	UserID    *uint      `json:"user_id"`
	AddressID *Address   `json:"address_id"`
	PaymentAt *time.Time `json:"payment_at"`
}

type OrderStatus string

const (
	ToPay     OrderStatus = "to_pay"
	ToShip    OrderStatus = "to_ship"
	ToReceive OrderStatus = "to_receive"
	ToReview  OrderStatus = "to_review"
	Complete  OrderStatus = "complete"
)

type OrderItem struct {
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint    `json:"order_id"`
	Order     Order   `gorm:"foreignKey:OrderID"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
	Price     float32 `gorm:"type:decimal(10,2)" json:"price"`
	Currency  string  `gorm:"type:char(3); default:myr" json:"currency"`
}

// NOTE: Repeat price and currency, as for snapshot of actual order price and currency used
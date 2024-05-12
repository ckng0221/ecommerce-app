package models

import (
	"ecommerce-app/utils"
	"time"
)

type Order struct {
	utils.DefaultModel
	UserID      uint         `json:"user_id"`
	User        *User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	AddressID   uint         `json:"address_id"`
	Address     *Address     `gorm:"foreignKey:AddressID" json:"address,omitempty"`
	PaymentAt   *time.Time   `json:"payment_at"`
	OrderStatus *OrderStatus `gorm:"type:enum('to_pay', 'to_ship', 'to_receive', 'to_review', 'complete'); default:to_pay" json:"order_status"`
	OrderItems  *[]OrderItem `json:"order_items,omitempty"`
}

type OrderUpdate struct {
	utils.DefaultModel
	UserID      *uint       `json:"user_id"`
	AddressID   *uint       `json:"address_id"`
	PaymentAt   *time.Time  `json:"payment_at"`
	OrderStatus OrderStatus `json:"order_status"`
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
	ID        uint     `gorm:"primaryKey" json:"id"`
	OrderID   uint     `json:"order_id"`
	Order     *Order   `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	ProductID uint     `json:"product_id"`
	Product   *Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity  int      `json:"quantity"`
	Price     float32  `gorm:"type:decimal(10,2)" json:"price"`
	Currency  string   `gorm:"type:char(3); default:myr" json:"currency"`
}

// NOTE: Repeat price and currency, as for snapshot of actual order price and currency used

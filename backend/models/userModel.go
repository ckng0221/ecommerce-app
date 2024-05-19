package models

import (
	"ecommerce-app/utils"
)

type Role string

const (
	Admin  Role = "admin"
	Member Role = "member"
)

type User struct {
	utils.DefaultModel
	Name             string   `gorm:"type:varchar(100)" json:"name"`
	Email            string   `gorm:"type:varchar(100)" json:"email"`
	Password         *string  `gorm:"type:varchar(255)" json:"-"`
	Role             Role     `gorm:"type:enum('admin', 'member'); default:member" json:"role"`
	ProfilePic       *string  `gorm:"type:varchar(255)" json:"profile_pic"`
	Sub              *string  `gorm:"type:varchar(100); unique" json:"sub"`
	Carts            *[]Cart  `gorm:"foreignKey:UserID" json:"carts,omitempty"`
	DefaultAddressID *uint    `json:"default_address_id"`
	DefaultAddress   *Address `gorm:"foreignKey:DefaultAddressID" json:"default_address,omitempty"`
}
type UserUpdate struct {
	Name             *string `json:"name"`
	Email            *string `json:"email"`
	Password         *string `json:"-"`
	Role             *Role   `json:"role"`
	ProfilePic       *string `json:"profile_pic"`
	DefaultAddressID *uint
}

type UserSignUp struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GoogleProfile struct {
	Sub            string
	Email          string
	Verified_email bool
	Name           string
	Given_name     string
	Family_name    string
	Picture        string
	Locale         string
}

type Address struct {
	utils.DefaultModel
	UserID        uint   `json:"user_id"`
	User          *User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Street        string `gorm:"type:varchar(100)" json:"street"`
	City          string `gorm:"type:varchar(50)" json:"city"`
	State         string `gorm:"type:varchar(50)" json:"state"`
	Zip           string `gorm:"type:varchar(10)" json:"zip"`
	ContactNumber string `gorm:"type:varchar(20)" json:"contact_number"`
}

var Roles = [2]string{string(Admin), string(Member)}

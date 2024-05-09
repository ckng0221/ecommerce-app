package models

import "gorm.io/gorm"

type Role string

const (
	Admin  Role = "admin"
	Member Role = "member"
)

type User struct {
	gorm.Model
	Name             string  `gorm:"type:varchar(100)"`
	Email            string  `gorm:"type:varchar(100)"`
	Password         *string `gorm:"type:varchar(255)" json:"-"`
	Role             Role    `gorm:"type:enum('admin', 'member'); default:member"`
	ProfilePic       *string `gorm:"type:varchar(255)"`
	Sub              *string `gorm:"type:varchar(100); unique"`
	DefaultAddressID *uint
	DefaultAddress   *Address
	Carts            []Cart `gorm:"foreignKey:UserID"`
}
type UserUpdate struct {
	Name             *string `json:"name"`
	Email            *string `json:"email"`
	Password         *string `json:"-"`
	Role             *Role   `json:"role"`
	ProfilePic       *string `json:"profile_pic"`
	DefaultAddressID *uint
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
	gorm.Model
	UserID uint   `json:"user_id"`
	User   User   `gorm:"foreignKey:UserID"`
	Street string `gorm:"type:varchar(100)"`
	City   string `gorm:"type:varchar(50)"`
	State  string `gorm:"type:varchar(50)"`
	Zip    string `gorm:"type:varchar(10)"`
}

var Roles = [2]string{string(Admin), string(Member)}

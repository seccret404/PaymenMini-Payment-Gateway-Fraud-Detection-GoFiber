package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name	string `json:"name"`
	Email	string 	`json:"email" gorm:"unique"`
	Password string `json:"-"`
	Role	string	`json:"role" gorm:"default:buyer"`
	Products	[]Product `json:"products"`
	Carts	[]Cart	`json:"cart"`
}
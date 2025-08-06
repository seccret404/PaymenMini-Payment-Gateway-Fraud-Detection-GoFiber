package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string	`json:"description"`
	Quantity	 	string `json:"quantity"`

	Carts	[]Cart `json:"carts"`

	UserID	uint `json:"user_id"`
	User		User `json:"user"`

}
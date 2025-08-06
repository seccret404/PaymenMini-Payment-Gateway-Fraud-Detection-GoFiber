package domain

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Quantity	string	`json:"quantity"`
	ProductID	 int `json:"product_id"`
	Product 	Product `json:"product"`
	UserID	uint	`json:"user_id"`
	User		User	`json:"user"`

}
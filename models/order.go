package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	User_id      string
	Item_ids     string
	Total_amount float32
}

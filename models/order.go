package models

import (
	"time"
)

type Order struct {
	ID           string `gorm:"column:id;primaryKey;" json:"id" validate:"required"`
	User_id      string
	Item_ids     string
	Total_amount float32
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

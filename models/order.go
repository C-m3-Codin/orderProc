package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	// ID                   string `gorm:"column:id;primaryKey;" json:"id" validate:"required"`
	ID                   uuid.UUID `gorm:"column:id;primaryKey;type:uuid;" json:"id" validate:"required"`
	User_id              string
	Item_ids             string
	Total_amount         float32
	Status               int
	OrderReceived        time.Time `gorm:"column:order_received;"`
	OrderProcessingStart time.Time `gorm:"column:order_processing_start;"`
	OrderCompleted       time.Time `gorm:"column:order_completed;"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            time.Time
}

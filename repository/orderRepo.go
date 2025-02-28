package repository

import "gorm.io/gorm"

type OrderRepo struct {
	DB *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {

	or := OrderRepo{
		DB: db,
	}

	return or
}

func (order OrderRepo) GetOrder(orderId string) {
	return
}

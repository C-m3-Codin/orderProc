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

func (orderRep OrderRepo) GetOrder(orderId uint) (order gorm.Model, err error) {

	order.ID = orderId
	err = orderRep.DB.First(order).Error

	return
}

func (orderRep OrderRepo) CreateOrder(order gorm.Model) (err error) {

	result := orderRep.DB.Create(&order)
	err = result.Error

	return
}

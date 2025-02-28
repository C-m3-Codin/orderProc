package repository

import (
	"c-m3-codin/ordProc/models"
	"fmt"

	"gorm.io/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {

	or := OrderRepo{
		DB: db,
	}

	return or
}

func (orderRep OrderRepo) GetOrder(orderId string) (order models.Order, err error) {
	fmt.Println("got here")
	order.ID = orderId
	fmt.Println(order)
	err = orderRep.DB.First(&order).Error

	return
}

func (orderRep OrderRepo) CreateOrder(order models.Order) (err error) {

	result := orderRep.DB.Create(&order)
	err = result.Error
	return
}

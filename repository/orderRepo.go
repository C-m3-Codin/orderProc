package repository

import (
	"c-m3-codin/ordProc/models"

	"github.com/google/uuid"
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
	order.ID, err = uuid.Parse(orderId)
	if err != nil {
		return
	}
	// fmt.Println(order)
	err = orderRep.DB.First(&order).Error

	return
}

func (orderRep OrderRepo) CreateOrder(order models.Order) (err error) {

	result := orderRep.DB.Create(&order)
	err = result.Error
	return
}

func (OrderRepo OrderRepo) UpdateOrder(order models.Order) (err error) {
	err = OrderRepo.DB.Model(&models.Order{}).Where("id = ?", order.ID).Updates(models.Order{Status: 2, OrderCompleted: order.OrderCompleted}).Error
	return
}

func (OrderRepo OrderRepo) GetUnproccessedOrders() (orders []models.Order, err error) {
	err = OrderRepo.DB.Where("status < ? ", 2).Find(&orders).Error
	return
}

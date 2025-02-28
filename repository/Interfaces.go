package repository

import "gorm.io/gorm"

type OrderRepoInteface interface {
	GetOrder(orderid string) (order gorm.Model, err error)
	CreateOrder(order gorm.Model) (err error)
}

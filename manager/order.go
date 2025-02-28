package manager

import (
	"c-m3-codin/ordProc/models"
	"c-m3-codin/ordProc/repository"
)

type OrderManager struct {
	repo repository.OrderRepo
}

type OrderManagerInterface interface {
	AcceptOrder(ord models.Order) (err error)
	GetOrder(ordid uint) (order models.Order, err error)
}

func NewOrderhandler(repo repository.OrderRepo) OrderManager {
	return OrderManager{
		repo: repo,
	}
}

func (o OrderManager) AcceptOrder(ord models.Order) (err error) {
	err = o.repo.CreateOrder(ord)
	return
}

func (o OrderManager) GetOrder(ordid uint) (order models.Order, err error) {
	order, err = o.repo.GetOrder(ordid)
	return
}

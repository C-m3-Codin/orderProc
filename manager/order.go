package manager

import (
	"c-m3-codin/ordProc/models"
	"c-m3-codin/ordProc/repository"
	"c-m3-codin/ordProc/services"

	"github.com/google/uuid"
)

type OrderManager struct {
	repo repository.OrderRepo
	q    services.Queue
}

type OrderManagerInterface interface {
	AcceptOrder(ord models.Order) (orderId string, err error)
	GetOrder(ordid string) (order models.Order, err error)
}

func NewOrderhandler(repo repository.OrderRepo, q services.Queue) OrderManager {
	return OrderManager{
		repo: repo,
		q:    q,
	}
}

func (o OrderManager) AcceptOrder(ord models.Order) (orderId string, err error) {
	orderId = uuid.NewString()
	ord.ID = orderId
	// err = o.repo.CreateOrder(ord)
	services.CacheReceivedOrder[orderId] = true
	o.q.PendingQueue <- ord
	return
}

func (o OrderManager) GetOrder(ordid string) (order models.Order, err error) {
	order, err = o.repo.GetOrder(ordid)
	return
}

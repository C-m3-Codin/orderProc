package manager

import (
	"c-m3-codin/ordProc/models"
	"c-m3-codin/ordProc/repository"
	"c-m3-codin/ordProc/services"
	"fmt"
	"time"

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
	ord.OrderReceived = time.Now()
	// err = o.repo.CreateOrder(ord)
	services.CacheReceivedOrders.Set(orderId, true)
	o.q.PendingQueue <- ord
	return
}

func (o OrderManager) GetOrder(ordid string) (order models.Order, err error) {
	order, err = o.repo.GetOrder(ordid)
	return
}

func (O OrderManager) LoadUpUnproccessed() {
	ords, err := O.repo.GetUnproccessedOrders()
	if err != nil {
		fmt.Println("Couldnt fetch unproccessed orders", err)
		return
	}
	for _, ord := range ords {
		O.q.Processing <- ord
	}
}

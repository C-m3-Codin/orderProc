package workers

import (
	"c-m3-codin/ordProc/repository"
	"c-m3-codin/ordProc/services"
	"fmt"
)

type WorkerPool struct {
	Q           services.Queue
	WorkerCount int
	orderRepo   repository.OrderRepo
}

type WorkerPoolInterface interface {
	StartCreateOrderWorkers()
	StopWorkers()
	GetMetrics()
}

func NewWorkerPool(count int, q services.Queue, orderRepo repository.OrderRepo) WorkerPool {
	wp := WorkerPool{Q: q, WorkerCount: count, orderRepo: orderRepo}
	return wp
}

func (wp WorkerPool) StartCreateOrderWorkers() {
	for i := range wp.WorkerCount {
		go wp.ListenForOrders(i)
	}
	return
}

func (wp WorkerPool) StopWorkers() {
	return
}

func (wp WorkerPool) GetMetrics() {
	return
}

func (wp WorkerPool) ListenForOrders(workerID int) {
	fmt.Println("Worker started, id:", workerID)

	for ord := range wp.Q.PendingQueue {
		fmt.Println("Worker", workerID, "received order:", ord)
		err := wp.orderRepo.CreateOrder(ord)
		if err != nil {
			fmt.Println("Error processing order, requeuing:", ord)
			wp.Q.PendingQueue <- ord
		}
	}

	fmt.Println("Worker", workerID, "stopped, channel closed.")
}

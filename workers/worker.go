package workers

import (
	"c-m3-codin/ordProc/repository"
	"c-m3-codin/ordProc/services"
	"fmt"
	"math/rand"
	"time"
)

type WorkerPool struct {
	Q           services.Queue
	WorkerCount int
	orderRepo   repository.OrderRepo
}

type WorkerPoolInterface interface {
	StartCreateOrderWorkers()
	StartProccessOrderWorkers()
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

func (wp WorkerPool) StartProccessOrderWorkers() {
	for i := range wp.WorkerCount {
		go wp.ProcccessOrders(i)
	}

}

func (wp WorkerPool) StopWorkers() {
	close(wp.Q.PendingQueue)
	close(wp.Q.Processing)

}

func (wp WorkerPool) GetMetrics() {
	return
}

func (wp WorkerPool) ListenForOrders(workerID int) {
	fmt.Println("Worker started, id:", workerID)

	for ord := range wp.Q.PendingQueue {
		fmt.Println("Worker", workerID, "received order:", ord)
		ord.Status = 1
		ord.OrderProcessingStart = time.Now()
		err := wp.orderRepo.CreateOrder(ord)
		if err != nil {
			ord.Status = 0
			fmt.Println("Error Creating order, requeuing:", ord)
			wp.Q.PendingQueue <- ord
		} else {
			delete(services.CacheReceivedOrder, ord.ID)
			wp.Q.Processing <- ord
		}
	}

	fmt.Println("Worker", workerID, "stopped, channel closed.")
}

func (wp WorkerPool) ProcccessOrders(workerID int) {
	fmt.Println("Worker started, id:", workerID)

	for ord := range wp.Q.Processing {
		fmt.Println("Worker", workerID, "received order:", ord)

		// simulating time for proccessing the order
		sleepDuration := rand.Intn(10) + 1
		fmt.Println("Gonna take time to proccess order : ", sleepDuration)
		time.Sleep(time.Duration(sleepDuration) * time.Second)
		ord.OrderCompleted = time.Now()
		err := wp.orderRepo.UpdateOrder(ord)
		if err != nil {
			fmt.Println("Error Processing order, requeuing:", ord)
			wp.Q.Processing <- ord
		} else {
			fmt.Println("Order Processing done")
		}
	}

	fmt.Println("Worker", workerID, "stopped, channel closed.")
}

package services

import "c-m3-codin/ordProc/models"

type Queue struct {
	PendingQueue chan models.Order
	Processing   chan models.Order
}

func NewQueue() (q Queue) {
	pend := make(chan models.Order, 200000)
	proc := make(chan models.Order, 200000)

	q = Queue{
		PendingQueue: pend,
		Processing:   proc,
	}
	return q

}

package manager

import (
	"c-m3-codin/ordProc/models"
	"fmt"
	"time"
)

func (o OrderManager) GetPending() (metrics models.Metrics, err error) {
	metrics.PendingCount = int64(len(o.q.PendingQueue))
	return
}

func (o OrderManager) GetProccessedCount() (metrics models.Metrics, err error) {
	metrics.Proccessed, err = o.repo.GetProccessedCount()
	return
}

func (o OrderManager) GetCompletedCount() (metrics models.Metrics, err error) {
	metrics.Completed, err = o.repo.GetCompletedCount()
	return
}

func (o OrderManager) GetTotalCount() (metrics models.Metrics, err error) {
	metrics.Completed, err = o.repo.GetTotalCount()
	return
}

func (o OrderManager) GetAverageProcessingTimeCount() (metrics models.Metrics, err error) {

	metrics.PendingCount = int64(len(o.q.PendingQueue))
	totalTime, err := o.repo.GetAverageProcessingTimeCount()
	if err != nil {
		fmt.Println("erroro.repo.GetAverageProcessingTimeCount()")
		return
	}
	metrics.TotalCount, err = o.repo.GetTotalCount()
	if err != nil {
		fmt.Println("erroro.repo.GetTotalCount()")
		return
	}
	ms := totalTime.Milliseconds()
	averageTime := ms / metrics.TotalCount
	metrics.AverageProcessingTime = time.Duration(averageTime)

	return
}

func (o OrderManager) GetAllMetrics() (metrics models.Metrics, err error) {
	fmt.Println("got here")
	totalTime, err := o.repo.GetAverageProcessingTimeCount()
	metrics.PendingCount = int64(len(o.q.PendingQueue))
	if err != nil {
		fmt.Println("error o.repo.GetAverageProcessingTimeCount()")
		return
	}
	metrics.Proccessed, err = o.repo.GetProccessedCount()
	if err != nil {
		fmt.Println("error o.repo.GetProccessedCount()")
		return
	}
	metrics.Completed, err = o.repo.GetCompletedCount()
	if err != nil {
		fmt.Println("error o.repo.GetCompletedCount()")
		return
	}
	metrics.TotalCount, err = o.repo.GetTotalCount()
	if err != nil {
		fmt.Println("erroro.repo.GetTotalCount()")
		return
	}
	ms := totalTime.Milliseconds()
	averageTime := ms / metrics.TotalCount
	metrics.AverageProcessingTime = time.Duration(averageTime)

	return
}

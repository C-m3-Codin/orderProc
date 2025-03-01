package main

import (
	"c-m3-codin/ordProc/handler"
	"c-m3-codin/ordProc/manager"
	"c-m3-codin/ordProc/models"
	"c-m3-codin/ordProc/repository"
	"c-m3-codin/ordProc/services"
	"c-m3-codin/ordProc/workers"
	"strconv"

	"github.com/alphadose/haxmap"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = services.GetConnections("postgres")

}

func main() {
	services.CacheReceivedOrders = haxmap.New[string, bool]()

	q := services.NewQueue()
	orderRepo := repository.NewOrderRepo(db)
	orderManager := manager.NewOrderhandler(orderRepo, q)
	orderHandler := handler.NewOrderhandler(orderManager)
	workerPool := workers.NewWorkerPool(1000, q, orderRepo)
	workerPool.StartCreateOrderWorkers()
	workerPool.StartProccessOrderWorkers()
	orderManager.LoadUpUnproccessed()
	func() {
		for i := range 200 {
			a := models.Order{
				Total_amount: float32(i) * 23.0,
				User_id:      strconv.Itoa(i),
				Item_ids:     strconv.Itoa(21 * i),
			}
			go orderManager.AcceptOrder(a)
		}

	}()
	r := gin.Default()
	r.GET("/ping", handler.Ping)
	r.GET("/order/:id", orderHandler.GetOrders)
	r.POST("/order", orderHandler.PostOrders)
	r.Run() // listen and serve on 0.0.0.0:8080

}

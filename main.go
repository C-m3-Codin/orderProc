package main

import (
	"c-m3-codin/ordProc/handler"
	"c-m3-codin/ordProc/manager"
	"c-m3-codin/ordProc/repository"
	"c-m3-codin/ordProc/services"
	"c-m3-codin/ordProc/workers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = services.GetConnections("sqlite3")

}

func main() {
	q := services.NewQueue()
	orderRepo := repository.NewOrderRepo(db)
	orderManager := manager.NewOrderhandler(orderRepo, q)
	orderHandler := handler.NewOrderhandler(orderManager)
	workerPool := workers.NewWorkerPool(1000, q, orderRepo)
	workerPool.StartCreateOrderWorkers()

	r := gin.Default()
	r.GET("/ping", handler.Ping)
	r.GET("/order/:id", orderHandler.GetOrders)
	r.POST("/order", orderHandler.PostOrders)
	r.Run() // listen and serve on 0.0.0.0:8080

}

package main

import (
	"c-m3-codin/ordProc/handler"
	"c-m3-codin/ordProc/manager"
	"c-m3-codin/ordProc/repository"
	"c-m3-codin/ordProc/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = services.GetConnections("sqlite3")

}

func main() {
	orderRepo := repository.NewOrderRepo(db)
	orderManager := manager.NewOrderhandler(orderRepo)
	orderHandler := handler.NewOrderhandler(orderManager)

	r := gin.Default()
	r.GET("/ping", handler.Ping)
	r.GET("/order", orderHandler.GetOrders)
	r.Run() // listen and serve on 0.0.0.0:8080

}

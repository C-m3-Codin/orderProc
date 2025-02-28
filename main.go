package main

import (
	"c-m3-codin/ordProc/handler"
	"c-m3-codin/ordProc/repository"
	"c-m3-codin/ordProc/services"

	"github.com/gin-gonic/gin"
)

var orderRepo repository.OrderRepo

func init() {
	db := services.GetConnections("sqlite3")
	orderRepo = repository.NewOrderRepo(db)
}

func main() {

	r := gin.Default()
	r.GET("/ping", handler.Ping)
	r.Run() // listen and serve on 0.0.0.0:8080

}

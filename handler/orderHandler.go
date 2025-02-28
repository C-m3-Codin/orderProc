package handler

import (
	"c-m3-codin/ordProc/manager"
	"c-m3-codin/ordProc/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderManager manager.OrderManager
}

type Order interface {
	GetOrder(id string)
}

func NewOrderhandler(orderManager manager.OrderManager) OrderHandler {
	return OrderHandler{
		OrderManager: orderManager,
	}

}

func (ord OrderHandler) GetOrders(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("Received order id ", id)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// userId := uint(id)
	order, err := ord.OrderManager.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(200, order)
}

func (ord OrderHandler) PostOrders(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderId, err := ord.OrderManager.AcceptOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	returnMessage := fmt.Sprintf("Received order %s ", orderId)
	c.JSON(200, returnMessage)
}

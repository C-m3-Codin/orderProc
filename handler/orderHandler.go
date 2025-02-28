package handler

import (
	"c-m3-codin/ordProc/manager"
	"c-m3-codin/ordProc/models"
	"net/http"
	"strconv"

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

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := uint(id)
	order, err := ord.OrderManager.GetOrder(userId)
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
	err := ord.OrderManager.AcceptOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(200, "Created your order")
}

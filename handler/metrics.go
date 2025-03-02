package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ord OrderHandler) GetMetrics(c *gin.Context) {
	fmt.Println("got here")
	metrics, err := ord.OrderManager.GetAllMetrics()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, metrics)
}

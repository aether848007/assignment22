package handler

import (
	"api-gateway/utils"
	"github.com/gin-gonic/gin"
)

const (
	inventoryServiceURL = "http://localhost:8081"
	orderServiceURL     = "http://localhost:8082"
)

func ProxyInventory(c *gin.Context) {
	utils.ForwardRequest(c, inventoryServiceURL)
}

func ProxyOrder(c *gin.Context) {
	utils.ForwardRequest(c, orderServiceURL)
}

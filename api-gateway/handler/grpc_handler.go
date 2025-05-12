package handler

import (
	"api-gateway/service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	pbinv "inventory-service/proto/inventory-service/proto"
	pbord "order-service/proto/order-service/proto"
)

// ------------------- INVENTORY -------------------

func CreateProduct(c *gin.Context) {
	var req struct {
		Name     string  `json:"name"`
		Category string  `json:"category"`
		Stock    int32   `json:"stock"`
		Price    float64 `json:"price"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := service.InventoryClient.CreateProduct(context.Background(), &pbinv.CreateProductRequest{
		Name:     req.Name,
		Category: req.Category,
		Stock:    req.Stock,
		Price:    req.Price,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp.Product)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	resp, err := service.InventoryClient.GetProductByID(context.Background(), &pbinv.GetProductRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Product)
}

func ListProducts(c *gin.Context) {
	resp, err := service.InventoryClient.ListProducts(context.Background(), &pbinv.ListProductsRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Products)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Name     string  `json:"name"`
		Category string  `json:"category"`
		Stock    int32   `json:"stock"`
		Price    float64 `json:"price"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := service.InventoryClient.UpdateProduct(context.Background(), &pbinv.UpdateProductRequest{
		Id:       id,
		Name:     req.Name,
		Category: req.Category,
		Stock:    req.Stock,
		Price:    req.Price,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	_, err := service.InventoryClient.DeleteProduct(context.Background(), &pbinv.DeleteProductRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

// ------------------- ORDER -------------------

func CreateOrder(c *gin.Context) {
	var req struct {
		UserID    string `json:"user_id"`
		ProductID string `json:"product_id"`
		Quantity  int32  `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := service.OrderClient.CreateOrder(context.Background(), &pbord.CreateOrderRequest{
		UserId: req.UserID,
		Items: []*pbord.OrderItem{
			{
				ProductId: req.ProductID,
				Quantity:  req.Quantity,
			},
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp.Order)
}

func GetOrderByID(c *gin.Context) {
	orderID := c.Param("id")
	resp, err := service.OrderClient.GetOrderByID(context.Background(), &pbord.GetOrderRequest{
		OrderId: orderID,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Order)
}

func UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")
	var req struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}
	resp, err := service.OrderClient.UpdateOrderStatus(context.Background(), &pbord.UpdateOrderStatusRequest{
		OrderId: orderID,
		Status:  req.Status,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Order)
}

func ListUserOrders(c *gin.Context) {
	userID := c.Query("user_id")
	resp, err := service.OrderClient.ListUserOrders(context.Background(), &pbord.ListOrdersRequest{
		UserId: userID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Orders)
}

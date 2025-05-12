package router

import (
	"api-gateway/handler"
	"api-gateway/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(middleware.Logger())

	api := r.Group("/api")

	// Inventory
	api.POST("/products", handler.CreateProduct)
	api.GET("/products/:id", handler.GetProductByID)
	api.GET("/products", handler.ListProducts)
	api.PUT("/products/:id", handler.UpdateProduct)
	api.DELETE("/products/:id", handler.DeleteProduct)

	// Order
	api.POST("/orders", handler.CreateOrder)
	api.GET("/orders/:id", handler.GetOrderByID)
	api.PATCH("/orders/:id", handler.UpdateOrderStatus)
	api.GET("/orders", handler.ListUserOrders) // ?user_id=<id>
}

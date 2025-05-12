package main

import (
	"api-gateway/router"
	"api-gateway/service"
	"github.com/gin-gonic/gin"
)

func main() {
	service.InitGRPCClients()

	r := gin.Default()
	router.SetupRoutes(r)

	r.Run(":8080")
}

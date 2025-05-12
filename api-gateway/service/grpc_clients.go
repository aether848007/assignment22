package service

import (
	pbinv "inventory-service/proto/inventory-service/proto"
	"log"
	pbord "order-service/proto/order-service/proto"

	"google.golang.org/grpc"
)

var (
	InventoryClient pbinv.InventoryServiceClient
	OrderClient     pbord.OrderServiceClient
)

func InitGRPCClients() {
	invConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("❌ Failed to connect to inventory-service: %v", err)
	}
	InventoryClient = pbinv.NewInventoryServiceClient(invConn)

	orderConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("❌ Failed to connect to order-service: %v", err)
	}
	OrderClient = pbord.NewOrderServiceClient(orderConn)
}

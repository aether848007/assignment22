package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	grpcstd "google.golang.org/grpc"
	"inventory-service/internal/interface/grpc"
	"inventory-service/internal/interface/repository"
	"inventory-service/internal/usecase"
	pb "inventory-service/proto/inventory-service/proto"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	db := client.Database("inventory_db")

	repo := repository.NewMongoRepo(db)
	uc := usecase.NewInventoryUsecase(repo)
	grpcServer := grpc.NewInventoryGRPCServer(uc)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpcstd.NewServer()
	pb.RegisterInventoryServiceServer(s, grpcServer)

	fmt.Println("✅ Inventory gRPC server is running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

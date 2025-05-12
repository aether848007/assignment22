package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"order-service/internal/interface/grpc"
	"order-service/internal/interface/repesitory"
	"order-service/internal/usecase"
	pb "order-service/proto/order-service/proto"

	grpcstd "google.golang.org/grpc"
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
	db := client.Database("order_db")

	repo := repository.NewMongoOrderRepo(db)
	uc := usecase.NewOrderUsecase(repo)
	grpcServer := grpc.NewOrderGRPCServer(uc)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpcstd.NewServer()
	pb.RegisterOrderServiceServer(s, grpcServer)

	fmt.Println("✅ Order gRPC server is running on :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package grpc

import (
	"context"

	"order-service/internal/entity"
	"order-service/internal/usecase"
	pb "order-service/proto/order-service/proto"
)

type OrderGRPCServer struct {
	pb.UnimplementedOrderServiceServer
	Usecase *usecase.OrderUsecase
}

func NewOrderGRPCServer(uc *usecase.OrderUsecase) *OrderGRPCServer {
	return &OrderGRPCServer{Usecase: uc}
}

func convertToProtoOrder(o *entity.Order) *pb.Order {
	return &pb.Order{
		Id:     o.ID.Hex(),
		Status: o.Status,
		Items: []*pb.OrderItem{
			{
				ProductId: o.ProductID,
				Quantity:  int32(o.Quantity),
			},
		},
	}
}

func (s *OrderGRPCServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	order := &entity.Order{
		ProductID: req.GetItems()[0].GetProductId(),
		Quantity:  int(req.GetItems()[0].GetQuantity()),
		Status:    "created",
	}
	if err := s.Usecase.Create(ctx, order); err != nil {
		return nil, err
	}
	return &pb.OrderResponse{Order: convertToProtoOrder(order)}, nil
}

func (s *OrderGRPCServer) GetOrderByID(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	order, err := s.Usecase.GetByID(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}
	return &pb.OrderResponse{Order: convertToProtoOrder(order)}, nil
}

func (s *OrderGRPCServer) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest) (*pb.OrderResponse, error) {
	if err := s.Usecase.UpdateStatus(ctx, req.GetOrderId(), req.GetStatus()); err != nil {
		return nil, err
	}
	order, err := s.Usecase.GetByID(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}
	return &pb.OrderResponse{Order: convertToProtoOrder(order)}, nil
}

func (s *OrderGRPCServer) ListUserOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := s.Usecase.List(ctx)
	if err != nil {
		return nil, err
	}
	var pbOrders []*pb.Order
	for _, o := range orders {
		pbOrders = append(pbOrders, convertToProtoOrder(o))
	}
	return &pb.ListOrdersResponse{Orders: pbOrders}, nil
}

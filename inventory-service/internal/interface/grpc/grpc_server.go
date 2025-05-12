package grpc

import (
	"context"
	"inventory-service/internal/entity"
	"inventory-service/internal/usecase"
	pb "inventory-service/proto/inventory-service/proto"
)

type InventoryGRPCServer struct {
	pb.UnimplementedInventoryServiceServer
	Usecase *usecase.InventoryUsecase
}

func NewInventoryGRPCServer(uc *usecase.InventoryUsecase) *InventoryGRPCServer {
	return &InventoryGRPCServer{Usecase: uc}
}

func (s *InventoryGRPCServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	product := &entity.Product{
		Name:     req.GetName(),
		Category: req.GetCategory(),
		Price:    req.GetPrice(),
		Stock:    int(req.GetStock()),
	}
	err := s.Usecase.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return &pb.ProductResponse{
		Product: &pb.Product{
			Id:       product.ID,
			Name:     product.Name,
			Category: product.Category,
			Price:    product.Price,
			Stock:    int32(product.Stock),
		},
	}, nil
}

func (s *InventoryGRPCServer) GetProductByID(ctx context.Context, req *pb.GetProductRequest) (*pb.ProductResponse, error) {
	product, err := s.Usecase.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.ProductResponse{
		Product: &pb.Product{
			Id:       product.ID,
			Name:     product.Name,
			Category: product.Category,
			Price:    product.Price,
			Stock:    int32(product.Stock),
		},
	}, nil
}

func (s *InventoryGRPCServer) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.ProductResponse, error) {
	product := &entity.Product{
		Name:     req.GetName(),
		Category: req.GetCategory(),
		Price:    req.GetPrice(),
		Stock:    int(req.GetStock()),
	}
	err := s.Usecase.Update(ctx, req.GetId(), product)
	if err != nil {
		return nil, err
	}
	product.ID = req.GetId()
	return &pb.ProductResponse{
		Product: &pb.Product{
			Id:       product.ID,
			Name:     product.Name,
			Category: product.Category,
			Price:    product.Price,
			Stock:    int32(product.Stock),
		},
	}, nil
}

func (s *InventoryGRPCServer) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.Empty, error) {
	err := s.Usecase.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (s *InventoryGRPCServer) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	products, err := s.Usecase.List(ctx, nil, 0, 0)
	if err != nil {
		return nil, err
	}
	var pbProducts []*pb.Product
	for _, p := range products {
		pbProducts = append(pbProducts, &pb.Product{
			Id:       p.ID,
			Name:     p.Name,
			Category: p.Category,
			Price:    p.Price,
			Stock:    int32(p.Stock),
		})
	}
	return &pb.ListProductsResponse{
		Products: pbProducts,
	}, nil
}

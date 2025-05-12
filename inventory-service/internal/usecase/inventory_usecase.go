package usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"inventory-service/internal/entity"
)

type Repository interface {
	Create(context.Context, *entity.Product) error
	GetByID(context.Context, string) (*entity.Product, error)
	Update(context.Context, string, *entity.Product) error
	Delete(context.Context, string) error
	List(context.Context, bson.M, int64, int64) ([]entity.Product, error)
}

type InventoryUsecase struct {
	repo Repository
}

func NewInventoryUsecase(r Repository) *InventoryUsecase {
	return &InventoryUsecase{repo: r}
}

func (u *InventoryUsecase) Create(ctx context.Context, p *entity.Product) error {
	return u.repo.Create(ctx, p)
}

func (u *InventoryUsecase) GetByID(ctx context.Context, id string) (*entity.Product, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *InventoryUsecase) Update(ctx context.Context, id string, p *entity.Product) error {
	return u.repo.Update(ctx, id, p)
}

func (u *InventoryUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}

func (u *InventoryUsecase) List(ctx context.Context, filter bson.M, limit int64, skip int64) ([]entity.Product, error) {
	return u.repo.List(ctx, filter, limit, skip)
}

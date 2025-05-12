package usecase

import (
	"context"
	"order-service/internal/entity"
)

type OrderRepository interface {
	Create(ctx context.Context, o *entity.Order) error
	UpdateStatus(ctx context.Context, id string, status string) error
	GetByID(ctx context.Context, id string) (*entity.Order, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*entity.Order, error)
}

type OrderUsecase struct {
	repo OrderRepository
}

func NewOrderUsecase(r OrderRepository) *OrderUsecase {
	return &OrderUsecase{repo: r}
}

func (u *OrderUsecase) Create(ctx context.Context, o *entity.Order) error {
	o.Status = "created"
	return u.repo.Create(ctx, o)
}

func (u *OrderUsecase) UpdateStatus(ctx context.Context, id string, status string) error {
	return u.repo.UpdateStatus(ctx, id, status)
}

func (u *OrderUsecase) GetByID(ctx context.Context, id string) (*entity.Order, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *OrderUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}

func (u *OrderUsecase) List(ctx context.Context) ([]*entity.Order, error) {
	return u.repo.List(ctx)
}

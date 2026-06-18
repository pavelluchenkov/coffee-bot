package service

import (
	"context"
	"coffee-bot/internal/models"
	"coffee-bot/internal/repository"
)

type OrderService struct {
	repo *repository.Repository
}

func NewOrderService(repo *repository.Repository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) StartOrder(ctx context.Context, userID int64) (int, error) {
	return s.repo.CreateOrder(ctx, userID)
}
func (s *OrderService) AddItem(ctx context.Context, orderID int, item models.MenuItem) error {
	return s.repo.AddOrderItem(ctx, models.OrderItem{
		OrderID:    orderID,
		MenuItemID: item.ID,
		Name:       item.Name,
		Size:       item.Size,
		Price:      item.Price,
	})
}
func (s *OrderService) GetTotal(ctx context.Context, orderID int) (int, error) {
	return s.repo.GetOrderTotal(ctx, orderID)
}
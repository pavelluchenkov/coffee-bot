package services

import (
	"coffee-bot/internal/models"
	"coffee-bot/internal/repository"
	"context"
)

type MenuService struct{
	repo *repository.Repository
}
func NewMenuService(repo *repository.Repository) *MenuService{
	return &MenuService{repo: repo}
}

func (m *MenuService) GetCategories(ctx context.Context) ([]string, error){
	return m.repo.GetCategories(ctx)
}
func (m *MenuService) GetByCategory(ctx context.Context, category string)([]models.MenuItem, error){
	return m.repo.GetMenuByCategory(ctx, category)
}
func (m *MenuService) GetByName(ctx context.Context, name string)([]models.MenuItem, error){
	return m.repo.GetMenuByName(ctx, name)
}
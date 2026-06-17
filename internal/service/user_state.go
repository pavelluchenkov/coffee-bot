package service

import (
	"coffee-bot/internal/models"
	"coffee-bot/internal/repository"
	"context"
)

type UserStateService struct{
	repo *repository.Repository
}
func NewUserStateService(repo *repository.Repository) *UserStateService{
	return &UserStateService{repo: repo}
}

func (u *UserStateService) Get (ctx context.Context, userID int64) (*models.UserState, error){
	return u.repo.GetUserState(ctx, userID)
}
func (u *UserStateService) Save (ctx context.Context, state models.UserState) error{
	return u.repo.SaveUserState(ctx, state)
}
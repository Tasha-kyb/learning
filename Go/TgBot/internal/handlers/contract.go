package handlers

import (
	"context"

	"github.com/internal/model"
)

type ProfileUseCase interface {
	CreateProfile(ctx context.Context, req model.Profile) (string, error)
	AddCategory(ctx context.Context, req model.Category) (string, error)
	GetAllCategories(ctx context.Context, userID int64) ([]model.Category, error)
	DeleteCategory(ctx context.Context, userID int64, id int) (string, error)
}

package usecase

import (
	"context"

	"github.com/internal/model"
)

type ProfileRepository interface {
	CreateProfile(ctx context.Context, req *model.Profile) error
	AddCategory(ctx context.Context, req *model.Category) (int, error)
	GetAllCategories(ctx context.Context, userID int64) ([]model.Category, error)
	DeleteCategory(ctx context.Context, userID int64, id int) (string, error)
	AddExpense(ctx context.Context, expence *model.Expense) (*model.Expense, error)
}

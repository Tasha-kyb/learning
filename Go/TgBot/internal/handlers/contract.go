package handlers

import (
	"context"

	"github.com/internal/model"
)

type ProfileUseCase interface {
	CreateProfile(ctx context.Context, req model.Profile) (string, error)
	AddCategory(ctx context.Context, req model.Category) (string, error)
	GetAllCategories(ctx context.Context, userID int64) (string, error)
	DeleteCategory(ctx context.Context, userID int64, id int) (string, error)
	AddExpense(ctx context.Context, expence *model.Expense) (string, error)
	TodayExpense(ctx context.Context, userID int64) (string, error)
	WeekExpense(ctx context.Context, userID int64) (string, error)
}

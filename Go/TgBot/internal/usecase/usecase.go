package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/internal/model"
)

type ProfileServiceT struct {
	repository ProfileRepository
}

func NewProfileService(repository ProfileRepository) *ProfileServiceT {
	return &ProfileServiceT{repository: repository}
}
func (p *ProfileServiceT) CreateProfile(ctx context.Context, req model.Profile) (string, error) {
	if req.ID == 0 || strings.TrimSpace(req.Username) == "" {
		return "", errors.New("❌ Не хватает параметров для создания профиля")
	}
	newProfile := &model.Profile{
		ID:         req.ID,
		Username:   req.Username,
		Created_at: time.Now(),
	}
	err := p.repository.Create(ctx, newProfile)
	if err != nil {
		return "", fmt.Errorf("❌ Ошибка при создании профиля, %w", err)
	}
	startMassage := `
	👋 Добро пожаловать в Expense Tracker!

	Я помогу вам отслеживать расходы и управлять бюджетами.

	✅ Вы зарегистрированы!
	📂 Созданы базовые категории:
   • Еда
   • Транспорт
   • Развлечения
   • Прочее
`
	return startMassage, nil
}

func (p *ProfileServiceT) AddCategory(ctx context.Context, req model.Category) (string, error) {
	if strings.TrimSpace(req.Name) == "" {
		return "", errors.New("❌ Не хватает параметров для создания категории")
	}
	newCategory := &model.Category{
		UserID: req.UserID,
		Name:   req.Name,
		Color:  req.Color,
	}
	id, err := p.repository.AddCategory(ctx, newCategory)
	if err != nil {
		if strings.Contains(err.Error(), "уже существует") {
			return "", fmt.Errorf("❌ Категория %s уже существует", req.Name)
		}
		return "", fmt.Errorf("❌ Ошибка при создании категории, %w", err)
	}
	addCategoryMassage := fmt.Sprintf(`
	✅ Категория создана!
	📂 Название: %s
	🎨 Цвет: %s
	🆔 ID: %d
	Используйте этот ID для удаления категории.
	`, req.Name, req.Color, id)

	return addCategoryMassage, nil
}

func (p *ProfileServiceT) GetAllCategories(ctx context.Context, userID int64) ([]model.Category, error) {
	categoriesDB, err := p.repository.GetAllCategories(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("❌ Ошибка при получении категорий: %w", err)
	}
	if categoriesDB == nil {
		return []model.Category{}, nil
	}
	var allCategories []model.Category
	for _, categoryDB := range categoriesDB {
		category := model.Category{
			Name:  categoryDB.Name,
			Color: categoryDB.Color,
			ID:    categoryDB.ID,
		}
		allCategories = append(allCategories, category)
	}
	if allCategories == nil {
		allCategories = []model.Category{}
	}
	return allCategories, nil
}
func (p *ProfileServiceT) DeleteCategory(ctx context.Context, userID int64, id int) (string, error) {
	if id == 0 {
		return "", errors.New("❌ Ошибка: некорректно указан id категории: %")
	}
	categoryName, err := p.repository.DeleteCategory(ctx, userID, id)
	if err != nil {
		return "", fmt.Errorf("❌ Ошибка при удалении категории: %w", err)
	}
	deleteCategoryMassage := fmt.Sprintf(`
	✅ Категория %s удалена
	Все расходы из этой категории перенесены в "Прочее"
	`, categoryName)
	return deleteCategoryMassage, nil
}

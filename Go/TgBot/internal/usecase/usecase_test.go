package usecase

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/internal/model"
)

type MockRepository struct {
	CreateProfileFunc    func(ctx context.Context, profile *model.Profile) error
	AddCategoryFunc      func(ctx context.Context, category *model.Category) (int, error)
	GetAllCategoriesFunc func(ctx context.Context, userID int64) ([]model.Category, error)
	DeleteCategoryFunc   func(ctx context.Context, userID int64, id int) (string, error)
	AddExpenseFunc       func(ctx context.Context, expense *model.Expense) (*model.Expense, error)
}

func (m MockRepository) CreateProfile(ctx context.Context, profile *model.Profile) error {
	if m.CreateProfileFunc != nil {
		return m.CreateProfileFunc(ctx, profile)
	}
	return nil
}
func (m MockRepository) AddCategory(ctx context.Context, category *model.Category) (int, error) {
	if m.AddCategoryFunc != nil {
		return m.AddCategoryFunc(ctx, category)
	}
	return 0, nil
}
func (m MockRepository) GetAllCategories(ctx context.Context, userID int64) ([]model.Category, error) {
	if m.GetAllCategoriesFunc != nil {
		return m.GetAllCategoriesFunc(ctx, userID)
	}
	return nil, nil
}
func (m MockRepository) DeleteCategory(ctx context.Context, userID int64, id int) (string, error) {
	if m.DeleteCategoryFunc != nil {
		return m.DeleteCategoryFunc(ctx, userID, id)
	}
	return "", nil
}
func (m MockRepository) AddExpense(ctx context.Context, expense *model.Expense) (*model.Expense, error) {
	if m.AddExpenseFunc != nil {
		return m.AddExpenseFunc(ctx, expense)
	}
	return &model.Expense{}, nil
}

func TestCreateProfile(t *testing.T) {
	tests := []struct {
		name        string
		input       model.Profile
		mockFunc    func(ctx context.Context, profile *model.Profile) error
		wantError   bool
		wantMessage string
	}{
		{
			name:  "Успешное создание профиля",
			input: model.Profile{ID: 123456, Username: "user"},
			mockFunc: func(ctx context.Context, profile *model.Profile) error {
				return nil
			},
			wantError:   false,
			wantMessage: "👋 Добро пожаловать",
		},
		{
			name:  "Ошибка: ID = 0",
			input: model.Profile{ID: 0, Username: "user"},
			mockFunc: func(ctx context.Context, profile *model.Profile) error {
				return nil
			},
			wantError: true,
		},
		{
			name:  "Ошибка: пустое имя",
			input: model.Profile{ID: 123456, Username: ""},
			mockFunc: func(ctx context.Context, profile *model.Profile) error {
				return nil
			},
			wantError: true,
		},
		{
			name:  "Ошибка в репозитории",
			input: model.Profile{ID: 123456, Username: "user"},
			mockFunc: func(ctx context.Context, profile *model.Profile) error {
				return errors.New("Ошибка БД")
			},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockRepository{
				CreateProfileFunc: tt.mockFunc,
			}
			service := NewProfileService(mockRepo)
			message, err := service.CreateProfile(context.Background(), tt.input)
			if !tt.wantError && err != nil {
				t.Error("Ошибка не ожидалась, но ее получили")

			}
			if tt.wantError && err == nil {
				t.Error("Ожидалась ошибка, но ее нет")
			}
			if !tt.wantError && !strings.Contains(message, tt.wantMessage) {
				t.Error("Ожидалась сообщение об успешном добавлении категории, но его нет")
			}
		})
	}

	t.Log("Тест завершен")
}
func TestAddCategory(t *testing.T) {
	tests := []struct {
		name        string
		input       model.Category
		mockFunc    func(ctx context.Context, category *model.Category) (int, error)
		wantError   bool
		wantMessage string
	}{
		{
			name:  "Успешное создание категории",
			input: model.Category{ID: 123456, Name: "Спорт"},
			mockFunc: func(ctx context.Context, category *model.Category) (int, error) {
				return 123456, nil
			},
			wantError:   false,
			wantMessage: "✅ Категория создана!",
		},
		{
			name:  "Ошибка: ID = 0",
			input: model.Category{ID: 0, Name: "Спорт"},
			mockFunc: func(ctx context.Context, category *model.Category) (int, error) {
				return 0, nil
			},
			wantError: true,
		},
		{
			name:  "Ошибка: нет названия категории",
			input: model.Category{ID: 123456, Name: ""},
			mockFunc: func(ctx context.Context, category *model.Category) (int, error) {
				return 0, nil
			},
			wantError: true,
		},
		{
			name:  "Ошибка в репозитории",
			input: model.Category{ID: 123456, Name: "Спорт"},
			mockFunc: func(ctx context.Context, category *model.Category) (int, error) {
				return 0, errors.New("Ошибка БД")
			},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockRepository{
				AddCategoryFunc: tt.mockFunc,
			}
			service := NewProfileService(mockRepo)
			message, err := service.AddCategory(context.Background(), tt.input)
			if !tt.wantError && err != nil {
				t.Error("Ошибка не ожидалась, но ее получили")

			}
			if tt.wantError && err == nil {
				t.Error("Ожидалась ошибка, но ее нет")
			}
			if !tt.wantError && !strings.Contains(message, tt.wantMessage) {
				t.Error("Ожидалась сообщение об успешной регистрации, но его нет")
			}
		})
	}

	t.Log("Тест завершен")
}
func TestGetAllCategories(t *testing.T) {
	tests := []struct {
		name        string
		userID      int64
		mockFunc    func(ctx context.Context, userID int64) ([]model.Category, error)
		wantError   bool
		wantMessage string
	}{
		{
			name:   "Успешное получение категорий",
			userID: 123,
			mockFunc: func(ctx context.Context, userID int64) ([]model.Category, error) {
				return []model.Category{
					{ID: 123, Name: "Красота", Color: "синий"},
					{ID: 65422432, Name: "Спорт", Color: ""},
				}, nil
			},
			wantError:   false,
			wantMessage: "📂 Ваши категории:",
		},
		{
			name:   "Ошибка в репозитории",
			userID: 123,
			mockFunc: func(tx context.Context, userID int64) ([]model.Category, error) {
				return nil, errors.New("Ошибка БД")
			},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockRepository{
				GetAllCategoriesFunc: tt.mockFunc,
			}
			service := NewProfileService(mockRepo)
			_, err := service.GetAllCategories(context.Background(), tt.userID)
			if !tt.wantError && err != nil {
				t.Error("Ошибка не ожидалась, но ее получили")

			}
			if tt.wantError && err == nil {
				t.Error("Ожидалась ошибка, но ее нет")
			}
		})
	}

	t.Log("Тест завершен")
}
func TestDeleteCategory(t *testing.T) {
	tests := []struct {
		name        string
		userID      int64
		id          int
		mockFunc    func(ctx context.Context, userID int64, id int) (string, error)
		wantError   bool
		wantMessage string
	}{
		{
			name:   "Успешное удаление категории",
			userID: 123,
			id:     123,
			mockFunc: func(ctx context.Context, userID int64, is int) (string, error) {
				return "Спорт", nil
			},
			wantError:   false,
			wantMessage: "✅ Категория",
		},
		{
			name:   "Некорректно указан id категории",
			userID: 123,
			id:     0,
			mockFunc: func(ctx context.Context, userID int64, is int) (string, error) {
				return "Спорт", nil
			},
			wantError: true,
		},
		{
			name:   "Ошибка в репозитории",
			userID: 123,
			id:     123,
			mockFunc: func(ctx context.Context, userID int64, is int) (string, error) {
				return "Спорт", errors.New("Ошибка БД")
			},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockRepository{
				DeleteCategoryFunc: tt.mockFunc,
			}
			service := NewProfileService(mockRepo)
			message, err := service.DeleteCategory(context.Background(), tt.userID, tt.id)
			if !tt.wantError && err != nil {
				t.Error("Ошибка не ожидалась, но ее получили")

			}
			if tt.wantError && err == nil {
				t.Error("Ожидалась ошибка, но ее нет")
			}
			if !tt.wantError && !strings.Contains(message, tt.wantMessage) {
				t.Error("Ожидалась сообщение об успешной регистрации, но его нет")
			}
		})
	}

	t.Log("Тест завершен")
}
func TestAddExpense(t *testing.T) {
	tests := []struct {
		name        string
		input       model.Expense
		mockFunc    func(ctx context.Context, expense *model.Expense) (*model.Expense, error)
		wantError   bool
		wantMessage string
	}{
		{
			name:  "Успешное создание расхода",
			input: model.Expense{UserID: 1, Amount: 123, Category: "Транспорт", Description: "Поездка в трамвае"},
			mockFunc: func(ctx context.Context, expense *model.Expense) (*model.Expense, error) {
				return expense, nil
			},
			wantError:   false,
			wantMessage: "✅ Расход добавлен!",
		},
		{
			name:  "Расход отрицательный",
			input: model.Expense{UserID: 1, Amount: -123, Category: "Категория", Description: "Поездка в трамвае"},
			mockFunc: func(ctx context.Context, expense *model.Expense) (*model.Expense, error) {
				return nil, errors.New("Сумма расхода должна быть положительной")
			},
			wantError: true,
		},
		{
			name:  "Расход нулевой",
			input: model.Expense{UserID: 1, Amount: 0, Category: "Категория", Description: "Поездка в трамвае"},
			mockFunc: func(ctx context.Context, expense *model.Expense) (*model.Expense, error) {
				return nil, errors.New("Не хватает данных для добавления расхода: расход равен нулю")
			},
			wantError: true,
		},
		{
			name:  "Не указана категория",
			input: model.Expense{UserID: 1, Amount: 123, Category: "", Description: "Поездка в трамвае"},
			mockFunc: func(ctx context.Context, expense *model.Expense) (*model.Expense, error) {
				return nil, errors.New("Не хватает данных для добавления расхода: не указана категория")
			},
			wantError: true,
		},
		{
			name:  "Не указано описание",
			input: model.Expense{UserID: 1, Amount: 123, Category: "Категория", Description: ""},
			mockFunc: func(ctx context.Context, expense *model.Expense) (*model.Expense, error) {
				return nil, errors.New("Не хватает данных для добавления расхода: не указано описание")
			},
			wantError: true,
		},
		{
			name:  "Категория не найдена в БД",
			input: model.Expense{UserID: 1, Amount: 123, Category: "Космос", Description: "Поездка в трамвае"},
			mockFunc: func(ctx context.Context, expense *model.Expense) (*model.Expense, error) {
				return nil, errors.New("Указанная категория не найдена в базе данных")
			},
			wantError: true,
		},
		{
			name:  "Не хватает описания расхода",
			input: model.Expense{UserID: 1, Amount: 123, Category: "Космос", Description: ""},
			mockFunc: func(ctx context.Context, expense *model.Expense) (*model.Expense, error) {
				return nil, errors.New("Не хватает описания расхода")
			},
			wantError: true,
		},
		{
			name:  "Ошибка в репозитории",
			input: model.Expense{UserID: 1, Amount: 123, Category: "Космос", Description: "Поездка в трамвае"},
			mockFunc: func(ctx context.Context, expense *model.Expense) (*model.Expense, error) {
				return nil, errors.New("Ошибка БД")
			},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockRepository{
				AddExpenseFunc: tt.mockFunc,
			}
			service := NewProfileService(mockRepo)
			message, err := service.AddExpense(context.Background(), &tt.input)
			if !tt.wantError && err != nil {
				t.Error("Ошибка не ожидалась, но ее получили")

			}
			if tt.wantError && err == nil {
				t.Error("Ожидалась ошибка, но ее нет")
			}
			if !tt.wantError && !strings.Contains(message, tt.wantMessage) {
				t.Error("Ожидалась сообщение об успешном добавлении категории, но его нет")
			}
		})
	}

	t.Log("Тест завершен")
}

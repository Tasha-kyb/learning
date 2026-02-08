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
	err := p.repository.CreateProfile(ctx, newProfile)
	if err != nil {
		return "", fmt.Errorf("❌ Ошибка при создании профиля, %w", err)
	}
	startMessage := `
	👋 Добро пожаловать в Expense Tracker!

	Я помогу вам отслеживать расходы и управлять бюджетами.

	✅ Вы зарегистрированы!
	📂 Созданы базовые категории:
   • Еда
   • Транспорт
   • Развлечения
   • Прочее
`
	return startMessage, nil
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
	addCategoryMessage := fmt.Sprintf(`
	✅ Категория создана!
	📂 Название: %s
	🎨 Цвет: %s
	🆔 ID: %d
	Используйте этот ID для удаления категории.
	`, req.Name, req.Color, id)

	return addCategoryMessage, nil
}
func (p *ProfileServiceT) GetAllCategories(ctx context.Context, userID int64) (string, error) {
	categoriesDB, err := p.repository.GetAllCategories(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("❌ Ошибка при получении категорий: %w", err)
	}
	if len(categoriesDB) == 0 {
		return "У вас пока нет категорий. \nСоздать категорию можно командой /category add", nil
	}
	response := "📂 Ваши категории:\n\n"
	for _, category := range categoriesDB {
		response += fmt.Sprintf("%s\n\n", category.Name)
		if category.Color != "" {
			response += fmt.Sprintf("%s\n\n", category.Color)
		}
		response += fmt.Sprintf("ID: %d\n", category.ID)
	}
	response += "\n💡 Используйте ID для удаления категории"
	return response, nil
}
func (p *ProfileServiceT) DeleteCategory(ctx context.Context, userID int64, id int) (string, error) {
	if id <= 0 {
		return "", errors.New("❌ Ошибка: некорректно указан id категории")
	}
	categoryName, err := p.repository.DeleteCategory(ctx, userID, id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return "", fmt.Errorf("❌ Ошибка: некорректно указан ID категории")
		}
		return "", fmt.Errorf("❌ Ошибка при удалении категории: %w", err)
	}
	deleteCategoryMassage := fmt.Sprintf(`
	✅ Категория %s удалена
	Все расходы из этой категории перенесены в "Прочее"
	`, categoryName)
	return deleteCategoryMassage, nil
}
func (p *ProfileServiceT) AddExpense(ctx context.Context, req *model.Expense) (string, error) {
	if req.Amount <= 0 {
		return "", errors.New("❌ Сумма расхода должна быть положительной")
	}
	if req.Category == "" || req.Description == "" {
		return "", errors.New("❌ Не хватает данных для добавления расхода")
	}
	if req.Created_at.IsZero() {
		req.Created_at = time.Now()
	}
	newExpense := &model.Expense{
		UserID:      req.UserID,
		Amount:      req.Amount,
		Category:    req.Category,
		Description: req.Description,
		Created_at:  req.Created_at,
	}
	expense, err := p.repository.AddExpense(ctx, newExpense)
	if err != nil {
		if strings.Contains(err.Error(), "не найдена в базе данных") {
			return "", fmt.Errorf("❌ Категория \"%s\" не найдена", req.Category)
		}
		return "", fmt.Errorf("❌ Ошибка при создании расхода %w", err)
	}
	addExpenseMessage := fmt.Sprintf(`
	✅ Расход добавлен!

	💰 Сумма: %.2f₽
	📂 Категория: %s
	📝 Описание: %s
	📅 Дата: %s

	💵 Осталось до лимита: X
	`, expense.Amount, expense.Category, expense.Description, expense.Created_at.Format("02.01.2006"))

	return addExpenseMessage, nil
}
func (p *ProfileServiceT) TodayExpense(ctx context.Context, userID int64) (string, error) {
	expenses, err := p.repository.TodayExpense(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("❌ Ошибка при при получении расходов за сегодня %w", err)
	}
	today := time.Now().Format("02.01.2006")
	if len(expenses) == 0 {
		return fmt.Sprintf(`📊 Расходы за сегодня (%s)
		
		Пока нет расходов за сегодня.
		Используйте /add для добавления расхода.`, today), nil
	}
	categoriesMap := make(map[string][]model.Expense)

	for _, expense := range expenses {
		categoriesMap[expense.Category] = append(categoriesMap[expense.Category], expense)
	}

	response := fmt.Sprintf("📊 Расходы за сегодня (%s)\n\n", today)
	total := 0.0

	for category, expenseList := range categoriesMap {
		sum := 0.0
		for _, exp := range expenseList {
			sum += exp.Amount
		}
		response += fmt.Sprintf("%s: %.2f₽\n", category, sum)

		for _, exp := range expenseList {
			response += fmt.Sprintf("   • %s: %.2f₽\n", exp.Description, exp.Amount)
		}
		total += sum
	}
	response += "\n━━━━━━━━━━━━━━━━━━━━\n"
	response += fmt.Sprintf("💰 Итого: %.2f₽", total)

	return response, nil
}
func (p *ProfileServiceT) WeekExpense(ctx context.Context, userID int64) (string, error) {
	expenses, err := p.repository.WeekExpense(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("❌ Ошибка при при получении расходов за неделю %w", err)
	}

	if len(expenses) == 0 {
		return "📊 Нет расходов за неделю", nil
	}

	total := 0.0
	for _, exp := range expenses {
		total += exp.Amount
	}
	response := fmt.Sprintf(`📊 Расходы за неделю 
	💰 Итого: %.2f₽
	📈 Средний расход в день: %.2f₽`, total, total/7)

	return response, nil
}

package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

/*type ExpenseRepositoryT struct {
	pool *pgxpool.Pool
}*/

type ProfileRepositoryT struct {
	pool *pgxpool.Pool
}

func NewProfileRepo(pool *pgxpool.Pool) *ProfileRepositoryT {
	return &ProfileRepositoryT{pool: pool}
}

func (p *ProfileRepositoryT) Create(ctx context.Context, profile *model.Profile) error {
	tr, err := p.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("Не удалось начать транзакцию, %w", err)
	}
	defer tr.Rollback(ctx)

	userQuery := `
	INSERT INTO users (user_id, username, created_at)
	VALUES ($1, $2, $3)
	ON CONFLICT (user_id) DO UPDATE
	SET username = EXCLUDED.username
	`
	_, err = tr.Exec(ctx, userQuery, profile.ID, profile.Username, profile.Created_at)
	if err != nil {
		log.Printf("Ошибка при создании пользователя: %v", err)
		return fmt.Errorf("Ошибка при создании пользователя, %w", err)
	}
	categoryQuery := `
	INSERT INTO categories (user_id, name) VALUES 
	($1, 'Еда'), 
	($1, 'Транспорт'), 
	($1, 'Развлечения'), 
	($1, 'Прочее')
	ON CONFLICT (user_id, name) DO NOTHING
	`
	_, err = tr.Exec(ctx, categoryQuery, profile.ID)
	if err != nil {
		log.Printf("Ошибка при создании категорий: %v", err)
		return fmt.Errorf("Ошибка при создании категорий, %w", err)
	}

	if err = tr.Commit(ctx); err != nil {
		log.Printf("Ошибка при завершении транзакции: %v", err)
		return fmt.Errorf("Ошибка при завершении транзакции: %w", err)
	}

	return nil
}

/*func (r *AddExpenseT) AddCategory(userID int) (domain.ExpenseServiceT, error) {
	if userID
}*/

/*func (r ExpenseRepositoryT) CreateExpense(expense domain.ExpenseT) (domain.ExpenseT, error) {

	query := `
	SELECT user_id FROM users WHERE user_id = $1
	`
	if query
	err := r.pool.QueryRow(query, id).Scan(
		&expense.UserID,
	)

	if err != nil {
		return domain.ExpenseT{}, fmt.Errorf("create list: %w", err)
	}

	return expense, nil
}*/

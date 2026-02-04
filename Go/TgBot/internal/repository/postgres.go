package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

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

func (p *ProfileRepositoryT) AddCategory(ctx context.Context, category *model.Category) (int, error) {
	query := `
	INSERT INTO categories (user_id, name, color) 
	VALUES ($1, $2, $3)
	ON CONFLICT (user_id, name) DO NOTHING
	RETURNING id`
	var id int
	err := p.pool.QueryRow(ctx, query, category.UserID, category.Name, category.Color).Scan(&id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, fmt.Errorf("Категория %s уже существует", category.Name)
		}
		log.Printf("Ошибка при создании категории: %v", err)
		return 0, fmt.Errorf("Ошибка при создании категории: %w", err)
	}
	return id, nil
}
func (p *ProfileRepositoryT) GetAllCategories(ctx context.Context, userID int64) ([]model.Category, error) {
	query := `
	SELECT id, name, color
	FROM categories WHERE user_id = $1
	ORDER by user_id
	`
	rows, err := p.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Ошибка запроса категорий из базы данных, %w", err)
	}
	defer rows.Close()

	var allCategories []model.Category
	for rows.Next() {
		var category model.Category
		err := rows.Scan(
			&category.Name,
			&category.Color,
			&category.ID,
		)
		if err != nil {
			return nil, fmt.Errorf("Ошибка при получении списка категорий, %w", err)
		}
		allCategories = append(allCategories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Ошибка при чтении категорий, %w", err)
	}
	if allCategories == nil {
		allCategories = []model.Category{}
	}
	return allCategories, nil
}
func (p *ProfileRepositoryT) DeleteCategory(ctx context.Context, userID int64, id int) (string, error) {
	query := `
	DELETE FROM categories WHERE user_id = $1 AND id = $2
	RETURNING name`
	var name string
	err := p.pool.QueryRow(ctx, query, userID, id).Scan(&name)
	if err != nil {
		log.Printf("Ошибка при удалении категории: %v", err)
		return "", fmt.Errorf("Ошибка при удалении категории: %w", err)
	}
	return name, nil
}

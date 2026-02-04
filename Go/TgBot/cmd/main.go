package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/internal/handlers"
	"github.com/internal/repository"
	"github.com/internal/repository/database"
	"github.com/internal/usecase"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := database.NewPool(ctx)
	if err != nil {
		log.Fatalf("Не удалось подключиться к БД: %v", err)
	}
	defer pool.Close()
	fmt.Println("Подключение к БД")

	profileRepo := repository.NewProfileRepo(pool)
	profileService := usecase.NewProfileService(profileRepo)

	// REST API обработчик
	httpHandler := handlers.NewProfileHandler(profileService)

	// TG обработчик
	tgHandler := handlers.NewTelegramUpdates(profileService)

	// запуск ТГ бота в фоне
	go tgHandler.StartUpdates()

	router := handlers.NewRouter(httpHandler)

	log.Println("ТГ бот запущен ")
	log.Println("HTTP сервер запущен на :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Ошибка сервера, %v", err)
	}

	log.Println("Бот остановлен")
}

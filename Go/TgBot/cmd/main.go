package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/internal/handlers"
	"github.com/internal/repository"
	"github.com/internal/repository/database"
	"github.com/internal/service"
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
	profileService := service.NewProfileService(profileRepo)
	httpServer := handlers.NewProfileHandler(profileService)
	router := http.NewServeMux()

	router.HandleFunc("/start", httpServer.Create)
	http.ListenAndServe(":8080", router)
	/*bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatal("Не получилось подключиться к Телеграмм")
		panic("Не получилось подключиться к Телеграмм")
	}
	router.HandleFunc("/start", httpServer.Create)
	http.ListenAndServe(":8080", router)*/

	log.Println("Бот остановлен")
}

package handlers

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/internal/model"
	"github.com/internal/usecase"
)

type TelegramHandlerT struct {
	usecase *usecase.ProfileServiceT
}

func NewTelegramUpdates(usecase *usecase.ProfileServiceT) *TelegramHandlerT {
	return &TelegramHandlerT{usecase: usecase}
}

func (t *TelegramHandlerT) StartUpdates() {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Бот запущен")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		switch {
		case update.Message == nil:
			continue
		case update.Message.Text == "/start":
			profile := model.Profile{
				ID:         int64(update.Message.From.ID),
				Username:   update.Message.From.UserName,
				Created_at: time.Now(),
			}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			response, err := t.usecase.CreateProfile(ctx, profile)

			if err != nil {
				log.Printf("Ошибка создания профиля, %v", err)
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "❌Ошибка создания профиля"))
				continue
			}
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, response))
		case strings.HasPrefix(update.Message.Text, "/category add"):
			parts := strings.Fields(update.Message.Text)
			if len(parts) < 3 {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "❌Ошибка: вы не указали название категории"))
				continue
			}
			categoryName := parts[2]
			color := ""
			if len(parts) >= 4 {
				color = parts[3]
			}
			newCategory := model.Category{
				UserID: int64(update.Message.From.ID),
				Name:   categoryName,
				Color:  color,
			}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			response, err := t.usecase.AddCategory(ctx, newCategory)

			if err != nil {
				log.Printf("Ошибка создания категории, %v", err)
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
				continue
			}
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, response))
		case update.Message.Text == "/categories":
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			categories, err := t.usecase.GetAllCategories(ctx, update.Message.From.ID)

			if err != nil {
				log.Printf("Ошибка получения категорий, %v", err)
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID,
					"❌Ошибка при получении категорий"))
				continue
			}
			if len(categories) == 0 {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID,
					"❌У вас пока нет категорий. \nСоздать категорию можно командой /category add"))
				continue
			}
			response := "📂 Ваши категории:\n\n"
			for _, category := range categories {
				response += fmt.Sprintf("%s\n\n", category.Name)
				if category.Color != "" {
					response += fmt.Sprintf("%s\n\n", category.Color)
				}
				response += fmt.Sprintf("ID: %d\n", category.ID)
			}
			response += "\n💡 Используйте ID для удаления категории"
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, response))
		case strings.HasPrefix(update.Message.Text, "/category delete"):
			parts := strings.Fields(update.Message.Text)
			if len(parts) < 3 {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID,
					"❌Ошибка: Вы не указали id категории для удаления"))
				continue
			}
			idstr := parts[2]
			id, err := strconv.Atoi(idstr)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID,
					"❌Ошибка: некорректно указан id категории"))
				continue
			}

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			response, err := t.usecase.DeleteCategory(ctx, update.Message.From.ID, id)
			if err != nil {
				log.Printf("❌Ошибка при удалении категории, %v", err)
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
				continue
			}
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, response))
		case update.Message.Text == "/help":
			helpText := `
			📖 Доступные команды:
			💰 Расходы:
			/add <сумма> <категория> <описание> — добавить расход
			/today — расходы за сегодня
			/week — расходы за неделю
			/month — расходы за месяц
			/stats — общая статистика

			📂 Категории:
			/category add <название> <цвет> — создать категорию
			/categories — список категорий
			/category delete <id> — удалить категорию

			💵 Бюджеты:
			/budget <категория> <сумма> <период> — установить лимит
			/budgets — список бюджетов
			/limit <сумма> — общий лимит

			📊 Экспорт:
			/export csv — экспорт в CSV
			/export json — экспорт в JSON
			/report <период> — детальный отчет`
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, helpText))
		case strings.HasPrefix(update.Message.Text, "/add"):
			parts := strings.Fields(update.Message.Text)
			if len(parts) < 4 {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID,
					"❌Ошибка: вы не указали все параметры (сумма, категория и описание)"))
				continue
			}
			amount, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID,
					"❌ Сумма должна быть числом"))
				continue
			}
			category := parts[2]
			discription := strings.Join(parts[3:], " ")
			newExpense := model.Expense{
				UserID:      int64(update.Message.From.ID),
				Amount:      amount,
				Category:    category,
				Description: discription,
			}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			response, err := t.usecase.AddExpense(ctx, &newExpense)

			if err != nil {
				log.Printf("Ошибка создания расхода, %v", err)
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
				continue
			}
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, response))
		default:
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "❌Неизвестная команда, используйте /help"))
		}
	}
}

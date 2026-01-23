package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/internal/model"
	"github.com/internal/repository"
)

type ProfileServiceT struct {
	repository *repository.ProfileRepositoryT
}

func NewProfileService(repository *repository.ProfileRepositoryT) *ProfileServiceT {
	return &ProfileServiceT{repository: repository}
}
func (p *ProfileServiceT) CreateProfile(ctx context.Context, req model.Profile) (string, error) {
	if req.ID == 0 || req.Username == "" {
		return "", errors.New("Не хватает параметров для создания профиля")
	}
	newProfile := &model.Profile{
		ID:         req.ID,
		Username:   req.Username,
		Created_at: time.Now(),
	}
	err := p.repository.Create(ctx, newProfile)
	if err != nil {
		return "", fmt.Errorf("Ошибка при создании профиля, %w", err)
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

/*type ExpenseServiceT struct {
	repo ExpenseRepositoryT
}*/
/*type MessageT struct {
	Text string `json:"text"`
	//Categories Expense `json:"categories"`
}

const TelegramBaseUrl = "https://api.telegram.org/bot"
const TelegramToken = "8364889336:AAFCJOUAGuFmIaPMUAA5Twtorxj-DGzJ_2M"
const MethodGetMe = "getMe"
const MethodGetUpdates = "getUpdates"
const MethodMessage = "sendMessage"

func GetUrlByMethod(methodName string) string {
	return TelegramBaseUrl + TelegramToken + "/" + methodName
}

func GetBodyByUrl(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	return body
}*/

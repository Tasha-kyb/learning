package model

import (
	"time"
)

// модель для получения профиля
type Profile struct {
	ID         int64     `json:"id"`
	Username   string    `json:"username"`
	Created_at time.Time `json:"-"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type StartT struct {
	Text string `json:"text"`
}

type ExpenseT struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Category   int       `json:"category"`
	Amount     float64   `json:"amount"`
	Text       string    `json:"text"`
	Created_at time.Time `json:"created_at"`
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/internal/model"
	"github.com/internal/service"
)

type ProfileHandlerT struct {
	service *service.ProfileServiceT
}

func NewProfileHandler(service *service.ProfileServiceT) *ProfileHandlerT {
	return &ProfileHandlerT{service: service}
}

func (p *ProfileHandlerT) Create(w http.ResponseWriter, r *http.Request) {
	var profile model.Profile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		http.Error(w, `{"error": "Некорректный JSON"}`, http.StatusBadRequest)
		return
	}

	response, err := p.service.CreateProfile(r.Context(), profile)
	if err != nil {
		http.Error(w, `{"error": "Внутренняя ошибка сервера"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&response)
}

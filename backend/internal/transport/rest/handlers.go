package rest

import (
	"backend/internal/models"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type ServiceInterface interface {
	GetStatuses(ctx context.Context) ([]models.PingStatus, error)
	UpdateStatus(ctx context.Context, ping models.PingStatus) error
}

type Handler struct {
	service ServiceInterface
}

func NewHandler(service ServiceInterface) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetStatusesData(w http.ResponseWriter, r *http.Request) {
	var statuses []models.PingStatus
	statuses, err := h.service.GetStatuses(r.Context())
	if err != nil {
		log.Println("Ошибка получения статусов")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewEncoder(w).Encode(statuses); err != nil {
		log.Println("Ошибка ответа")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(statuses)
	log.Println("Выполнен запрос для получения статусов пингов")
}

func (h *Handler) UpdateStatusData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("Метод не разрешён")
		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		return
	}

	var ps models.PingStatus
	if err := json.NewDecoder(r.Body).Decode(&ps); err != nil {
		log.Println("Неверный JSON")
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateStatus(r.Context(), ps); err != nil {
		log.Println("Неверный JSON")
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	log.Println("Выполнен запрос для добавления статуса пинга")
}

package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/exp/slog"

	"github.com/DenisGoldiner/webapp/internal"
)

type EmployeeHandler struct {
	service internal.EmployeeService
}

func NewEmployeeHandler(service internal.EmployeeService) EmployeeHandler {
	return EmployeeHandler{
		service: service,
	}
}

func (h EmployeeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetEmployee(w, r)
	case http.MethodPost:
		h.CreateEmployee(w, r)
	default:
		msg := fmt.Sprintf("method %s is not supported", r.Method)
		slog.Info(msg)
		http.Error(w, msg, http.StatusMethodNotAllowed)
	}
}

func (h EmployeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	slog.Info("Get")

	ctx := r.Context()

	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "id must be a valid integer", http.StatusBadRequest)
		return
	}

	res, err := h.service.GetEmployee(ctx, id)
	// Використовуємо кастомну помилку з Лаби 1
	if errors.Is(err, internal.ErrEmployeeNotFound) {
		slog.Warn("request failed", "error", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		slog.Error("request failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respEmployee := Employee{
		ID:       res.ID,
		Name:     res.Name,
		Position: res.Position.Name,
		Salary:   res.Salary,
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(respEmployee); err != nil {
		slog.Error("failed to encode response", "error", err)
		return
	}
}

func (h EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	slog.Info("Create")

	if r.Body == nil {
		http.Error(w, "Body must not be nil", http.StatusBadRequest)
		return
	}

	var payload CreateEmployeePayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		err = fmt.Errorf("failed to decode the body: %w", err)
		slog.Error("Create request failed", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	slog.Info("CreateEmployeePayload", "payload", payload)

	empID, err := h.service.CreateEmployee(r.Context(), payload.toServiceParams())
	// Валідація бізнес-логіки
	if errors.Is(err, internal.ErrInvalidInput) || errors.Is(err, internal.ErrInvalidSalary) {
		slog.Error("Create request failed", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err != nil {
		slog.Error("Create request failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := CreateEmployeeResponse{
		ID: empID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // HTTP 201 Created
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error("failed to encode response", "error", err)
		return
	}
}

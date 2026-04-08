package rest

import (
	"encoding/json"
	"net/http"

	"github.com/VictoriaBlahodarna/GO_lab1/lab3/internal"
)

type TopPaidHandler struct {
	service internal.EmployeeService
}

func NewTopPaidHandler(service internal.EmployeeService) http.Handler {
	return TopPaidHandler{service: service}
}

func (h TopPaidHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	topPaid := h.service.GetTopPaidEmployees()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(topPaid)
}

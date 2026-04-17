package rest

import (
	"encoding/json"
	"net/http"

	"golang.org/x/exp/slog"

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

	serializableTopPaid := make(map[string]internal.Employee)
	for pos, emp := range topPaid {
		serializableTopPaid[pos.Name] = emp
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(serializableTopPaid); err != nil {
		slog.Error("failed to encode top paid response", "error", err)
	}
}

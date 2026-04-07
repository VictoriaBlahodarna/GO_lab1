package rest

import (
	"strings"

	"github.com/DenisGoldiner/webapp/internal"
)

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Salary   uint   `json:"salary"`
}

// CreateEmployeePayload відповідає за парсинг JSON запитів на створення працівника
type CreateEmployeePayload struct {
	Name     string `json:"name"`
	Position string `json:"position"`
	Salary   uint   `json:"salary"`
}

func (p CreateEmployeePayload) toServiceParams() internal.CreateEmployeePayload {
	return internal.CreateEmployeePayload{
		Name:     strings.TrimSpace(p.Name),
		Position: strings.TrimSpace(p.Position),
		Salary:   p.Salary,
	}
}

type CreateEmployeeResponse struct {
	ID int `json:"id"`
}

package internal

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrEmployeeNotFound = errors.New("employee not found")
	ErrInvalidInput     = errors.New("invalid input")
	ErrInvalidSalary    = errors.New("employee salary is out of position range")
)

type EmployeeStorage interface {
	Get(ctx context.Context, id int) (Employee, error)
	Create(ctx context.Context, params CreateEmployeePayload) (int, error)
}

type Position struct {
	Name      string
	MinSalary uint
	MaxSalary uint
}

func (p Position) GetLevel() string {
	switch {
	case p.MinSalary >= 3000:
		return "Senior/Management"
	case p.MinSalary >= 1500:
		return "Middle"
	default:
		return "Junior"
	}
}

type Employee struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Position Position `json:"position"`
	Salary   uint     `json:"salary"`
}

type CreateEmployeePayload struct {
	Name     string `json:"name"`
	Position string `json:"position"`
	Salary   uint   `json:"salary"`
}

func (e Employee) IsSalaryValid() bool {
	return e.Salary >= e.Position.MinSalary && e.Salary <= e.Position.MaxSalary
}

func (e Employee) String() string {
	return fmt.Sprintf("%d: %s (%s), $%d\n", e.ID, e.Name, e.Position.GetLevel(), e.Salary)
}

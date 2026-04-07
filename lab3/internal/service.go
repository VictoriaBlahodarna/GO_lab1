package internal

import (
	"context"
	"fmt"
)

type EmployeeService struct {
	storage EmployeeStorage
}

func NewEmployeeService(storage EmployeeStorage) EmployeeService {
	return EmployeeService{storage: storage}
}

func (s EmployeeService) GetEmployee(ctx context.Context, id int) (Employee, error) {
	if id <= 0 {
		return Employee{}, fmt.Errorf("%w: id must be greater than 0", ErrInvalidInput)
	}

	res, err := s.storage.Get(ctx, id)
	if err != nil {
		return Employee{}, fmt.Errorf("%w: failed to get employee from storage", err)
	}

	return res, nil
}

func (s EmployeeService) CreateEmployee(ctx context.Context, params CreateEmployeePayload) (int, error) {
	if params.Name == "" {
		return 0, fmt.Errorf("%w: company or employee name cannot be empty", ErrInvalidInput)
	}

	if params.Salary == 0 {
		return 0, fmt.Errorf("%w: employee salary is invalid", ErrInvalidSalary)
	}

	if params.Position == "" {
		return 0, fmt.Errorf("%w: position must be provided", ErrInvalidInput)
	}

	empID, err := s.storage.Create(ctx, params)
	if err != nil {
		return 0, fmt.Errorf("failed to create employee in storage: %w", err)
	}

	return empID, nil
}

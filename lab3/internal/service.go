package internal

import (
	"context"
	"fmt"
)

type EmployeeService struct {
	employeesStorage EmployeeStorage
}

func NewEmployeeService(storage EmployeeStorage) EmployeeService {
	return EmployeeService{employeesStorage: storage}
}

func (s EmployeeService) GetEmployee(ctx context.Context, id int) (Employee, error) {
	if id <= 0 {
		return Employee{}, fmt.Errorf("%w: id must be greater than 0", ErrInvalidInput)
	}

	res, err := s.employeesStorage.Get(ctx, id)
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

	empID, err := s.employeesStorage.Create(ctx, params)
	if err != nil {
		return 0, fmt.Errorf("failed to create employee in storage: %w", err)
	}

	return empID, nil
}

func (s EmployeeService) BulkCreateEmployees(ctx context.Context, params []CreateEmployeePayload) error {
	if len(params) == 0 {
		return nil
	}
	return s.employeesStorage.BulkCreate(ctx, params)
}

func (s EmployeeService) GetTopPaidEmployees() map[Position]Employee {
	topPaid := make(map[Position]Employee)

	for _, e := range s.employeesStorage.GetAll() {
		currentTop, exists := topPaid[e.Position]
		if !exists {
			topPaid[e.Position] = e
			continue
		}

		if e.Salary > currentTop.Salary {
			topPaid[e.Position] = e
		}
	}

	return topPaid
}

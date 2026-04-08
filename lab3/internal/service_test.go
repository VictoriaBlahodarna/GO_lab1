package internal_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/VictoriaBlahodarna/GO_lab1/lab3/internal"
)

type mockStorage struct {
	createFunc func(ctx context.Context, params internal.CreateEmployeePayload) (int, error)
	getFunc    func(ctx context.Context, id int) (internal.Employee, error)
}

func (m mockStorage) Create(ctx context.Context, params internal.CreateEmployeePayload) (int, error) {
	if m.createFunc != nil {
		return m.createFunc(ctx, params)
	}
	return 0, nil
}

func (m mockStorage) Get(ctx context.Context, id int) (internal.Employee, error) {
	if m.getFunc != nil {
		return m.getFunc(ctx, id)
	}
	return internal.Employee{}, nil
}

func TestEmployeeService_CreateEmployee(t *testing.T) {
	ctx := context.Background()

	t.Run("success flow - створює працівника успішно", func(t *testing.T) {
		mock := mockStorage{
			createFunc: func(ctx context.Context, params internal.CreateEmployeePayload) (int, error) {
				return 42, nil
			},
		}
		service := internal.NewEmployeeService(mock)

		validPayload := internal.CreateEmployeePayload{
			Name:     "John Doe",
			Position: "Developer",
			Salary:   3000,
		}

		id, err := service.CreateEmployee(ctx, validPayload)

		require.NoError(t, err)
		require.Equal(t, 42, id)
	})

	t.Run("fail flow - помилка від бази даних", func(t *testing.T) {
		expectedErr := errors.New("db connection lost")
		mock := mockStorage{
			createFunc: func(ctx context.Context, params internal.CreateEmployeePayload) (int, error) {
				return 0, expectedErr
			},
		}
		service := internal.NewEmployeeService(mock)

		validPayload := internal.CreateEmployeePayload{
			Name:     "John Doe",
			Position: "Developer",
			Salary:   3000,
		}

		id, err := service.CreateEmployee(ctx, validPayload)

		require.Error(t, err)
		require.ErrorContains(t, err, "failed to create employee in storage")
		require.Equal(t, 0, id)
	})

	t.Run("edge case - пусте ім'я", func(t *testing.T) {
		mock := mockStorage{}
		service := internal.NewEmployeeService(mock)

		invalidPayload := internal.CreateEmployeePayload{
			Name:     "",
			Position: "Developer",
			Salary:   3000,
		}

		id, err := service.CreateEmployee(ctx, invalidPayload)

		require.Error(t, err)
		require.ErrorIs(t, err, internal.ErrInvalidInput)
		require.Equal(t, 0, id)
	})

	t.Run("edge case - нульова зарплата", func(t *testing.T) {
		mock := mockStorage{}
		service := internal.NewEmployeeService(mock)

		invalidPayload := internal.CreateEmployeePayload{
			Name:     "Jane",
			Position: "Manager",
			Salary:   0,
		}

		id, err := service.CreateEmployee(ctx, invalidPayload)

		require.Error(t, err)
		require.ErrorIs(t, err, internal.ErrInvalidSalary)
		require.Equal(t, 0, id)
	})
}

func TestEmployeeService_GetEmployee(t *testing.T) {
	ctx := context.Background()

	t.Run("success flow - отримує працівника", func(t *testing.T) {
		expectedEmployee := internal.Employee{
			ID:     42,
			Name:   "Alice",
			Salary: 5000,
		}

		mock := mockStorage{
			getFunc: func(ctx context.Context, id int) (internal.Employee, error) {
				return expectedEmployee, nil
			},
		}
		service := internal.NewEmployeeService(mock)

		emp, err := service.GetEmployee(ctx, 42)

		require.NoError(t, err)
		require.Equal(t, expectedEmployee.Name, emp.Name)
	})

	t.Run("edge case - від'ємний або нульовий ID", func(t *testing.T) {
		mock := mockStorage{}
		service := internal.NewEmployeeService(mock)

		emp, err := service.GetEmployee(ctx, -5)

		require.Error(t, err)
		require.ErrorIs(t, err, internal.ErrInvalidInput)
		require.Equal(t, internal.Employee{}, emp)
	})
}

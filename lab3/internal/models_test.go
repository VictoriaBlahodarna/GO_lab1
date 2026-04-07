package internal_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/DenisGoldiner/webapp/internal"
)

func TestPosition_GetLevel(t *testing.T) {
	// Table-driven тестування
	tests := []struct {
		name      string
		minSalary uint
		expected  string
	}{
		{"Senior level", 3500, "Senior/Management"},
		{"Middle level bounds", 1500, "Middle"},
		{"Middle level normal", 2000, "Middle"},
		{"Junior level", 800, "Junior"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos := internal.Position{MinSalary: tt.minSalary}
			require.Equal(t, tt.expected, pos.GetLevel())
		})
	}
}

func TestEmployee_IsSalaryValid(t *testing.T) {
	pos := internal.Position{MinSalary: 1000, MaxSalary: 3000}

	t.Run("salary is exactly min limit", func(t *testing.T) {
		emp := internal.Employee{Salary: 1000, Position: pos}
		require.True(t, emp.IsSalaryValid())
	})

	t.Run("salary is valid and in range", func(t *testing.T) {
		emp := internal.Employee{Salary: 2000, Position: pos}
		require.True(t, emp.IsSalaryValid())
	})

	t.Run("salary too low", func(t *testing.T) {
		emp := internal.Employee{Salary: 500, Position: pos}
		require.False(t, emp.IsSalaryValid())
	})

	t.Run("salary too high", func(t *testing.T) {
		emp := internal.Employee{Salary: 4000, Position: pos}
		require.False(t, emp.IsSalaryValid())
	})
}

func TestEmployee_String(t *testing.T) {
	emp := internal.Employee{
		ID:   42,
		Name: "Alice",
		Position: internal.Position{
			MinSalary: 2000, // Значить це Middle
		},
		Salary: 2500,
	}

	expectedString := "42: Alice (Middle), $2500\n"
	require.Equal(t, expectedString, emp.String())
}

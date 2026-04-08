package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	"github.com/VictoriaBlahodarna/GO_lab1/lab3/internal"
)

type Client struct {
	dbExec sqlx.ExtContext
}

func NewClient(dbExec sqlx.ExtContext) Client {
	return Client{dbExec: dbExec}
}

func (c Client) Get(ctx context.Context, id int) (internal.Employee, error) {
	q := `
		SELECT e.id, e.name, p.name as position_name, p.min_salary, p.max_salary, e.salary 
		FROM employees e
		JOIN positions p ON e.position_id = p.id
		WHERE e.id = $1
	`

	type rawData struct {
		ID           int    `db:"id"`
		Name         string `db:"name"`
		PositionName string `db:"position_name"`
		MinSalary    uint   `db:"min_salary"`
		MaxSalary    uint   `db:"max_salary"`
		Salary       uint   `db:"salary"`
	}

	var row rawData
	err := sqlx.GetContext(ctx, c.dbExec, &row, q, id)
	if err != nil {
		return internal.Employee{}, fmt.Errorf("failed to fetch employee (id=%d): %w", id, err)
	}

	return internal.Employee{
		ID:   row.ID,
		Name: row.Name,
		Position: internal.Position{
			Name:      row.PositionName,
			MinSalary: row.MinSalary,
			MaxSalary: row.MaxSalary,
		},
		Salary: row.Salary,
	}, nil
}

func (c Client) Create(ctx context.Context, params internal.CreateEmployeePayload) (int, error) {
	var positionID int
	posQ := "SELECT id FROM positions WHERE name = $1"
	err := sqlx.GetContext(ctx, c.dbExec, &positionID, posQ, params.Position)
	if err != nil {
		return 0, fmt.Errorf("failed to find position '%s': %w", params.Position, err)
	}

	q := "INSERT INTO employees (name, position_id, salary) VALUES ($1, $2, $3) RETURNING id"

	var empID int
	err = sqlx.GetContext(ctx, c.dbExec, &empID, q, params.Name, positionID, params.Salary)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == pqErrCodeUniqueViolation {
			return 0, fmt.Errorf("employee already exists: %w", err)
		}
		return 0, fmt.Errorf("failed to create employee: %w", err)
	}

	return empID, nil
}

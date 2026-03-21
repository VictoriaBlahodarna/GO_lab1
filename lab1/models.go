package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrEmptyCompanyName = errors.New("company name cannot be empty")
	ErrEmployeeNotFound = errors.New("employee not found")
	ErrInvalidSalary    = errors.New("employee salary is out of position range")
)

type Company struct {
	name      string
	employees []Employee
}

func NewCompany(name string) (Company, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return Company{}, ErrEmptyCompanyName
	}
	return Company{
		name:      name,
		employees: []Employee{},
	}, nil
}

func (c *Company) AddEmployee(name string, position Position, salary uint) error {
	emp := NewEmployee(len(c.employees)+1, name, position, salary)
	if !emp.IsSalaryValid() {
		return ErrInvalidSalary
	}
	c.employees = append(c.employees, emp)
	return nil
}

func (c *Company) GetEmployee(id int) (Employee, error) {
	for _, e := range c.employees {
		if e.id == id {
			return e, nil
		}
	}
	return Employee{}, ErrEmployeeNotFound
}

func (c *Company) GetEmployeesByPosition(position Position) []Employee {
	employees := []Employee{}
	for _, e := range c.employees {
		if e.position == position {
			employees = append(employees, e)
		}
	}
	return employees
}

func (c *Company) GetTopPaidEmployees() map[Position]Employee {
	topPaid := make(map[Position]Employee)

	for _, e := range c.employees {
		currentTop, exists := topPaid[e.position]
		if !exists {
			topPaid[e.position] = e
			continue
		}

		if e.salary > currentTop.salary {
			topPaid[e.position] = e
		}
	}

	return topPaid
}

func (c *Company) String() string {
	emploeesPerPosition := make(map[Position][]Employee)
	for _, e := range c.employees {
		emploeesPerPosition[e.position] = append(emploeesPerPosition[e.position], e)
	}

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Company: %s, Employees: %d", c.name, len(c.employees)))

	for position, employees := range emploeesPerPosition {
		sb.WriteString(fmt.Sprintf("\n%s: \n", position.name))
		for _, e := range employees {
			sb.WriteString(e.String())
		}
	}
	return sb.String()
}

type Position struct {
	name      string
	minSalary Dollar
	maxSalary Dollar
}

func NewPosition(name string, minSalary, maxSalary uint) Position {
	return Position{
		name:      name,
		minSalary: Dollar(minSalary),
		maxSalary: Dollar(maxSalary),
	}
}

func (p Position) GetLevel() string {
	switch {
	case p.minSalary >= Dollar(3000*OneDollar):
		return "Senior/Management"
	case p.minSalary >= Dollar(1500*OneDollar):
		return "Middle"
	default:
		return "Junior"
	}
}

type Employee struct {
	id       int
	name     string
	position Position
	salary   Dollar
}

func NewEmployee(id int, name string, position Position, salary uint) Employee {
	return Employee{
		id:       id,
		name:     name,
		position: position,
		salary:   Dollar(salary),
	}
}

func (e Employee) IsSalaryValid() bool {
	return e.salary >= e.position.minSalary && e.salary <= e.position.maxSalary
}

func (e Employee) String() string {
	return fmt.Sprintf("%d: %s (%s), %s\n", e.id, e.name, e.position.GetLevel(), e.salary)
}

type Dollar uint

const OneDollar = 100

// Example: Dollar = 3075 is shown as $30.75
func (d Dollar) String() string {
	return fmt.Sprintf("$%d.%02d", d/OneDollar, d%OneDollar)
}

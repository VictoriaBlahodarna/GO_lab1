package main

import (
	"fmt"
	"strings"
)

// Position - struct that contains the name of the position,
//
//	minimum and maximum salary in cents
type Company struct {
	name      string
	employees []Employee
}

// TODO: Add validation for empty company name and trim spaces
func NewCompany(name string) Company {
	return Company{
		name:      name,
		employees: []Employee{},
	}
}

// TODO add validation for salary
func (c *Company) AddEmployee(name string, position Position, salary uint) {
	c.employees = append(c.employees, NewEmployee(len(c.employees)+1, name, position, salary))
}

func (c *Company) GetEmployee(id int) (Employee, bool) {
	for _, e := range c.employees {
		if e.id == id {
			return e, true
		}
	}
	return Employee{}, false
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

func (e Employee) String() string {
	return fmt.Sprintf("%d: %s, %s\n", e.id, e.name, e.salary)
}

type Dollar uint

// Example: Dollar = 3075 is shown as $30.75
func (d Dollar) String() string {
	return fmt.Sprintf("$%d.%02d", d/100, d%100)
}

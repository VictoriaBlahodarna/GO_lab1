package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCompany(t *testing.T) {
	tests := []struct {
		name     string
		compName string
		want     Company
	}{
		{
			name:     "Create new company",
			compName: "TestCorp",
			want: Company{
				name:      "TestCorp",
				employees: []Employee{},
			},
		},
		{
			name:     "Create company with empty name",
			compName: "",
			want: Company{
				name:      "",
				employees: []Employee{},
			},
		},
		{
			name:     "Create company with spaces name",
			compName: "   ",
			want: Company{
				name:      "",
				employees: []Employee{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCompany(tt.compName)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestCompany_AddEmployee(t *testing.T) {
	type args struct {
		name     string
		position Position
		salary   uint
	}
	tests := []struct {
		name         string
		initialState Company
		args         args
		wantLen      int
		wantLastEmp  Employee
	}{
		{
			name:         "Add first employee",
			initialState: NewCompany("TestCorp"),
			args: args{
				name:     "John Doe",
				position: NewPosition("Dev", 100, 200),
				salary:   150,
			},
			wantLen: 1,
			wantLastEmp: Employee{
				id:       1,
				name:     "John Doe",
				position: NewPosition("Dev", 100, 200),
				salary:   150,
			},
		},
		{
			name: "Add second employee",
			initialState: Company{
				name: "TestCorp",
				employees: []Employee{
					{
						id:       1,
						name:     "John Doe",
						position: NewPosition("Dev", 100, 200),
						salary:   150,
					},
				},
			},
			args: args{
				name:     "Jane Smith",
				position: NewPosition("QA", 80, 150),
				salary:   100,
			},
			wantLen: 2,
			wantLastEmp: Employee{
				id:       2,
				name:     "Jane Smith",
				position: NewPosition("QA", 80, 150),
				salary:   100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &tt.initialState
			c.AddEmployee(tt.args.name, tt.args.position, tt.args.salary)
			require.Len(t, c.employees, tt.wantLen)
			require.Equal(t, tt.wantLastEmp, c.employees[len(c.employees)-1])
		})
	}
}

func TestCompany_GetEmployee(t *testing.T) {
	devPos := NewPosition("Dev", 1000, 2000)
	emp1 := NewEmployee(1, "Alice", devPos, 1500)
	emp2 := NewEmployee(2, "Bob", devPos, 1600)

	comp := Company{
		name:      "TestCorp",
		employees: []Employee{emp1, emp2},
	}

	tests := []struct {
		name  string
		c     *Company
		id    int
		want  Employee
		want1 bool
	}{
		{
			name:  "Existing employee 1",
			c:     &comp,
			id:    1,
			want:  emp1,
			want1: true,
		},
		{
			name:  "Existing employee 2",
			c:     &comp,
			id:    2,
			want:  emp2,
			want1: true,
		},
		{
			name:  "Non-existing employee",
			c:     &comp,
			id:    3,
			want:  Employee{},
			want1: false,
		},
		{
			name:  "Negative ID",
			c:     &comp,
			id:    -1,
			want:  Employee{},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.c.GetEmployee(tt.id)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.want1, got1)
		})
	}
}

func TestCompany_GetEmployeesByPosition(t *testing.T) {
	devPos := NewPosition("Dev", 1000, 2000)
	qaPos := NewPosition("QA", 800, 1500)
	managerPos := NewPosition("Manager", 2000, 3000)

	emp1 := NewEmployee(1, "Alice", devPos, 1500)
	emp2 := NewEmployee(2, "Bob", devPos, 1600)
	emp3 := NewEmployee(3, "Charlie", qaPos, 1000)

	comp := Company{
		name:      "TestCorp",
		employees: []Employee{emp1, emp2, emp3},
	}

	tests := []struct {
		name     string
		c        *Company
		position Position
		want     []Employee
	}{
		{
			name:     "Multiple employees",
			c:        &comp,
			position: devPos,
			want:     []Employee{emp1, emp2},
		},
		{
			name:     "Single employee",
			c:        &comp,
			position: qaPos,
			want:     []Employee{emp3},
		},
		{
			name:     "No employees",
			c:        &comp,
			position: managerPos,
			want:     []Employee{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.GetEmployeesByPosition(tt.position)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestCompany_String(t *testing.T) {
	devPos := NewPosition("Dev", 1000, 2000)
	emp1 := NewEmployee(1, "Alice", devPos, 1500)

	comp := Company{
		name:      "TestCorp",
		employees: []Employee{emp1},
	}

	emptyComp := Company{
		name:      "EmptyCorp",
		employees: []Employee{},
	}

	tests := []struct {
		name string
		c    *Company
		want []string
	}{
		{
			name: "Company with 1 employee",
			c:    &comp,
			want: []string{
				"Company: TestCorp, Employees: 1",
				"Dev:",
				"1: Alice, $15.00",
			},
		},
		{
			name: "Empty company",
			c:    &emptyComp,
			want: []string{
				"Company: EmptyCorp, Employees: 0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.String()
			for _, w := range tt.want {
				require.Contains(t, got, w)
			}
		})
	}
}

func TestNewPosition(t *testing.T) {
	type args struct {
		name      string
		minSalary uint
		maxSalary uint
	}
	tests := []struct {
		name string
		args args
		want Position
	}{
		{
			name: "Create normal position",
			args: args{
				name:      "Dev",
				minSalary: 1000,
				maxSalary: 2000,
			},
			want: Position{
				name:      "Dev",
				minSalary: Dollar(1000),
				maxSalary: Dollar(2000),
			},
		},
		{
			name: "Create position with 0 salary",
			args: args{
				name:      "Intern",
				minSalary: 0,
				maxSalary: 0,
			},
			want: Position{
				name:      "Intern",
				minSalary: Dollar(0),
				maxSalary: Dollar(0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPosition(tt.args.name, tt.args.minSalary, tt.args.maxSalary)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestNewEmployee(t *testing.T) {
	devPos := NewPosition("Dev", 1000, 2000)
	type args struct {
		id       int
		name     string
		position Position
		salary   uint
	}
	tests := []struct {
		name string
		args args
		want Employee
	}{
		{
			name: "Create normal employee",
			args: args{
				id:       1,
				name:     "Alice",
				position: devPos,
				salary:   1500,
			},
			want: Employee{
				id:       1,
				name:     "Alice",
				position: devPos,
				salary:   Dollar(1500),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewEmployee(tt.args.id, tt.args.name, tt.args.position, tt.args.salary)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestEmployee_String(t *testing.T) {
	devPos := NewPosition("Dev", 1000, 2000)
	tests := []struct {
		name string
		e    Employee
		want string
	}{
		{
			name: "Format employee",
			e:    NewEmployee(1, "Alice", devPos, 1500),
			want: "1: Alice, $15.00\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.String()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestDollar_String(t *testing.T) {
	tests := []struct {
		name string
		d    Dollar
		want string
	}{
		{
			name: "Zero dollars",
			d:    Dollar(0),
			want: "$0.00",
		},
		{
			name: "Exactly dollars",
			d:    Dollar(100),
			want: "$1.00",
		},
		{
			name: "Dollars and cents",
			d:    Dollar(3075),
			want: "$30.75",
		},
		{
			name: "Only cents",
			d:    Dollar(99),
			want: "$0.99",
		},
		{
			name: "Single digit cent",
			d:    Dollar(5),
			want: "$0.05",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.d.String()
			require.Equal(t, tt.want, got)
		})
	}
}

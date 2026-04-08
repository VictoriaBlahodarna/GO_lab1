package postgres

const pqErrCodeUniqueViolation = "23505"

type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	PositionID int    `db:"position_id"`
	Salary     uint   `db:"salary"`
}

type Position struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	MinSalary uint   `db:"min_salary"`
	MaxSalary uint   `db:"max_salary"`
}

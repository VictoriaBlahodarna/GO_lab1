package main

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/VictoriaBlahodarna/GO_lab1/lab3/internal"
	"github.com/VictoriaBlahodarna/GO_lab1/lab3/internal/adapters/postgres"
	"github.com/VictoriaBlahodarna/GO_lab1/lab3/internal/ports/csv"
)

func main() {
	log.Println("Starting CSV Importer")
	start := time.Now()

	dbExec, err := newDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	employeesClient := postgres.NewClient(dbExec)
	employeesService := internal.NewEmployeeService(employeesClient)

	parser := csv.NewParser(employeesService)

	ctx := context.Background()
	filepath := "employees_10000.csv"

	log.Printf("Starting import of file: %s", filepath)

	err = parser.Run(ctx, filepath)
	if err != nil {
		log.Fatalf("import failed with error: %v", err)
	}

	duration := time.Since(start)
	log.Printf("Successfully imported all records in %v", duration)
}

func newDB() (sqlx.ExtContext, error) {
	dsn := "postgres://postgres:postgres@localhost:5432/employees?sslmode=disable"
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

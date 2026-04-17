package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/exp/slog"

	"github.com/VictoriaBlahodarna/GO_lab1/lab3/internal"
	"github.com/VictoriaBlahodarna/GO_lab1/lab3/internal/adapters/postgres"
	"github.com/VictoriaBlahodarna/GO_lab1/lab3/internal/ports/rest"
)

func main() {
	app := newApplication()

	slog.Info("Starting Employee application")

	app.start()
}

type application struct {
	server *http.Server
}

func newApplication() application {
	dbExec, err := newDB()
	if err != nil {
		log.Fatal(err)
	}

	server := newServer(dbExec)

	return application{
		server: server,
	}
}

func newDB() (sqlx.ExtContext, error) {
	// База даних залишена travellers для сумісності з docker-compose.yml
	dsn := "postgres://postgres:postgres@localhost:5432/employees?sslmode=disable"
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func newServer(dbExec sqlx.ExtContext) *http.Server {
	employeesClient := postgres.NewClient(dbExec)
	employeesService := internal.NewEmployeeService(employeesClient)

	handlers := map[string]http.Handler{
		"/api/v1/employees":          rest.NewEmployeeHandler(employeesService),
		"/api/v1/employees/top-paid": rest.NewTopPaidHandler(employeesService),
	}

	mux := http.NewServeMux()
	for route, handler := range handlers {
		mux.Handle(route, handler)
	}

	return &http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}
}

func (app application) start() {
	if err := app.server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return
		}

		log.Printf("failed to start the HTTP server, error: %v", err)
	}
}

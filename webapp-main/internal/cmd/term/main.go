package main

import (
	"context"
	"github.com/DenisGoldiner/webapp/internal/ports/ftp"
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/exp/trace"
	"log"
	"os"

	"github.com/DenisGoldiner/webapp/internal"
	"github.com/DenisGoldiner/webapp/internal/adapters/postgres"
)

func main() {
	run()
}

func run() {

	fr := trace.NewFlightRecorder()
	if err := fr.Start(); err != nil {
		// handle error
	}

	defer func() {
		if err := fr.Stop(); err != nil {
			// handle error
		}
	}()

	app := fiber.New()

	app.Get("/trace", func(ctx fiber.Ctx) error {
		f, err := os.OpenFile("file.out", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
		if err != nil {
			// handle error
		}

		if _, err := fr.WriteTo(f); err != nil {
			// handle error
		}
		return ctx.SendStatus(fiber.StatusOK)
	})

	dbExec, err := newDB()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	travellersClient := postgres.NewClient(dbExec)
	travellersService := internal.NewTravellers(travellersClient)
	travellersParser := ftp.NewParser(travellersService)

	if err = travellersParser.Run(ctx, "/Users/denys/Go/src/github.com/DenisGoldiner/webapp/internal/integration/data/test_1.csv"); err != nil {
		log.Printf("error running travellers import: %v", err)
		return
	}

	if err := app.Listen(":8080"); err != nil {
		// handle error
		app.Shutdown()
	}
}

func newDB() (sqlx.ExtContext, error) {
	dsn := "postgres://postgres:postgres@localhost:5432/travellers?sslmode=disable"
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

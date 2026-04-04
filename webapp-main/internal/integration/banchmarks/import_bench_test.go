package banchmarks

import (
	"context"
	"github.com/DenisGoldiner/webapp/internal"
	"github.com/DenisGoldiner/webapp/internal/adapters/postgres"
	"github.com/DenisGoldiner/webapp/internal/ports/ftp"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"testing"
)

func BenchmarkSample(b *testing.B) {
	dbExec, err := newDB()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	travellersClient := postgres.NewClient(dbExec)
	travellersService := internal.NewTravellers(travellersClient)
	travellersParser := ftp.NewParser(travellersService)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err = travellersParser.Run(ctx, "/Users/denys/Go/src/github.com/DenisGoldiner/webapp/internal/integration/data/test_1.csv"); err != nil {
			b.Fatalf("unespected error, %v", err)
		}
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

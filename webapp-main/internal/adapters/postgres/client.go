package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	"github.com/DenisGoldiner/webapp/internal"
)

type Client struct {
	dbExec sqlx.ExtContext
}

func NewClient(dbExec sqlx.ExtContext) Client {
	return Client{dbExec: dbExec}
}

func (c Client) Get(ctx context.Context, id uuid.UUID) (internal.Traveller, error) {
	q := "select id, first_name, last_name, age from travellers where id = $1"

	rows, err := c.dbExec.QueryxContext(ctx, q, id)
	if err != nil {
		return internal.Traveller{}, fmt.Errorf("failed to fetch traveler: %w", err)
	}

	defer func() { _ = rows.Close() }()

	var travelers []Traveller

	for rows.Next() {
		var traveler Traveller

		if err = rows.StructScan(&traveler); err != nil {
			return internal.Traveller{}, fmt.Errorf("failed to scan traveler: %w", err)
		}

		travelers = append(travelers, traveler)
	}

	if len(travelers) == 0 {
		return internal.Traveller{}, fmt.Errorf("no travelers with id %s: %w", id, internal.ErrNoResource)
	}

	return internal.Traveller{
		ID:        travelers[0].ID,
		FirstName: travelers[0].FirstName,
		LastName:  travelers[0].LastName,
		Age:       travelers[0].Age,
	}, err
}

func (c Client) Create(ctx context.Context, params internal.CreateTravellerPayload) (uuid.UUID, error) {
	q := "insert into travellers (first_name, last_name, age) values ($1, $2, $3) returning id"

	rows, err := c.dbExec.QueryxContext(ctx, q, params.FirstName, params.LastName, params.Age)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == pqErrCodeUniqueViolation {
			return uuid.Nil, fmt.Errorf("traveler already exists: %w", internal.ErrAlreadyExists)
		}

		return uuid.Nil, fmt.Errorf("failed to create traveler: %w", err)
	}

	defer func() { _ = rows.Close() }()

	var travelerIDs []uuid.UUID

	for rows.Next() {
		var travelerID uuid.UUID

		if err = rows.Scan(&travelerID); err != nil {
			return uuid.Nil, fmt.Errorf("failed to scan traveler id: %w", err)
		}

		travelerIDs = append(travelerIDs, travelerID)
	}

	if len(travelerIDs) == 0 {
		return uuid.Nil, fmt.Errorf("failed to create traveler: no id returned")
	}

	return travelerIDs[0], nil
}

func (c Client) DeleteTraveller() {

}

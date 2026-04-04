package ftp

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/DenisGoldiner/webapp/internal"
	"io"
	"os"
	"strconv"
)

type Parser struct {
	service internal.Travellers
}

func NewParser(service internal.Travellers) Parser {
	return Parser{service: service}
}

func (p Parser) Run(ctx context.Context, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open the file %s: %w", filePath, err)
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	reader := csv.NewReader(bytes.NewReader(data))

	travelers, err := p.parse(reader)
	if err != nil {
		return err
	}

	if err = p.process(ctx, travelers); err != nil {
		return err
	}

	return nil
}

func (p Parser) parse(r *csv.Reader) ([]internal.Traveller, error) {
	var travelers []internal.Traveller

	for i := 0; ; i++ {
		row, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("failed to parse row #%d: %w", i, err)
		}

		age, err := strconv.Atoi(row[2])
		if err != nil {
			return nil, fmt.Errorf("failed to parse age vaue %s in #%d: %w", row[2], i, err)
		}

		traveler := internal.Traveller{
			FirstName: row[0],
			LastName:  row[1],
			Age:       age,
		}

		travelers = append(travelers, traveler)
	}

	return travelers, nil
}

func (p Parser) process(ctx context.Context, travelers []internal.Traveller) error {
	for _, traveler := range travelers {
		if _, err := p.service.Create(ctx, traveler); err != nil {
			return fmt.Errorf("failed to create traveller: %w", err)
		}
	}

	return nil
}

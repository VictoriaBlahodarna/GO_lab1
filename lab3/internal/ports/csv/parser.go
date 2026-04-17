package csv

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/VictoriaBlahodarna/GO_lab1/lab3/internal"
)

type Parser struct {
	service internal.EmployeeService
}

func NewParser(service internal.EmployeeService) Parser {
	return Parser{service: service}
}

func (p Parser) Run(ctx context.Context, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open the file %s: %w", filePath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	const batchSize = 500

	batch := make([]internal.CreateEmployeePayload, 0, batchSize)

	for i := 0; ; i++ {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to parse row #%d: %w", i, err)
		}

		if len(row) != 3 {
			return fmt.Errorf("invalid number of columns in row #%d: %v", i, row)
		}

		salary, err := strconv.ParseUint(row[2], 10, 32)
		if err != nil {
			return fmt.Errorf("failed to parse salary value %s in row #%d: %w", row[2], i, err)
		}

		employee := internal.CreateEmployeePayload{
			Name:     strings.TrimSpace(row[0]),
			Position: strings.TrimSpace(row[1]),
			Salary:   uint(salary),
		}

		batch = append(batch, employee)

		if len(batch) >= batchSize {
			if err := p.service.BulkCreateEmployees(ctx, batch); err != nil {
				return err
			}
			batch = batch[:0]
		}
	}

	if len(batch) > 0 {
		if err := p.service.BulkCreateEmployees(ctx, batch); err != nil {
			return err
		}
	}

	return nil
}

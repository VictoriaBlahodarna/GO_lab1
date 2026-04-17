package csv

import (
	"context"
	"os"
	"testing"

	"github.com/VictoriaBlahodarna/GO_lab1/lab3/internal"
)

type mockStorage struct{}

func (m mockStorage) Get(ctx context.Context, id int) (internal.Employee, error) {
	return internal.Employee{}, nil
}

func (m mockStorage) Create(ctx context.Context, params internal.CreateEmployeePayload) (int, error) {
	return 1, nil
}

func (m mockStorage) GetAll() []internal.Employee {
	return nil
}

func (m mockStorage) BulkCreate(ctx context.Context, params []internal.CreateEmployeePayload) error {
	return nil
}

func BenchmarkParseCSV(b *testing.B) {
	storage := mockStorage{}
	service := internal.NewEmployeeService(storage)
	parser := NewParser(service)

	filePath := "../../../employees_10000.csv"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		b.Fatalf("Test file %s does not exist", filePath)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := parser.Run(context.Background(), filePath)
		if err != nil {
			b.Fatalf("Parser Run failed: %v", err)
		}
	}
}

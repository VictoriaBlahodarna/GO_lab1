package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var (
	firstNames = []string{"Alex", "Alice", "Bob", "Victoria", "John", "Jane", "Mike", "Sarah", "David", "Emma"}
	lastNames  = []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis"}
	positions  = []string{"Developer", "Manager", "QA"}
)

func main() {
	filename := "employees_10000.csv"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Не вдалося створити файл: %v\n", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Опціонально можна записати заголовки (якщо наш парсер їх підтримуватиме)
	// writer.Write([]string{"Name", "Position", "Salary"})

	recordsCount := 10000
	for i := 0; i < recordsCount; i++ {
		name := fmt.Sprintf("%s %s", firstNames[rand.Intn(len(firstNames))], lastNames[rand.Intn(len(lastNames))])
		position := positions[rand.Intn(len(positions))]
		salary := generateSalaryForPosition(position)

		record := []string{name, position, strconv.Itoa(salary)}
		if err := writer.Write(record); err != nil {
			fmt.Printf("Помилка запису рядка %d: %v\n", i, err)
			return
		}
	}

	fmt.Printf("Успішно згенеровано %d рядків у файл %s\n", recordsCount, filename)
}

func generateSalaryForPosition(pos string) int {
	switch pos {
	case "Developer":
		// Ліміти з init.sql: 1000 - 8000
		return 1000 + rand.Intn(7000)
	case "Manager":
		// Ліміти з init.sql: 2000 - 10000
		return 2000 + rand.Intn(8000)
	case "QA":
		// Ліміти з init.sql: 800 - 4000
		return 800 + rand.Intn(3200)
	default:
		return 1000
	}
}

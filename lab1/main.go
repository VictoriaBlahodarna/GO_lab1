package main

import "fmt"

func main() {
	// Створюємо посад
	developer := NewPosition("Developer", 200000, 800000)
	manager := NewPosition("Manager", 300000, 1000000)
	qaEngineer := NewPosition("QA Engineer", 150000, 500000)

	// Створюємо структуру компанії
	company := NewCompany("Tech Solutions Inc.")

	// Наймаємо співробітників
	company.AddEmployee("Олександр", developer, 450000) // Отримуватиме $4500.00
	company.AddEmployee("Марія", developer, 550000)     // Отримуватиме $5500.00
	company.AddEmployee("Іван", manager, 700000)        // Отримуватиме $7000.00
	company.AddEmployee("Анна", qaEngineer, 250000)     // Отримуватиме $2500.00

	// Виводимо всю інформацію про компанію.
	// Оскільки метод String() визначено для вказівника (*Company), до fmt.Println
	// передаємо посилання на компанію за допомогою амперсанда (&).
	fmt.Println(&company)
}

package main

import "fmt"

func main() {
	developer := NewPosition("Developer", 200000, 800000)
	manager := NewPosition("Manager", 300000, 1000000)
	qaEngineer := NewPosition("QA Engineer", 150000, 500000)

	company, err := NewCompany("Tech Solutions Inc.")
	if err != nil {
		fmt.Printf("Failed to create company: %v\n", err)
		return
	}

	if err := company.AddEmployee("Олександр", developer, 450000); err != nil {
		fmt.Printf("Помилка найму Олександра: %v\n", err)
	}

	if err := company.AddEmployee("Марія", developer, 550000); err != nil {
		fmt.Printf("Помилка найму Марії: %v\n", err)
	}

	if err := company.AddEmployee("Іван", manager, 700000); err != nil {
		fmt.Printf("Помилка найму Івана: %v\n", err)
	}

	if err := company.AddEmployee("Анна", qaEngineer, 250000); err != nil {
		fmt.Printf("Помилка найму Анни: %v\n", err)
	}

	// неправильна зарплата
	if err := company.AddEmployee("Олексій", qaEngineer, 900000); err != nil {
		fmt.Printf("Очікувана помилка найму Олексія: %v\n", err)
	}

	fmt.Println(&company)

	fmt.Println("\nТоп працівники за посадами:")
	topPaid := company.GetTopPaidEmployees()
	for pos, emp := range topPaid {
		fmt.Printf("%s: %s (Зарплата: %s)\n", pos.name, emp.name, emp.salary)
	}
}

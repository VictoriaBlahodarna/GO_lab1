package main

import "fmt"

func main() {
	developer := NewPosition("Developer", 2000*OneDollar, 8000*OneDollar)
	manager := NewPosition("Manager", 3000*OneDollar, 10000*OneDollar)
	qaEngineer := NewPosition("QA Engineer", 1500*OneDollar, 5000*OneDollar)

	company, err := NewCompany("Tech Solutions Inc.")
	if err != nil {
		fmt.Printf("Failed to create company: %v\n", err)
		return
	}

	if err := company.AddEmployee("Олександр", developer, 4500*OneDollar); err != nil {
		fmt.Printf("Помилка найму Олександра: %v\n", err)
	}

	if err := company.AddEmployee("Марія", developer, 5500*OneDollar); err != nil {
		fmt.Printf("Помилка найму Марії: %v\n", err)
	}

	if err := company.AddEmployee("Іван", manager, 7000*OneDollar); err != nil {
		fmt.Printf("Помилка найму Івана: %v\n", err)
	}

	if err := company.AddEmployee("Анна", qaEngineer, 2500*OneDollar); err != nil {
		fmt.Printf("Помилка найму Анни: %v\n", err)
	}

	// неправильна зарплата
	if err := company.AddEmployee("Олексій", qaEngineer, 9000*OneDollar); err != nil {
		fmt.Printf("Очікувана помилка найму Олексія: %v\n", err)
	}

	fmt.Println(&company)

	fmt.Println("\nТоп працівники за посадами:")
	topPaid := company.GetTopPaidEmployees()
	for pos, emp := range topPaid {
		fmt.Printf("%s: %s (Зарплата: %s)\n", pos.name, emp.name, emp.salary)
	}
}

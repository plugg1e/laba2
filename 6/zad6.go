package main

import "fmt"

func countSubordinates(employees map[string]string) map[string]int {
	subordinatesCount := make(map[string]int)

	for employee, manager := range employees {
		for manager != "" {
			subordinatesCount[manager]++
			manager = employees[manager]
		}
	}

	return subordinatesCount
}

func main() {
	employees := map[string]string{
		"A": "B",
		"C": "B",
		"D": "E",
		"B": "E",
		"E": "E",
	}

	result := countSubordinates(employees)
	fmt.Println("Количество подчиненных:", result)
}

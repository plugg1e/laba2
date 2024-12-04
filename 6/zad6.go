package main

import (
	"fmt"
)

// Функция для подсчета подчиненных
func countSubordinates(employees map[string]string, manager string, subordinatesCount map[string]int) int {
	if subordinatesCount[manager] != 0 {
		return subordinatesCount[manager]
	}

	count := 0
	for employee, m := range employees {
		if m == manager {
			count += 1 + countSubordinates(employees, employee, subordinatesCount)
		}
	}

	subordinatesCount[manager] = count
	return count
}

func main() {
	employees := map[string]string{
		"A": "B",
		"C": "B",
		"D": "E",
		"B": "E",
		"E": "E",
	}

	subordinatesCount := make(map[string]int)

	// Инициализируем счетчик для всех сотрудников
	for employee := range employees {
		subordinatesCount[employee] = 0
	}

	// Подсчитываем количество подчиненных для каждого менеджера
	for manager := range employees {
		countSubordinates(employees, manager, subordinatesCount)
	}

	fmt.Println("Subordinates count:", subordinatesCount)
}

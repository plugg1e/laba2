package main

import "fmt"

// подсчет подчиненых
func countSubordinates(employees map[string]string) map[string]int {

	// мап для хранения количества подчиненных
	subordinatesCount := make(map[string]int)

	// рекрсия по подсчету подчиненных
	var count func(manager string) int
	count = func(manager string) int {

		if val, ok := subordinatesCount[manager]; ok {
			return val
		}

		// инициализциф колва подчиненных для менеджера
		subordinatesCount[manager] = 0

		// обход по всем сотрудникам и подсчитываем подчиненных
		for employee, m := range employees {
			if m == manager {
				subordinatesCount[manager] += 1 + count(employee)
			}
		}

		return subordinatesCount[manager]
	}

	// количество подчиненных для каждого менеджера
	for manager := range employees {
		count(manager)
	}

	return subordinatesCount
}

func main() {
	employees := map[string]string{
		"A": "B",
		"C": "B",
		"D": "E",
		"B": "E",
		"E": "D",
	}

	result := countSubordinates(employees)
	fmt.Println("Количество подчиненных:", result)
}

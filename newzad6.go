package main

import (
	"fmt"
	"hash/fnv"
)

type Item struct {
	key   string
	value string
	next  *Item
}

type HashTab struct {
	sizeArr int
	tabl    []*Item
}

func NewHashTab(size int) *HashTab {
	return &HashTab{
		sizeArr: size,
		tabl:    make([]*Item, size),
	}
}

func (ht *HashTab) Hash(itemKey string) int {
	h := fnv.New32a()
	h.Write([]byte(itemKey))
	return int(h.Sum32()) % ht.sizeArr
}

func (ht *HashTab) AddHash(key, value string) {
	index := ht.Hash(key)
	if ht.tabl[index] != nil && ht.tabl[index].key == key {
		fmt.Printf(" ключ '%s' уже существует. значение не добавлено.\n", ht.tabl[index].key)
		return
	}
	newNode := &Item{key: key, value: value, next: ht.tabl[index]}
	ht.tabl[index] = newNode
}

func countSubordinates(employees *HashTab) map[string]int {
	// мап для хранения количества подчиненных
	subordinatesCount := make(map[string]int)

	// рекрсия по подсчету подчиненных
	var count func(manager string, visited map[string]bool) int
	count = func(manager string, visited map[string]bool) int {
		if val, ok := subordinatesCount[manager]; ok {
			return val
		}

		// инициализциф колва подчиненных для менеджера
		subordinatesCount[manager] = 0

		// обход по всем сотрудникам и подсчитываем подчиненных
		for i := 0; i < employees.sizeArr; i++ {
			current := employees.tabl[i]
			for current != nil {
				if current.value == manager {
					// Проверка на зацикливание
					if visited[current.key] {
						continue // Пропускаем, если сотрудник уже был посещен
					}
					visited[current.key] = true
					subordinatesCount[manager] += 1 + count(current.key, visited)
				}
				current = current.next
			}
		}

		return subordinatesCount[manager]
	}

	// количество подчиненных для каждого менеджера
	for i := 0; i < employees.sizeArr; i++ {
		current := employees.tabl[i]
		for current != nil {
			visited := make(map[string]bool)
			count(current.value, visited)
			current = current.next
		}
	}

	return subordinatesCount
}

func main() {
	employees := NewHashTab(10)
	employees.AddHash("A", "B")
	employees.AddHash("C", "B")
	employees.AddHash("D", "E")
	employees.AddHash("B", "E")
	employees.AddHash("E", "E")

	result := countSubordinates(employees)
	fmt.Println("Количество подчиненных:", result)
}

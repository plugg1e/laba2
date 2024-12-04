package main

import (
	"fmt"
	"hash/fnv"
)

const (
	tableSize = 100
)

type Set struct {
	table []*Node
}

type Node struct {
	key  string
	next *Node
}

func NewSet() *Set {
	return &Set{
		table: make([]*Node, tableSize),
	}
}

func (s *Set) hash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32()) % tableSize
}

func (s *Set) Add(key string) {
	index := s.hash(key)
	if s.table[index] == nil {
		s.table[index] = &Node{key: key}
	} else {
		current := s.table[index]
		for current.next != nil {
			if current.key == key {
				return // элемент существует
			}
			current = current.next
		}
		if current.key == key {
			return // элемент существует
		}
		current.next = &Node{key: key}
	}
}

func (s *Set) Remove(key string) {
	index := s.hash(key)
	if s.table[index] == nil {
		return
	}
	if s.table[index].key == key {
		s.table[index] = s.table[index].next
		return
	}
	prev := s.table[index]
	current := prev.next
	for current != nil {
		if current.key == key {
			prev.next = current.next
			return
		}
		prev = current
		current = current.next
	}
}

func (s *Set) Contains(key string) bool {
	index := s.hash(key)
	current := s.table[index]
	for current != nil {
		if current.key == key {
			return true
		}
		current = current.next
	}
	return false
}

func main() {
	set := NewSet()

	// добавление элементов
	set.Add("apple")
	set.Add("banana")

	// проверка наличия
	fmt.Println("содержит 'яяблк':", set.Contains("apple"))
	fmt.Println("содержит 'банан':", set.Contains("banana"))
	fmt.Println("содержит 'вишня':", set.Contains("cherry"))
	fmt.Println("содержит 'хлеб':", set.Contains("date"))

	// удаление элемента
	set.Remove("banana")
	fmt.Println("содержит 'банан' после удаления:", set.Contains("banana")) // false
}

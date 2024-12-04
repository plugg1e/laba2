package main

import (
	"fmt"
	"strings"
)

func evaluateExpression(expression string) int {
	// Удаляем пробелы из выражения
	expression = strings.ReplaceAll(expression, " ", "")

	// Стек для операндов
	operands := make([]int, 0)
	// Стек для операторов
	operators := make([]rune, 0)

	// Функция для применения оператора к двум операндам
	applyOperator := func() {
		operator := operators[len(operators)-1]
		operators = operators[:len(operators)-1]

		b := operands[len(operands)-1]
		a := operands[len(operands)-2]
		operands = operands[:len(operands)-2]

		switch operator {
		case '&':
			operands = append(operands, a&b)
		case '|':
			operands = append(operands, a|b)
		case '^':
			operands = append(operands, a^b)
		}
	}

	for _, char := range expression {
		switch char {
		case '0', '1':
			operands = append(operands, int(char-'0'))
		case '!':
			// Применяем отрицание к следующему операнду
			next := operands[len(operands)-1]
			operands[len(operands)-1] = 1 - next
		case '&', '|', '^':
			// Применяем операторы в соответствии с приоритетом
			for len(operators) > 0 && operators[len(operators)-1] != '(' {
				applyOperator()
			}
			operators = append(operators, char)
		case '(':
			operators = append(operators, char)
		case ')':
			// Применяем все операторы до открывающей скобки
			for len(operators) > 0 && operators[len(operators)-1] != '(' {
				applyOperator()
			}
			operators = operators[:len(operators)-1]
		}
	}

	// Применяем оставшиеся операторы
	for len(operators) > 0 {
		applyOperator()
	}

	return operands[0]
}

func main() {
	expression := "1 & (0 | 1) ^ 0"
	result := evaluateExpression(expression)
	fmt.Println("Результат:", result)
}

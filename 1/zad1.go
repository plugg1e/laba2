package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	data []string
}

func (s *Stack) Push(item string) {
	s.data = append(s.data, item)
}

func (s *Stack) Pop() string {
	if len(s.data) == 0 {
		return ""
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item
}

func (s *Stack) Peek() string {
	if len(s.data) == 0 {
		return ""
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func precedence(op string) int {
	switch op {
	case "!":
		return 3
	case "&":
		return 2
	case "|", "^":
		return 1
	default:
		return 0
	}
}

func infixToPostfix(expression string) string {
	var output strings.Builder
	var stack Stack

	for _, char := range expression {
		token := string(char)

		switch token {
		case "0", "1":
			output.WriteString(token)
		case "(":
			stack.Push(token)
		case ")":
			for !stack.IsEmpty() && stack.Peek() != "(" {
				output.WriteString(stack.Pop())
			}
			stack.Pop()
		default:
			for !stack.IsEmpty() && precedence(stack.Peek()) >= precedence(token) {
				output.WriteString(stack.Pop())
			}
			stack.Push(token)
		}
	}

	for !stack.IsEmpty() {
		output.WriteString(stack.Pop())
	}

	return output.String()
}

func evaluatePostfix(expression string) int {
	var stack Stack

	for _, char := range expression {
		token := string(char)

		switch token {
		case "0", "1":
			stack.Push(token)
		case "!":
			val := stack.Pop()
			if val == "0" {
				stack.Push("1")
			} else {
				stack.Push("0")
			}
		case "&", "|", "^":
			val2 := stack.Pop()
			val1 := stack.Pop()
			result := evaluateOperation(val1, val2, token)
			stack.Push(result)
		}
	}

	return toInt(stack.Pop())
}

func evaluateOperation(val1, val2, op string) string {
	a := toInt(val1)
	b := toInt(val2)

	switch op {
	case "&":
		return toStr(a & b)
	case "|":
		return toStr(a | b)
	case "^":
		return toStr(a ^ b)
	default:
		return "0"
	}
}

func toInt(val string) int {
	if val == "1" {
		return 1
	}
	return 0
}

func toStr(val int) string {
	if val == 1 {
		return "1"
	}
	return "0"
}

func main() {
	expression := "1&(0|1)^0|1"
	postfix := infixToPostfix(expression)
	result := evaluatePostfix(postfix)
	fmt.Println("овтет:", result)
}

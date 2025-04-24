package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

func PrefixToInfix(prefix string) (string, error) {
	if prefix == "" {
		return "", fmt.Errorf("порожній вираз")
	}

	tokens := strings.Split(prefix, " ")
	stack := []string{}
	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]
		if token == "" {
			continue
		}
		if num, err := strconv.Atoi(token); err == nil {
			if num < 0 {
				stack = append(stack, fmt.Sprintf("(%d)", num))
			} else {
				stack = append(stack, token)
			}
		} else if isOperator(token) {
			if len(stack) < 2 {
				return "", fmt.Errorf("некоректний вираз: недостатньо операндів (%d) для оператора '%s'", len(stack), token)
			}
			op1 := stack[len(stack)-1]
			op2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			expr := "(" + op1 + token + op2 + ")"
			stack = append(stack, expr)
		} else {
			return "", fmt.Errorf("некоректний токен: '%s'", token)
		}
	}
	if len(stack) != 1 {
		return "", fmt.Errorf("некоректний вираз: неправильна структура (фінальний стек: %v)", stack)
	}
	return stack[0], nil
}
func isOperator(token string) bool {
	switch token {
	case "+", "-", "*", "/", "^":
		return true
	default:
		return false
	}
}

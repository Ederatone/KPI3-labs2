package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

func PrefixToInfix(prefix string) (string, error) {
	tokens := strings.Split(prefix, " ")
	stack := []string{}

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]

		if isOperator(token) {
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid prefix expression: not enough operands for operator %s", token)
			}
			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			expression := "(" + operand1 + token + operand2 + ")"
			stack = append(stack, expression)
		} else if _, err := strconv.Atoi(token); err == nil {
			stack = append(stack, token)
		} else {
			return "", fmt.Errorf("invalid token: %s", token)
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("invalid prefix expression")
	}

	return stack[0], nil
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/" || token == "^"
}

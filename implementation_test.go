package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToInfix(t *testing.T) {
	tests := []struct {
		name          string
		prefix        string
		expectedInfix string
		expectError   bool
	}{
		{
			name:          "Simple addition",
			prefix:        "+ 2 3",
			expectedInfix: "(2+3)",
			expectError:   false,
		},
		{
			name:          "Simple subtraction",
			prefix:        "- 5 2",
			expectedInfix: "(5-2)",
			expectError:   false,
		},
		{
			name:          "Simple multiplication",
			prefix:        "* 4 5",
			expectedInfix: "(4*5)",
			expectError:   false,
		},
		{
			name:          "Simple division",
			prefix:        "/ 10 2",
			expectedInfix: "(10/2)",
			expectError:   false,
		},
		{
			name:          "Simple exponentiation",
			prefix:        "^ 2 3",
			expectedInfix: "(2^3)",
			expectError:   false,
		},
		{
			name:          "Complex expression",
			prefix:        "+ 5 * - 4 2 ^ 3 2",
			expectedInfix: "(5+((4-2)*(3^2)))",
			expectError:   false,
		},
		{
			name:          "Invalid expression - not enough operands",
			prefix:        "+ 2",
			expectedInfix: "",
			expectError:   true,
		},
		{
			name:          "Invalid expression - invalid token",
			prefix:        "+ 2 a",
			expectedInfix: "",
			expectError:   true,
		},
		{
			name:          "Empty expression",
			prefix:        "",
			expectedInfix: "",
			expectError:   true,
		},
		{
			name:          "Complex with negative numbers",
			prefix:        "+ -5 * - 4 -2 ^ 3 2",
			expectedInfix: "((-5)+((4-(-2))*(3^2)))",
			expectError:   false,
		},
		{
			name:          "Expression structure from previous debug",
			prefix:        "+ -5 * - 4 -2 ^ 3 2",
			expectedInfix: "",
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualInfix, err := PrefixToInfix(tt.prefix)

			if tt.expectError {
				assert.Error(t, err, fmt.Sprintf("Prefix: %s", tt.prefix))
			} else {
				assert.NoError(t, err, fmt.Sprintf("Prefix: %s", tt.prefix))
				assert.Equal(t, tt.expectedInfix, actualInfix, fmt.Sprintf("Prefix: %s", tt.prefix))
			}
		})
	}
}

func ExamplePrefixToInfix() {
	infix, _ := PrefixToInfix("+ 5 * - 4 2 ^ 3 2")
	fmt.Println(infix)
}

package lab2

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpressionHandler(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
		expectedInfix  string
		expectedError  string
	}{
		{"Valid expression", `{"expression": "+ 3 4"}`, http.StatusOK, "(3+4)", ""},
		{"Invalid expression", `{"expression": "+ 3"}`, http.StatusOK, "", "некоректний вираз"},
		{"Empty expression", `{"expression": ""}`, http.StatusOK, "", "порожній вираз"},
		{"Invalid JSON", `{"expr": "+ 3 4"}`, http.StatusBadRequest, "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/convert", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			expressionHandler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.expectedStatus, res.StatusCode)

			var resp ExpressionResponse
			err := json.NewDecoder(res.Body).Decode(&resp)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedInfix, resp.Infix)
			assert.Equal(t, tt.expectedError, resp.Error)
		})
	}
}

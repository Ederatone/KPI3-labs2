package lab2

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ExpressionRequest struct {
	Expression string `json:"expression"`
}

type ExpressionResponse struct {
	Infix string `json:"infix,omitempty"`
	Error string `json:"error,omitempty"`
}

func expressionHandler(w http.ResponseWriter, r *http.Request) {
	var req ExpressionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	// Якщо вираз порожній (або відсутнє поле "expression" в JSON), повертаємо 400 Bad Request
	if req.Expression == "" {
		w.WriteHeader(http.StatusBadRequest) // Повертаємо 400 Bad Request
		return                               // Без тіла відповіді, як очікує тест
	}

	infix, err := PrefixToInfix(req.Expression)
	resp := ExpressionResponse{Infix: infix}
	if err != nil {
		resp.Error = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

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
		{"Empty expression", `{"expression": ""}`, http.StatusBadRequest, "", ""},        // Змінено статус на 400
		{"Missing Expression Field", `{"expr": "+ 3 4"}`, http.StatusBadRequest, "", ""}, // Перейменовано тест та залишено статус 400
		{"Invalid JSON", `invalid json`, http.StatusBadRequest, "", ""},                  // Додано тест на справді невалідний JSON
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

			// Для статусів 200 та 400 тіло відповіді обробляється по-різному
			if tt.expectedStatus == http.StatusOK {
				var resp ExpressionResponse
				err := json.NewDecoder(res.Body).Decode(&resp)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedInfix, resp.Infix)
				assert.Equal(t, tt.expectedError, resp.Error)
			} else if tt.expectedStatus == http.StatusBadRequest {
				// Для 400 Bad Request тіло має бути порожнім (згідно з вимогою тесту)
				bodyBytes, _ := bytes.ReadAll(res.Body)
				assert.Empty(t, string(bodyBytes), "Body should be empty for 400 status")
			}
		})
	}
}

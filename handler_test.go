package lab2

import (
	"bytes"
	"encoding/json"
	"io"
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
		{"Invalid expression", `{"expression": "+ 3"}`, http.StatusOK, "", "некоректний вираз: недостатньо операндів (1) для оператора '+'"},
		{"Empty expression", `{"expression": ""}`, http.StatusBadRequest, "", "Expression field is required"},
		{"Missing Expression Field", `{"expr": "+ 3 4"}`, http.StatusBadRequest, "", "Expression field is required"},
		{"Invalid JSON", `invalid json`, http.StatusBadRequest, "", "Invalid JSON"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/convert", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			expressionHandler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.expectedStatus, res.StatusCode, "Test case: "+tt.name)

			if tt.expectedStatus == http.StatusOK {
				var resp ExpressionResponse
				err := json.NewDecoder(res.Body).Decode(&resp)
				assert.NoError(t, err, "Test case: "+tt.name+"; decoding OK response body")
				assert.Equal(t, tt.expectedInfix, resp.Infix, "Test case: "+tt.name+"; checking Infix")
				assert.Equal(t, tt.expectedError, resp.Error, "Test case: "+tt.name+"; checking Error")
			} else if tt.expectedStatus == http.StatusBadRequest {
				bodyBytes, err := io.ReadAll(res.Body)
				assert.NoError(t, err, "Test case: "+tt.name+"; reading Bad Request response body")
				if tt.expectedError != "" {
					var respError map[string]string
					err = json.Unmarshal(bodyBytes, &respError)
					assert.NoError(t, err, "Test case: "+tt.name+"; body should be valid JSON for 400 status")
					assert.Equal(t, tt.expectedError, respError["error"], "Test case: "+tt.name+"; Error message in JSON response should match")
				} else {
					assert.Empty(t, string(bodyBytes), "Test case: "+tt.name+"; Body should ideally be empty for 400 status when no specific error expected")
				}
			}
		})
	}
}

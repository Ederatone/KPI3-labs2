package lab2

import (
	"encoding/json"
	"net/http"
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

	// Якщо вираз порожній, повертаємо 400 Bad Request
	if req.Expression == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	infix, err := PrefixToInfix(req.Expression)
	resp := ExpressionResponse{Infix: infix}
	if err != nil {
		resp.Error = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

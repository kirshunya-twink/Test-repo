package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSum(t *testing.T) {
	body := []byte(`{"a":2,"b":3}`)
	req := httptest.NewRequest(http.MethodPost, "/sum", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	Sum(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var resp SumResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Result != 5 {
		t.Fatalf("expected result 5, got %d", resp.Result)
	}
}

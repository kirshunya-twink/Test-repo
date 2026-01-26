package handler

import (
	"bytes"
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

	expected := `{"result":5}`
	if w.Body.String() != expected {
		t.Fatalf("expected %s, got %s", expected, w.Body.String())
	}
}

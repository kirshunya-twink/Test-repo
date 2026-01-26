package e2e

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kirshunya-twink/go-ci-playground/internal/handler"
)

func TestHealthEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	handler.Health(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	if string(body) != "ok" {
		t.Fatalf("expected body 'ok', got '%s'", string(body))
	}
}

func TestSumEndpoint(t *testing.T) {
	// создаём тело запроса с a=2, b=3
	body, _ := json.Marshal(map[string]int{"a": 2, "b": 3})
	req := httptest.NewRequest(http.MethodPost, "/sum", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler.Sum(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var result map[string]int
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}

	expected := 5
	if result["result"] != expected {
		t.Fatalf("expected result %d, got %d", expected, result["result"])
	}
}

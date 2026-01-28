package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kirshunya-twink/go-ci-playground/internal/handler"
)

func TestHealthEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.GET("/health", handler.Health)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var response map[string]string
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}

	if response["status"] != "live" {
		t.Fatalf("expected status 'live', got '%s'", response["status"])
	}
}

func TestSumEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.POST("/sum", handler.Sum)

	// создаём тело запроса с a=2, b=3
	body, _ := json.Marshal(map[string]int{"a": 2, "b": 3})
	req := httptest.NewRequest(http.MethodPost, "/sum", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var result map[string]int
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}

	expected := 5
	if result["result"] != expected {
		t.Fatalf("expected result %d, got %d", expected, result["result"])
	}
}

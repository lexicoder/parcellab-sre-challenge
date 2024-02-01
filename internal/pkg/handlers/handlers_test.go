package handlers

import (
	"net/http"
	"net/http/httptest"
	"parcellab-sre-challenge/internal/pkg/config"
	"parcellab-sre-challenge/internal/pkg/health"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSalutationHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	cfg := config.NewConfig()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SalutationHandler(cfg.Salutation))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, rr.Body.String(), "Hi")
}

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(health.HealthGet())
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Contains(t, rr.Body.String(), "status")
	assert.Contains(t, rr.Body.String(), "timestamp")
	assert.Contains(t, rr.Body.String(), "component")
	assert.Contains(t, rr.Body.String(), "name")
	assert.Contains(t, rr.Body.String(), "version")
}

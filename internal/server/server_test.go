package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	handleHealth(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), `"ok"`) {
		t.Fatalf("expected ok in body, got %s", w.Body.String())
	}
}

func TestTraceEndpoint(t *testing.T) {
	body := `{"name":"simple","object_distance":100,"elements":[{"kind":"thin_lens","focal":50},{"kind":"space","length":200}]}`
	req := httptest.NewRequest(http.MethodPost, "/api/trace", strings.NewReader(body))
	w := httptest.NewRecorder()
	handleTrace(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "report") {
		t.Fatalf("expected report in body, got %s", w.Body.String())
	}
}

func TestTraceInvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/trace", strings.NewReader(`{bad`))
	w := httptest.NewRecorder()
	handleTrace(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestTraceMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/trace", nil)
	w := httptest.NewRecorder()
	handleTrace(w, req)
	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", w.Code)
	}
}

func TestTelescopeEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/telescope?f1=100&f2=-25&spacing=75", nil)
	w := httptest.NewRecorder()
	handleTelescope(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "result") {
		t.Fatalf("expected result in body, got %s", w.Body.String())
	}
}

func TestTelescopeMissingParam(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/telescope?f1=100&f2=-25", nil)
	w := httptest.NewRecorder()
	handleTelescope(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

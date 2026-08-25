package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"abcd-optics/internal/element"
	"abcd-optics/internal/system"
)

type Config struct {
	Addr string
}

func ListenAndServe(cfg Config) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/trace", handleTrace)
	mux.HandleFunc("/api/telescope", handleTelescope)
	mux.HandleFunc("/health", handleHealth)
	return http.ListenAndServe(cfg.Addr, mux)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}

func handleTrace(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "POST required")
		return
	}
	body, err := io.ReadAll(io.LimitReader(r.Body, 1<<20))
	if err != nil {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("read body: %v", err))
		return
	}
	spec, err := element.ParseSpec(body)
	if err != nil {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("parse spec: %v", err))
		return
	}
	report, err := system.BuildReport(spec)
	if err != nil {
		writeError(w, http.StatusUnprocessableEntity, fmt.Sprintf("build report: %v", err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"report": report.String(),
	})
}

func handleTelescope(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "GET or POST required")
		return
	}
	q := r.URL.Query()
	f1, err := strconv.ParseFloat(q.Get("f1"), 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("f1: %v", err))
		return
	}
	f2, err := strconv.ParseFloat(q.Get("f2"), 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("f2: %v", err))
		return
	}
	spacing, err := strconv.ParseFloat(q.Get("spacing"), 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("spacing: %v", err))
		return
	}
	t, err := system.AnalyzeTelescope(f1, f2, spacing)
	if err != nil {
		writeError(w, http.StatusUnprocessableEntity, fmt.Sprintf("telescope: %v", err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": t.String(),
	})
}

func writeError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

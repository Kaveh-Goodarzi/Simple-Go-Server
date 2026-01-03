package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed!", http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("OK"))
	})

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	if w.Body.String() != "OK" {
		t.Fatalf("expected body OK, got %s", w.Body.String())
	}
}

func TestEchoEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/echo?msg=hello", nil)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := r.URL.Query().Get("msg")
		if msg == "" {
			http.Error(w, "missing msg parameter", http.StatusBadRequest)
			return
		}
		w.Write([]byte(msg))
	})

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	if w.Body.String() != "hello" {
		t.Fatalf("expected hello, got %s", w.Body.String())
	}
}

func TestEchoMissingParam(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/echo", nil)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := r.URL.Query().Get("msg")
		if msg == "" {
			http.Error(w, "missing msg parameter", http.StatusBadRequest)
			return
		}
	})

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

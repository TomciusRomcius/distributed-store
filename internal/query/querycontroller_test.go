package query

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func QueryControllerSetShouldInsertANewKeyValuePair(t *testing.T) {
	store := NewKeyvalStore(NewFileDataStorage(), NewMemtable(), &QueryParser{})
	handler := QueryController(store)
	req := httptest.NewRequest(
		http.MethodPost,
		"/query",
		strings.NewReader(`{"query":"set path = stringvalue"}`),
	)

	rec := httptest.NewRecorder()

	handler(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}
}

func QueryControllerDelShouldRemoveAKeyValuePair(t *testing.T) {
	store := NewKeyvalStore(NewFileDataStorage(), NewMemtable(), &QueryParser{})
	handler := QueryController(store)
	req := httptest.NewRequest(
		http.MethodPost,
		"/query",
		strings.NewReader(`{"query":"del path"}`),
	)

	rec := httptest.NewRecorder()

	handler(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}
}

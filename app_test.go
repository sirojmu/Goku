package goku

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouteWithParameter(t *testing.T) {
	app := New()
	app.Get("/produk/:id", func(context *Context) error {
		return context.JSON(http.StatusOK, map[string]string{"id": context.Param("id")})
	})

	response := httptest.NewRecorder()
	app.ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/produk/42", nil))

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, ingin %d", response.Code, http.StatusOK)
	}
	if response.Body.String() != "{\"id\":\"42\"}\n" {
		t.Fatalf("respons tidak sesuai: %s", response.Body.String())
	}
}

func TestUnknownRouteReturnsNotFound(t *testing.T) {
	response := httptest.NewRecorder()
	New().ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/tidak-ada", nil))
	if response.Code != http.StatusNotFound {
		t.Fatalf("status = %d, ingin %d", response.Code, http.StatusNotFound)
	}
}

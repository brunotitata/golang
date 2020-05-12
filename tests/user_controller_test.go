package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	_ "github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func TestCriarNovoUsuario(t *testing.T) {
	router := gin.Default()
	w := performRequest(router, "GET", "usuarios")

	assert.Equal(t, http.StatusOK, w.Code)
}

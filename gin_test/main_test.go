package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/user/:id", GetUser)
	return router
}

func TestUserRoute(t *testing.T) {
	router := setupRouter()

	t.Run("Valid user ID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/user/123", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		expectedBody := `{"id":"123"}`
		assert.JSONEq(t, expectedBody, w.Body.String())
	})

	t.Run("Invalid Method", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/user/123", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

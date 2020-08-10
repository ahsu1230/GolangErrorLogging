package services_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahsu1230/golangwebservertutorial/src/router"
	
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	engine := router.Setup()
	return engine
}

func TestNoRoute(t *testing.T) {
	router := setupRouter()
	
	// Record HTTP response
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "asdf", nil)
	router.ServeHTTP(w, req)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, w.Code)
}

func TestPing1(t *testing.T) {
	router := setupRouter()

	// Record HTTP response
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/ping1", nil)
	router.ServeHTTP(w, req)

	// Validate results
	assert.EqualValues(t, http.StatusOK, w.Code)
}

func TestPong3(t *testing.T) {
	router := setupRouter()

	// Record HTTP response
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/pong3", nil)
	router.ServeHTTP(w, req)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, w.Code)
}
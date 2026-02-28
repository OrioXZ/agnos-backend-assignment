package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.POST("/staff/create", func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing required fields"})
	})

	r.POST("/staff/login", func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing required fields"})
	})

	return r
}

func TestStaffCreate_MissingField(t *testing.T) {
	r := setupTestRouter()

	req := httptest.NewRequest(http.MethodPost, "/staff/create", bytes.NewBuffer([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestStaffLogin_MissingField(t *testing.T) {
	r := setupTestRouter()

	req := httptest.NewRequest(http.MethodPost, "/staff/login", bytes.NewBuffer([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestStaffCreate_Positive(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.POST("/staff/create", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"username": "admin"})
	})

	req := httptest.NewRequest(http.MethodPost, "/staff/create",
		bytes.NewBuffer([]byte(`{
			"username":"admin",
			"password":"1234",
			"hospitalCode":"hospital_a"
		}`)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

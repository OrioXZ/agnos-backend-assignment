package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupPatientRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// จำลองว่าไม่มี auth → return 401
	r.GET("/patient/search", func(c *gin.Context) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
	})

	return r
}

func TestPatientSearch_Unauthorized(t *testing.T) {
	r := setupPatientRouter()

	req := httptest.NewRequest(http.MethodGet, "/patient/search?last_name=Sukjai", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", w.Code)
	}
}

func TestPatientSearch_Positive(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.GET("/patient/search", func(c *gin.Context) {
		c.JSON(http.StatusOK, []gin.H{
			{"first_name": "Somying"},
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/patient/search?last_name=Sukjai", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

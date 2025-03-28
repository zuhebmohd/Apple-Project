package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"web-app/routes"

	"github.com/gin-gonic/gin"
)

func TestGetClusters(t *testing.T) {
	router := gin.Default()
	routes.SetupRoutes(router)

	req, _ := http.NewRequest("GET", "/api/clusters", nil)
	// Add basic auth header for an authorized user (e.g., readonly)
	req.SetBasicAuth("readonly", "readonly")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

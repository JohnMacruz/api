package apireq

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPostMethod(t *testing.T) {

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/", PostMethod)

	req, err := http.NewRequest(http.MethodPost, "/", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestGetMethod(t *testing.T) {

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/", GetMethod)

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"broker-backend/controllers"
	"broker-backend/database"
	"broker-backend/utils"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	database.ConnectDB()
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	private := r.Group("/")
	private.Use(func(c *gin.Context) {
		token, _ := utils.GenerateToken("testuser")
		c.Request.Header.Set("Authorization", "Bearer "+token)
		c.Next()
	})

	private.GET("/holdings", controllers.GetHoldings)
	private.GET("/orderbook", controllers.GetOrderbook)
	private.GET("/positions", controllers.GetPositions)

	return r
}

func TestHealthEndpoint(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestSignUp(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	payload := `{"username":"testuser","password":"password123"}`
	req, _ := http.NewRequest("POST", "/signup", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestLogin(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	payload := `{"username":"testuser","password":"password123"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestHoldingsEndpoint(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/holdings", nil)

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestOrderbookEndpoint(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/orderbook", nil)

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestPositionsEndpoint(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/positions", nil)

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

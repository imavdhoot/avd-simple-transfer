package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gorm.io/gorm"
	"github.com/gin-gonic/gin"

	"github.com/imavdhoot/avd-simple-transfer/config"
	"github.com/imavdhoot/avd-simple-transfer/src/handler"
	"github.com/imavdhoot/avd-simple-transfer/src/model"
  "github.com/imavdhoot/avd-simple-transfer/src/middleware"
)

func setupTestDB2(t *testing.T) *gorm.DB {
	os.Setenv("GO_ENV", "test")
  os.Setenv("PG_DB", "transfers_test")  
	db := config.ConnectDB()
	sqlDB, _ := db.DB()
	sqlDB.Exec("TRUNCATE accounts, transactions RESTART IDENTITY")
	db.Create(&model.Account{AccountID: 201, Balance: 500.0})
	return db
}

func setupRouter2(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	h := handler.New(db)
  r.Use(gin.Logger(), gin.Recovery(), middleware.ErrorHandler())
	r.GET("/api/v1/accounts/:account_id", h.GetAccount)
	return r
}

func TestGetAccount(t *testing.T) {
	db := setupTestDB2(t)
	router := setupRouter2(db)

	tests := []struct {
		name       string
		url        string
		wantStatus int
	}{
		{"existing account", "/api/v1/accounts/201", 200},
		{"non-existent account", "/api/v1/accounts/999", 404},
		{"invalid ID", "/api/v1/accounts/abc", 400},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, tt.url, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("got %v, want %v", w.Code, tt.wantStatus)
			}
		})
	}
}

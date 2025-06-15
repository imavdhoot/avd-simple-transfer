package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gorm.io/gorm"
	"github.com/gin-gonic/gin"

	"github.com/imavdhoot/avd-simple-transfer/config"
	"github.com/imavdhoot/avd-simple-transfer/src/handler"
  "github.com/imavdhoot/avd-simple-transfer/src/middleware"
)

func setupTestDB(t *testing.T) *gorm.DB {
	os.Setenv("GO_ENV", "test")
  os.Setenv("PG_DB", "transfers_test")
	db := config.ConnectDB()
	sqlDB, _ := db.DB() 
	sqlDB.Exec("TRUNCATE accounts, transactions RESTART IDENTITY")
	return db
}

func setupRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	h := handler.New(db)

  r.Use(gin.Logger(), gin.Recovery(), middleware.ErrorHandler())
	r.POST("/api/v1/accounts", h.CreateAccount)
	return r
}

func TestCreateAccount(t *testing.T) {
	db := setupTestDB(t)
	router := setupRouter(db)

	tests := []struct {
		name       string
		payload    map[string]interface{}
		wantStatus int
	}{
		{"valid account", map[string]interface{}{"account_id": 101, "initial_balance": 100.5}, 201},
		{"missing account_id", map[string]interface{}{"initial_balance": 100.5}, 400},
		{"missing balance", map[string]interface{}{"account_id": 102}, 400},
		{"negative balance", map[string]interface{}{"account_id": 103, "initial_balance": -50}, 400},
		{"duplicate account", map[string]interface{}{"account_id": 101, "initial_balance": 10}, 409},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest(http.MethodPost, "/api/v1/accounts", bytes.NewBuffer(b))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("got %v, want %v", w.Code, tt.wantStatus)
			}
		})
	}
}

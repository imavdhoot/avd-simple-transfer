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
	"github.com/imavdhoot/avd-simple-transfer/src/model"
  "github.com/imavdhoot/avd-simple-transfer/src/middleware"
)

func setupTestDB3(t *testing.T) *gorm.DB {
	os.Setenv("GO_ENV", "test")
  os.Setenv("PG_DB", "transfers_test")
	db := config.ConnectDB()
	sqlDB, _ := db.DB()
	sqlDB.Exec("TRUNCATE accounts, transactions RESTART IDENTITY")
	db.Create(&model.Account{AccountID: 301, Balance: 200.0})
	db.Create(&model.Account{AccountID: 302, Balance: 50.0})
	return db
}

func setupRouter3(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	h := handler.New(db)

  r.Use(gin.Logger(), gin.Recovery(), middleware.ErrorHandler())
	r.POST("/api/v1/transactions", h.SubmitTransaction)
	return r
}

func TestSubmitTransaction(t *testing.T) {
	db := setupTestDB3(t)
	router := setupRouter3(db)

	tests := []struct {
		name       string
		payload    map[string]interface{}
		wantStatus int
	}{
		{"valid transaction", map[string]interface{}{"source_account_id": 301, "destination_account_id": 302, "amount": 50}, 200},
		{"insufficient funds", map[string]interface{}{"source_account_id": 302, "destination_account_id": 301, "amount": 999}, 406},
		{"missing source", map[string]interface{}{"destination_account_id": 301, "amount": 50}, 400},
		{"invalid amount", map[string]interface{}{"source_account_id": 301, "destination_account_id": 302, "amount": -100}, 400},
		{"non-existent dest", map[string]interface{}{"source_account_id": 301, "destination_account_id": 999, "amount": 10}, 404},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest(http.MethodPost, "/api/v1/transactions", bytes.NewBuffer(b))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("got %v, want %v", w.Code, tt.wantStatus)
			}
		})
	}
}

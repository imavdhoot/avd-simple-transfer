package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/imavdhoot/avd-simple-transfer/config"
	"github.com/imavdhoot/avd-simple-transfer/src/handler"
	"github.com/imavdhoot/avd-simple-transfer/src/middleware"
)

func main() {
	db := config.ConnectDB()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("db.DB: %v ", err)
	}
	defer sqlDB.Close()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.ErrorHandler())

	// Simple Health check route (no auth, no DB)
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})	

	h := handler.New(db)

	// Version-1 routes
	v1 := r.Group("/api/v1")
	v1.POST("/accounts", h.CreateAccount)
	v1.GET("/accounts/:account_id", h.GetAccount)
	v1.POST("/transactions", h.SubmitTransaction)

	log.Println("🚀  Listening on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("server error:", err)
	}
}

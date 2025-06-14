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

	h := handler.New(db)
	r.POST("/accounts", h.CreateAccount)
	r.GET("/accounts/:id", h.GetAccount)
	r.POST("/transactions", h.SubmitTransaction)

	log.Println("🚀  Listening on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("server error:", err)
	}
}

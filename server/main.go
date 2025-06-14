package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/imavdhoot/avd-simple-transfer/config"
	"github.com/imavdhoot/avd-simple-transfer/src/handler"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	r := gin.Default()

	h := handler.New(db)
	r.POST("/accounts", h.CreateAccount)
	r.GET("/accounts/:id", h.GetAccount)
	r.POST("/transactions", h.SubmitTransaction)

	log.Println("🚀  Listening on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("server error:", err)
	}
}

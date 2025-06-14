package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imavdhoot/avd-simple-transfer/src/service"
)

type TransactionHandler struct {
	svc *service.TransactionService
}

func (h *TransactionHandler) SubmitTransaction(c *gin.Context) {
	var req struct {
		SourceAccountID      int64   `json:"source_account_id"`
		DestinationAccountID int64   `json:"destination_account_id"`
		Amount               float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.svc.Transfer(c, req.SourceAccountID, req.DestinationAccountID, req.Amount)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

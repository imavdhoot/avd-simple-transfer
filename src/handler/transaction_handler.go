package handler

import (
	"log"
	"time"
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/imavdhoot/avd-simple-transfer/src/dto"
	"github.com/imavdhoot/avd-simple-transfer/src/service"
	"github.com/imavdhoot/avd-simple-transfer/src/utils"
)

type TransactionHandler struct {
	svc *service.TransactionService
}

func (h *TransactionHandler) SubmitTransaction(c *gin.Context) {
	rid := c.GetString("request_id")
	log.Printf("[RID=%s][HandlerSubmitTransaction] request received", rid)

	var req dto.SubmitTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			log.Printf("[RID=%s][HandlerSubmitTransaction] Validation error %+v", rid, ve)
			c.JSON(http.StatusBadRequest, utils.NewValidationResp(c, ve))
			return
		}		

		c.Error(err)
		return
	}
	
	log.Printf("[RID=%s][HandlerSubmitTransaction] request body %+v", rid, req)

	txn, err := h.svc.Transfer(c, req.SourceAccountID, req.DestinationAccountID, req.Amount)
	if err != nil {
		c.Error(err)
		return
	}

	log.Printf("[RID=%s][HandlerSubmitTransaction]txn body %+v", rid, txn)

	resp := dto.SubmitTransactionResponse{
		TransactionID: txn.ID,
		Message:       "success",
		Status:        http.StatusOK,
		CreatedAt:     txn.CreatedAt.Format(time.RFC3339),
		RequestID:     rid,
	}

	log.Printf("[RID=%s][HandlerSubmitTransaction] response body %+v", rid, resp)
	c.JSON(http.StatusOK, resp)
}

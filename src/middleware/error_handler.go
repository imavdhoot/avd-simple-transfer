package middleware

import (
	"log"
	"errors"
	"context"
	"net/http"
	"strconv"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/go-playground/validator/v10"
	"github.com/imavdhoot/avd-simple-transfer/src/constant"
)

type ErrorResponse struct {
	Error     string      `json:"error"`
	Code      string      `json:"code,omitempty"`
	Status    int         `json:"status"`
	RequestID string      `json:"request_id,omitempty"`
	Details   interface{} `json:"details,omitempty"`
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		// Step 1: Set a request ID for traceability
		rid := uuid.New().String()
		c.Set("request_id", rid)

		ctx := context.WithValue(c.Request.Context(), "request_id", rid)
		c.Request = c.Request.WithContext(ctx)

		c.Next()

		// Step 3: Handle errors from c.Error()
		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		status := http.StatusInternalServerError
		code := "INTERNAL_ERROR"

		log.Printf("[DEBUG] switch duplicate-test: %T  ->  %v", err, err)
		switch {
		case errors.Is(err, constant.ErrAccountExists):
			code, status = "ACCOUNT_EXIST", http.StatusConflict

		case errors.Is(err, constant.ErrAccountNotFound):
			code, status = "ACCOUNT_NOT_FOUND", http.StatusNotFound

		case errors.Is(err, constant.ErrInsufficientFund):
			code, status = "INSUFFICIENT_FUNDS", http.StatusNotAcceptable

		case errors.Is(err, constant.ErrSrcAccountNotFound):
			code, status = "SRC_ACCOUNT_NOT_FOUND", http.StatusNotFound

		case errors.Is(err, constant.ErrDstAccountNotFound):
			code, status = "DST_ACCOUNT_NOT_FOUND", http.StatusNotFound

		case errors.As(err, new(validator.ValidationErrors)),
				 errors.As(err, new(*json.SyntaxError)),
				 errors.As(err, new(*json.UnmarshalTypeError)),
				 errors.As(err, new(*strconv.NumError)):
			code, status = "INVALID_REQUEST", http.StatusBadRequest
		}		

		resp := ErrorResponse{
			Error:     err.Error(),
			Code:      code,
			Status:    status,
			RequestID: rid,
		}

		log.Printf("[RID=%s][ErrorHandler] response body %+v", rid, resp)
		c.JSON(status, resp)
	}
}

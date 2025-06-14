package middleware

import (
	"errors"
	"net/http"
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

		c.Next() // run handler

		// Step 2: Handle errors from c.Error()
		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		status := http.StatusInternalServerError
		code := "INTERNAL_ERROR"

		switch {
		case errors.Is(err, constant.ErrAccountNotFound):
			code = "ACCOUNT_NOT_FOUND"
			status = http.StatusNotFound
		case errors.Is(err, constant.ErrInsufficientFund):
			code = "INSUFFICIENT_FUNDS"
			status = http.StatusConflict
		case errors.As(err, new(validator.ValidationErrors)),
				 errors.As(err, new(*json.SyntaxError)),
				 errors.As(err, new(*json.UnmarshalTypeError)):
				code =  "INVALID_REQUEST"
				status = http.StatusBadRequest
		}

		resp := ErrorResponse{
			Error:     err.Error(),
			Code:      code,
			Status:    status,
			RequestID: rid,
		}

		c.JSON(status, resp)
	}
}

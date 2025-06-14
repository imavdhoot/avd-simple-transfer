package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/imavdhoot/avd-simple-transfer/src/middleware"
)

func NewValidationResp(c *gin.Context, ve validator.ValidationErrors) middleware.ErrorResponse {
	details := map[string]string{}
	for _, fe := range ve {
		details[fe.Field()] = msgForTag(fe)
	}
	return middleware.ErrorResponse{
		Error:     "Validation failed",
		Code:      "INVALID_REQUEST",
		Status:    http.StatusBadRequest,
		RequestID: c.GetString("request_id"),
		Details:   details,
	}
}

// msgForTag maps validator tags → human text.
func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "is required"
	case "gt":
		return fmt.Sprintf("must be greater than %s", fe.Param())
	default:
		return fmt.Sprintf("failed on '%s'", fe.Tag())
	}
}

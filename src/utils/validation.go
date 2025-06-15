package utils

import (
	"log"
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

	rid := c.GetString("request_id")

	validationResp := middleware.ErrorResponse{
		Error:     "Validation failed",
		Code:      "INVALID_REQUEST",
		Status:    http.StatusBadRequest,
		RequestID: rid,
		Details:   details,
	}

	log.Printf("[RID=%s][NewValidationResp] response body %+v", rid, validationResp)
	return validationResp
}

// msgForTag maps validator tags to human readable errors
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

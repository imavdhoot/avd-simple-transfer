package handler

import (
	"log"
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/imavdhoot/avd-simple-transfer/src/dto"
	"github.com/imavdhoot/avd-simple-transfer/src/model"
	"github.com/imavdhoot/avd-simple-transfer/src/service"
	"github.com/imavdhoot/avd-simple-transfer/src/utils"
)

type AccountHandler struct {
	svc *service.AccountService
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	rid := c.GetString("request_id")
	log.Printf("[RID=%s][HandlerCreateAccount] request received", rid)

	var req dto.CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			log.Printf("[RID=%s][HandlerCreateAccount] Validation error %+v", rid, ve)
			c.JSON(http.StatusBadRequest, utils.NewValidationResp(c, ve))
			return
		}

		c.Error(err)
		return
	}
	
	log.Printf("[RID=%s][HandlerCreateAccount] request body %+v", rid, req)
	account := model.Account{AccountID: req.AccountID, Balance: req.InitialBalance}
	
	err := h.svc.Create(c, account)
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusCreated)
}

func (h *AccountHandler) GetAccount(c *gin.Context) {
	rid := c.GetString("request_id")
	idStr := c.Param("account_id")
	log.Printf("[RID=%s][HandlerGetAccount] request received AccountID:: %s", rid, idStr)

	var uri dto.AccountURI
	if err := c.ShouldBindUri(&uri); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			c.JSON(http.StatusBadRequest, utils.NewValidationResp(c, ve))
			return
		}
		c.Error(err);
		return
	}
	
	acc, err := h.svc.Get(c, uri.AccountID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, acc)
}

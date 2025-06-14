package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imavdhoot/avd-simple-transfer/src/model"
	"github.com/imavdhoot/avd-simple-transfer/src/service"
)

type AccountHandler struct {
	svc *service.AccountService
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req struct {
		AccountID      int64   `json:"account_id"`
		InitialBalance float64 `json:"initial_balance"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.svc.Create(c, model.Account{AccountID: req.AccountID, Balance: req.InitialBalance})
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (h *AccountHandler) GetAccount(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	acc, err := h.svc.Get(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, acc)
}

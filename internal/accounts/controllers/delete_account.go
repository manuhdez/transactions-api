package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
)

type DeleteAccountController struct {
	service service.DeleteAccountService
}

func NewDeleteAccountController(s service.DeleteAccountService) DeleteAccountController {
	return DeleteAccountController{s}
}

func (c DeleteAccountController) Handle(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.Delete(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

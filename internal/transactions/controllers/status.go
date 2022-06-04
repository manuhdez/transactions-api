package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusController struct{}

func NewStatusController() StatusController {
	return StatusController{}
}

type statusResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

func (StatusController) Handle(ctx *gin.Context) {
	status := statusResponse{"ok", "transactions"}
	ctx.JSON(http.StatusOK, status)
}

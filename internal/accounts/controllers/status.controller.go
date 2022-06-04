package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type statusResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

type StatusController struct{}

func NewStatusController() StatusController {
	return StatusController{}
}

func (c StatusController) Handle(ctx *gin.Context) {
	r := statusResponse{"ok", "accounts service"}
	ctx.JSON(http.StatusOK, r)
}

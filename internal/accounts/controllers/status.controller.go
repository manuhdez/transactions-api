package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type statusResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

func StatusController(ctx *gin.Context) {
	r := statusResponse{"ok", "accounts service"}
	ctx.JSON(http.StatusOK, r)
}

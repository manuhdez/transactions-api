package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type statusResponse struct {
	Status string `json:"status"`
}

func StatusController(ctx *gin.Context) {
	r := statusResponse{"ok"}
	ctx.JSON(http.StatusOK, r)
}

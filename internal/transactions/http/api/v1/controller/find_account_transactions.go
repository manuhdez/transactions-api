package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FindAccountTransactions struct {
	// repository transaction.Repository
}

func NewFindAccountTransactions() FindAccountTransactions {
	// return FindAccountTransactions{repository}
	return FindAccountTransactions{}
}

type jsonTransaction struct {
	Type     string  `json:"type"`
	Amount   float32 `json:"amount"`
	Currency string  `json:"currency"`
}

type response struct {
	Transactions []jsonTransaction `json:"transactions"`
}

func (ctlr *FindAccountTransactions) Handle(ctx *gin.Context) {
	// accountId := ctx.Param("id")

	res := response{[]jsonTransaction{}}
	ctx.JSON(http.StatusOK, res)
}

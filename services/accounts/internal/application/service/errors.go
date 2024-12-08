package service

import "errors"

var (
	ErrInvalidTransactionType  = errors.New("invalid transaction type")
	ErrUnauthorizedTransaction = errors.New("unauthorized transaction")
)

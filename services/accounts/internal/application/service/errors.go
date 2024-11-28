package service

import "errors"

var (
	ErrInvalidTransactionType = errors.New("invalid transaction type")
	ErrUnauthorized           = errors.New("unauthorized transaction")
)

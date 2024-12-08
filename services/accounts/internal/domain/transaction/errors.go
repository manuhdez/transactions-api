package transaction

import "errors"

var (
	ErrTransferAmountTooBig   = errors.New("transfer amount too big")
	ErrTransferAmountTooSmall = errors.New("transfer amount too small")
	ErrInsufficientBalance    = errors.New("insufficient balance")
)

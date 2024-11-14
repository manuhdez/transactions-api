package request

import (
	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
	domainshared "github.com/manuhdez/transactions-api/shared/domain"
)

type CreateAccount struct {
	Id       string  `json:"id" validate:"required"`
	UserId   string  `json:"user_id" validate:"required"`
	Balance  float32 `json:"balance" default:"0"`
	Currency string  `json:"currency" default:"EUR"`
}

func (req CreateAccount) Decode() account.Account {
	return account.NewWithUserID(req.Id, domainshared.NewID(req.UserId), req.Balance, req.Currency)
}

package service

import (
	"context"

	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
)

type DeleteAccountService struct {
	repository account.Repository
}

func NewDeleteAccountService(r account.Repository) DeleteAccountService {
	return DeleteAccountService{r}
}

func (s DeleteAccountService) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}

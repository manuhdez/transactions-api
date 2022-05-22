package service

import (
	"context"

	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
)

type FindAllService struct {
	repository account.Repository
}

func NewFindAllService(r account.Repository) FindAllService {
	return FindAllService{r}
}

func (s FindAllService) Find(ctx context.Context) ([]account.Account, error) {
	return s.repository.FindAll(ctx)
}

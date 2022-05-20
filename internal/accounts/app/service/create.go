package service

import "github.com/manuhdez/transactions-api/internal/accounts/domain/account"

type CreateService struct {
	repository account.Repository
}

func NewCreateService(repository account.Repository) CreateService {
	return CreateService{repository}
}

func (s CreateService) Create(a account.Account) error {
	return s.repository.Create(a)
}

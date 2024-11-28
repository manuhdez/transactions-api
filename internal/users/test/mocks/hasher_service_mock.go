package mocks

import (
	"fmt"

	"github.com/manuhdez/transactions-api/internal/users/internal/domain/service"
)

type HasherService struct {
	domain_service.HashService
	HashError    error
	CompareError error
}

func (h HasherService) Hash(password string) (string, error) {
	return password, h.HashError
}

func (h HasherService) Compare(hashed, password string) error {
	if h.CompareError != nil {
		return h.CompareError
	}

	if hashed != password {
		return fmt.Errorf("invalid password: got %s want %s", hashed, password)
	}

	return nil
}

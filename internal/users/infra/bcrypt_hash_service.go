package infra

import "golang.org/x/crypto/bcrypt"

const (
	hashCost = bcrypt.DefaultCost
)

type BcryptHashService struct {
}

func NewBcryptService() BcryptHashService {
	return BcryptHashService{}
}

func (b BcryptHashService) Hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	return string(hashed), err
}

func (b BcryptHashService) Compare(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}

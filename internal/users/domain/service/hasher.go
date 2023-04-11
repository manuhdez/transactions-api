package service

type HashService interface {
	Hash(value string) (string, error)
	Compare(hashed, password string) error
}

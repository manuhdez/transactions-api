package infra

import (
	"encoding/json"

	"github.com/manuhdez/transactions-api/internal/users/domain/user"
)

type UserJson struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func NewUserJson(u user.User) UserJson {
	return UserJson{u.Id, u.FirstName, u.LastName, u.Email}
}

func (u UserJson) ToJson() ([]byte, error) {
	str, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	return str, nil
}

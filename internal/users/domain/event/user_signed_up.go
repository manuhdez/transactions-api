package event

import "encoding/json"

const UserSignedUpType Type = "event.users.user_signed_up"

type UserSignedUp struct {
	body UserSignedUpBody
}

type UserSignedUpBody struct {
	Id        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func NewUserSignedUp(data UserSignedUpBody) *UserSignedUp {
	return &UserSignedUp{body: data}
}

func (e UserSignedUp) Type() Type {
	return UserSignedUpType
}

func (e UserSignedUp) Body() []byte {
	body, err := json.Marshal(e.body)
	if err != nil {
		return nil
	}
	return body
}

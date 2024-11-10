package event

const userSignedUpType = "event.users.user_signed_up"

type UserSignedUp struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func NewUserSignedUp(id, firstName, lastName, email string) *UserSignedUp {
	return &UserSignedUp{ID: id, FirstName: firstName, LastName: lastName, Email: email}
}

func (u UserSignedUp) Type() string {
	return userSignedUpType
}

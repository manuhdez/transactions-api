package user

type User struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func New(id, first, last, email, password string) User {
	return User{
		Id:        id,
		FirstName: first,
		LastName:  last,
		Email:     email,
		Password:  password,
	}
}

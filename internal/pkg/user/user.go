package user

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Surname  string    `json:"surname"`
}

func New(login, password, name, surname string) *User {
	return &User{
		ID:       uuid.New(),
		Login:    login,
		Password: password,
		Name:     name,
		Surname:  surname,
	}
}

package user

import "github.com/google/uuid"

type User struct {
	ID      uuid.UUID `json:"id"`
	Login   string    `json:"login"`
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
}

func New(login, name, surname string) *User {
	return &User{
		ID:      uuid.New(),
		Login:   login,
		Name:    name,
		Surname: surname,
	}
}

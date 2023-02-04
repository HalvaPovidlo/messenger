package auth

import (
	"crypto/subtle"
	"errors"
	"github.com/HalvaPovidlo/messenger/internal/pkg/user"
)

type service struct {
	cache map[string]*user.User //login >> user
}

func (s *service) Verify(login, password string) (*user.User, bool) {
	u, ok := s.cache[login]
	if !ok {
		return nil, false
	}

	return u, subtle.ConstantTimeCompare([]byte(password), []byte(u.Password)) == 1
}

func (s *service) Register(name, surname, login, password string) error {
	if _, ok := s.cache[login]; ok {
		return errors.New("login already taken")
	}
	s.cache[login] = user.New(login, password, name, surname)
	return nil
}

package auth

import (
	"crypto/subtle"
	"errors"

	"github.com/google/uuid"
)

type credentials struct {
	ID       uuid.UUID `json:"id"`
	Password []byte    `json:"password"`
}

type service struct {
	cache   map[string]credentials // login -> credentials
	storage *storage
}

func New() (*service, error) {
	s := &service{
		storage: NewStorage(),
	}
	all, err := s.storage.GetAll()
	if err != nil {
		return nil, err
	}

	s.cache = make(map[string]credentials, len(all.Credentials))
	for k, v := range all.Credentials {
		s.cache[k] = v
	}

	return s, nil
}

func (s *service) Verify(login, password string) (uuid.UUID, bool) {
	d, ok := s.cache[login]
	if !ok {
		return uuid.Nil, false
	}

	return d.ID, subtle.ConstantTimeCompare([]byte(password), d.Password) == 1
}

func (s *service) Register(name, surname, login, password string) error {
	if _, ok := s.cache[login]; ok {
		return errors.New("login already exists")
	}

	creds := credentials{
		ID:       uuid.New(),
		Password: []byte(password),
	}

	err := s.storage.Add(login, creds)
	if err != nil {
		return err
	}

	s.cache[login] = creds
	return nil
}

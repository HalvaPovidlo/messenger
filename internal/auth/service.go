package auth

import (
	"crypto/subtle"
	"errors"
	"sync"

	"github.com/google/uuid"
)

type credentials struct {
	ID       uuid.UUID `json:"id"`
	Password []byte    `json:"password"`
}

type service struct {
	cache   map[string]credentials // login -> credentials
	cacheMX *sync.RWMutex
	storage *storage
}

func New() (*service, error) {
	s := &service{
		storage: NewStorage(),
		cacheMX: &sync.RWMutex{},
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
	s.cacheMX.RLock()
	d, ok := s.cache[login]
	s.cacheMX.RUnlock()
	if !ok {
		return uuid.Nil, false
	}

	return d.ID, subtle.ConstantTimeCompare([]byte(password), d.Password) == 1
}

func (s *service) Register(name, surname, login, password string) error {
	s.cacheMX.Lock()
	if _, ok := s.cache[login]; ok {
		s.cacheMX.Unlock()
		return errors.New("login already exists")
	}

	creds := credentials{
		ID:       uuid.New(),
		Password: []byte(password),
	}

	s.cache[login] = creds
	s.cacheMX.Unlock()

	err := s.storage.Add(login, creds)
	if err != nil {
		s.cacheMX.Lock()
		delete(s.cache, login)
		s.cacheMX.Unlock()
		return err
	}

	return nil
}

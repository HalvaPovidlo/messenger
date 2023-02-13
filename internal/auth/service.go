package auth

import (
	"crypto/subtle"
	"errors"
	"github.com/HalvaPovidlo/messenger/internal/pkg/user"
	"sync"

	"github.com/google/uuid"
)

type userService interface {
	Add(user user.User) error
}

type credentials struct {
	ID       uuid.UUID `json:"id"`
	Password []byte    `json:"password"`
}

type service struct {
	cache   map[string]credentials // login -> credentials
	cacheMX *sync.RWMutex
	storage *storage
	users   userService
}

func New(users userService) (*service, error) {
	s := &service{
		users:   users,
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

	u := user.New(login, name, surname)
	creds := credentials{
		ID:       u.ID,
		Password: []byte(password),
	}

	s.cache[login] = creds
	s.cacheMX.Unlock()
	err := s.users.Add(*u)
	if err != nil {
		s.cacheMX.Lock()
		delete(s.cache, login)
		s.cacheMX.Unlock()
		return err
	}

	err = s.storage.Add(login, creds)
	if err != nil {
		s.cacheMX.Lock()
		delete(s.cache, login)
		s.cacheMX.Unlock()
		return err
	}

	return nil
}

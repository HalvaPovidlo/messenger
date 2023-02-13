package user

import (
	"github.com/HalvaPovidlo/messenger/internal/pkg/user"
	"sync"
)

type dataBase interface {
	Add(u user.User) error
	GetAll() ([]user.User, error)
}

type service struct {
	storage dataBase
	mx      *sync.RWMutex
	cache   []user.User
}

func New(st dataBase) *service {
	return &service{
		storage: st,
		mx:      &sync.RWMutex{},
		cache:   make([]user.User, 0, 100),
	}
}

func (s *service) Users() ([]user.User, error) {
	s.mx.RLock()
	cache := make([]user.User, len(s.cache))
	copy(cache, s.cache)
	s.mx.RUnlock()

	if len(cache) > 0 {
		return cache, nil
	}
	s.mx.Lock()
	defer s.mx.Unlock()

	users, err := s.storage.GetAll()
	if err != nil {
		return nil, err
	}
	s.cache = users
	return users, nil
}

func (s *service) Add(user user.User) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	err := s.storage.Add(user)
	if err != nil {
		return err
	}
	s.cache = append(s.cache, user)
	return nil
}

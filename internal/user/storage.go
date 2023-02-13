package user

import (
	"encoding/json"
	"errors"
	"github.com/HalvaPovidlo/messenger/internal/pkg/user"
	"io"
	"os"
	"sync"
)

type users struct {
	Data []user.User `json:"data"`
}

type storage struct {
	*sync.Mutex
}

func NewStorage() *storage {
	return &storage{Mutex: &sync.Mutex{}}
}

func (s *storage) GetAll() ([]user.User, error) {
	s.Lock()
	defer s.Unlock()

	file, err := os.Open("users.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users users
	byteValue, _ := io.ReadAll(file)
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		return nil, err
	}
	return users.Data, nil

}

func (s *storage) Add(u user.User) error {
	var newUsers users
	var err error
	newUsers.Data, err = s.GetAll()
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	s.Lock()
	defer s.Unlock()
	if newUsers.Data == nil {
		newUsers.Data = make([]user.User, 0, 100)
	}
	newUsers.Data = append(newUsers.Data, u)

	file, err := json.Marshal(newUsers)
	if err != nil {
		return err
	}

	return os.WriteFile("users.json", file, 0644)
}

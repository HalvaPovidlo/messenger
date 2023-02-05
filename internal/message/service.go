package message

import (
	"strings"
	"sync"

	"github.com/HalvaPovidlo/messenger/internal/pkg/message"
	"github.com/google/uuid"
)

type database interface {
	History(chatID string) ([]message.Message, error)
	Message(msg message.Message, chatID string) error
	Chats(id string) ([]string, error)
}

const averageMessageLength = 11
const uuidLength = 36

type key string

type service struct {
	storage database
	cache   map[key][]message.Message
	cacheMX *sync.RWMutex
}

func New(database database) *service {
	history := make(map[key][]message.Message, 100)
	return &service{
		cache:   history,
		storage: database,
		cacheMX: &sync.RWMutex{},
	}
}

func (s *service) Message(from, to uuid.UUID, text string) error {
	msg := message.Message{ID: from, Text: text}
	key := buildKey(from, to)

	err := s.storage.Message(msg, string(key))
	if err != nil {
		return err
	}

	s.cacheMX.Lock()
	defer s.cacheMX.Unlock()
	v := s.cache[key]
	if v == nil {
		v = make([]message.Message, 0, 100)
	}
	v = append(v, msg)
	s.cache[key] = v

	return nil
}

func (s *service) History(person1, person2 uuid.UUID) (string, error) {
	history, err := s.history(person1, person2)
	if err != nil {
		return "", err
	}

	var response strings.Builder
	response.Grow((uuidLength + averageMessageLength + 3) * len(history))
	for i := 0; i < len(history); i++ {
		id := history[i].ID.String()
		response.WriteString(id)
		response.WriteString(": ")
		response.WriteString(history[i].Text)
		response.WriteString("\n")
	}
	return response.String(), nil
}

func (s *service) history(person1, person2 uuid.UUID) ([]message.Message, error) {
	key := buildKey(person1, person2)
	s.cacheMX.RLock()
	v := s.cache[key]
	s.cacheMX.RUnlock()
	if len(v) == 0 {
		history, err := s.storage.History(string(key))
		if err != nil {
			return nil, err
		}
		s.cacheMX.Lock()
		s.cache[key] = history
		s.cacheMX.Unlock()
		return history, nil
	}
	return v, nil
}

func buildKey(from, to uuid.UUID) key {
	f := from.String()
	t := to.String()
	if f > t {
		return key(f + "_" + t)
	}
	return key(t + "_" + f)
}

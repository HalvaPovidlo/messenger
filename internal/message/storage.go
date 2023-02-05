package message

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"github.com/HalvaPovidlo/messenger/internal/pkg/message"
)

type Messages struct {
	History []message.Message `json:"messages"`
}

type storage struct {
	*sync.Mutex
}

func NewStorage() *storage {
	return &storage{Mutex: &sync.Mutex{}}
}

func (s *storage) History(chatID string) ([]message.Message, error) {
	s.Lock()
	defer s.Unlock()

	file, err := os.Open(chatID + ".json")
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var messages Messages
	byteValue, _ := io.ReadAll(file)
	err = json.Unmarshal(byteValue, &messages)
	if err != nil {
		return nil, err
	}
	return messages.History, nil

}

func (s *storage) Message(msg message.Message, chatID string) error {
	var messages Messages
	var err error
	messages.History, err = s.History(chatID)
	if err != nil {
		return err
	}

	s.Lock()
	defer s.Unlock()
	if messages.History == nil {
		messages.History = make([]message.Message, 0, 100)
	}
	messages.History = append(messages.History, msg)

	file, err := json.Marshal(messages)
	if err != nil {
		return err
	}

	return os.WriteFile(chatID+".json", file, 0644)
}

func (s *storage) Chats(id string) ([]string, error) {
	return nil, nil
}

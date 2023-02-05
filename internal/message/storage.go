package message

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/HalvaPovidlo/messenger/internal/pkg/message"
)

type Messages struct {
	History []message.Message `json:"messages"`
	Bytes   []byte            `json:"bytes"`
}

type storage struct {
}

func NewStorage() *storage {
	return &storage{}
}

func (s *storage) History(chatID string) ([]message.Message, error) {
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

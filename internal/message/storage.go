package message

import (
	"encoding/json"
	"errors"
	"github.com/HalvaPovidlo/messenger/internal/pkg/message"
	"io"
	"os"
)

type Messages struct {
	History []message.Item `json:"messages"`
}

type storage struct {
}

func NewStorage() *storage {
	return &storage{}
}

func (s *storage) History(chatID string) ([]message.Item, error) {

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

func (s *storage) Message(msg message.Item, chatID string) error {
	var messages Messages
	var err error
	messages.History, err = s.History(chatID)
	if err != nil {
		return err
	}
	if messages.History == nil {
		messages.History = make([]message.Item, 0, 100)
	}
	messages.History = append(messages.History, msg)
	file, err := json.Marshal(messages)
	if err != nil {
		return err
	}

	err = os.WriteFile(chatID+".json", file, 0644)
	return err
}
func (s *storage) Chats(id string) ([]string, error) {
	return nil, nil
}

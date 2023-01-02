package message

import (
	"github.com/HalvaPovidlo/messenger/internal/pkg/message"
)

type database interface {
	History(chatID string) ([]message.Item, error)
	Message(msg message.Item, chatID string) error
	Chats(id string) ([]string, error)
}

type key string

type personal struct {
	storage database
	cache   map[key][]message.Item
}

func New() *personal {
	history := make(map[key][]message.Item, 100)
	return &personal{cache: history}
}

func (k *personal) Message(from, to, text string) error {
	msg := message.Item{Name: from, Text: text}
	key := buildKey(from, to)

	err := k.storage.Message(msg, string(key))
	if err != nil {
		return err
	}

	v := k.cache[key]
	if v == nil {
		v = make([]message.Item, 0, 100)
	}
	v = append(v, msg)
	k.cache[key] = v

	return nil
}

func (k *personal) History(person1, person2 string) ([]message.Item, error) {
	key := buildKey(person1, person2)
	v := k.cache[key]
	if len(v) == 0 {
		history, err := k.storage.History(string(key))
		if err != nil {
			return nil, err
		}
		k.cache[key] = history
		return history, nil
	}
	return v, nil
}

func buildKey(from, to string) key {
	if from > to {
		return key(from + "_" + to)
	}
	return key(to + "_" + from)
}

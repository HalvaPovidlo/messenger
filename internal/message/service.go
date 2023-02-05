package message

import (
	"github.com/google/uuid"

	"github.com/HalvaPovidlo/messenger/internal/pkg/message"
)

type database interface {
	History(chatID string) ([]message.Message, error)
	Message(msg message.Message, chatID string) error
	Chats(id string) ([]string, error)
}

type key string

type personal struct {
	storage database
	cache   map[key][]message.Message
}

func New(database2 database) *personal {
	history := make(map[key][]message.Message, 100)
	return &personal{
		cache:   history,
		storage: database2,
	}
}

func (k *personal) Message(from, to uuid.UUID, text string) error {
	msg := message.Message{ID: from, Text: text}
	key := buildKey(from, to)

	err := k.storage.Message(msg, string(key))
	if err != nil {
		return err
	}

	v := k.cache[key]
	if v == nil {
		v = make([]message.Message, 0, 100)
	}
	v = append(v, msg)
	k.cache[key] = v

	return nil
}

func (k *personal) History(person1, person2 uuid.UUID) ([]message.Message, error) {
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

func buildKey(from, to uuid.UUID) key {
	f := from.String()
	t := to.String()
	if f > t {
		return key(f + "_" + t)
	}
	return key(t + "_" + f)
}

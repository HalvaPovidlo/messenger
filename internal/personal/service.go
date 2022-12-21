package personal

import (
	"errors"
	"github.com/HalvaPovidlo/messenger/internal/pkg/message"
)

type key string

type personal struct {
	history map[key][]message.Item
}

func New() *personal {
	history := make(map[key][]message.Item, 100)
	return &personal{history: history}
}

func (k *personal) Message(from, to, msg string) error {
	key := buildKey(from, to)
	v := k.history[key]
	if v == nil {
		v = make([]message.Item, 0, 100)
	}
	v = append(v, message.Item{Name: from, Text: msg})
	k.history[key] = v
	return nil
}

func (k *personal) History(person1, person2 string) ([]message.Item, error) {
	key := buildKey(person1, person2)
	v := k.history[key]
	if len(v) == 0 {
		return nil, errors.New("no such dialogue")
	}
	return v, nil
}

func buildKey(from, to string) key {
	if from > to {
		return key(from + "_" + to)
	}
	return key(to + "_" + from)
}

//func (k *personal) PersonalHistory(c echo.Context) ([]message.History, error) {
//	person1 := c.Param("from")
//	person2 := c.Param("to")
//	key := buildKey(person1, person2)
//	v := k.history[key]
//	if len(v) == 0 {
//		return nil, errors.New("no such dialogue")
//	}
//
//	return c.String(http.StatusOK, v), nil
//}

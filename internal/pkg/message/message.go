package message

import (
	"github.com/google/uuid"
)

type Message struct {
	ID   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}

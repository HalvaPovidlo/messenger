package main

import (
	"github.com/HalvaPovidlo/messenger/internal/api/v1/messages"
	"github.com/HalvaPovidlo/messenger/internal/message"
)

func main() {
	personalService := message.New(message.NewStorage())
	msgService := messages.NewService(personalService)
	msgHandler := messages.NewMessagesHandler(msgService)
	msgHandler.Run("9090")
}

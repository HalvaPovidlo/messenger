package main

import (
	"github.com/HalvaPovidlo/messenger/internal/api/v1/messages"
	"github.com/HalvaPovidlo/messenger/internal/personal"
)

func main() {
	personalService := personal.New()
	msgService := messages.NewService(personalService)
	msgHandler := messages.NewMessagesHandler(msgService)
	msgHandler.Run("9090")
}

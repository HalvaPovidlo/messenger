package main

import (
	"github.com/HalvaPovidlo/messenger/internal/api/v1/messages"
)

func main() {
	personalService := messages.New()
	msgService := messages.NewService(personalService)
	msgHandler := messages.NewMessagesHandler(msgService)
	msgHandler.Run("9090")
}

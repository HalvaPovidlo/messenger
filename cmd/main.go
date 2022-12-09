package main

import "github.com/HalvaPovidlo/messenger/internal/api/v1/messages"

func main() {
	msgService := messages.NewService()
	msgHandler := messages.NewMessagesHandler(msgService)
	msgHandler.Run("9090")
}

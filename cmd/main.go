package main

import (
	"github.com/HalvaPovidlo/messenger/internal/api/v1"
	apiMsg "github.com/HalvaPovidlo/messenger/internal/api/v1/message"
	"github.com/HalvaPovidlo/messenger/internal/message"
)

func main() {
	personalService := message.New(message.NewStorage())
	apiMsgService := apiMsg.NewService(personalService)
	apiMsgHandler := v1.NewMessagesHandler(apiMsgService)
	apiMsgHandler.Run("9090")
}

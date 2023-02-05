package main

import (
	apiv1 "github.com/HalvaPovidlo/messenger/internal/api/v1"
	"github.com/HalvaPovidlo/messenger/internal/auth"
	"github.com/HalvaPovidlo/messenger/internal/message"
)

func main() {
	personalService := message.New(message.NewStorage())
	authService, err := auth.New()
	if err != nil {
		panic(err)
	}

	server := apiv1.NewHandler(personalService, authService)
	server.Run("9090")
}

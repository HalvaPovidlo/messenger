package main

import (
	apiv1 "github.com/HalvaPovidlo/messenger/internal/api/v1"
	"github.com/HalvaPovidlo/messenger/internal/auth"
	"github.com/HalvaPovidlo/messenger/internal/message"
	"github.com/HalvaPovidlo/messenger/internal/user"
)

func main() {
	userService := user.New(user.NewStorage())
	personalService := message.New(message.NewStorage())
	authService, err := auth.New(userService)
	if err != nil {
		panic(err)
	}

	server := apiv1.NewHandler(personalService, authService, userService)
	server.Run("9090")
}

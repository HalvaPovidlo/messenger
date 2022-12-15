package messages

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type messageService interface {
	History(c echo.Context) error
	Message(c echo.Context) error
	PersonalMessage(c echo.Context) error
}

type Handler struct {
	messages messageService
}

func NewMessagesHandler(messages messageService) *Handler {
	return &Handler{
		messages: messages,
	}
}

func (h *Handler) Run(port string) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/msg", h.messages.Message)
	e.GET("/msg", h.messages.History)
	e.GET("/msg/:to", h.messages.PersonalMessage)

	e.Logger.Fatal(e.Start(":" + port))
}

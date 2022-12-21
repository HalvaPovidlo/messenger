package messages

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type messageService interface {
	History(c echo.Context) error
	Message(c echo.Context) error
	PersonalMessage(c echo.Context) error
	PersonalHistory(c echo.Context) error
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

	e.POST("/msg", h.messages.Message)                  //принимает сообщение в общий чат
	e.GET("/msg", h.messages.History)                   //возвращает историю сообщений общего чата
	e.POST("/msg/:to", h.messages.PersonalMessage)      //принимает сообщение в личный чут
	e.GET("/msg/:from/:to", h.messages.PersonalHistory) //возвращает историю личного чата
	e.Logger.Fatal(e.Start(":" + port))
}

package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type registrationHandler interface {
	Register(c echo.Context) error
	Auth(username, password string, c echo.Context) (bool, error)
}

type messageHandler interface {
	History(c echo.Context) error
	Message(c echo.Context) error
	PersonalMessage(c echo.Context) error
	PersonalHistory(c echo.Context) error
}

type Handler struct {
	messages messageHandler

	registration registrationHandler
}

func NewMessagesHandler(messages messageHandler, registration registrationHandler) *Handler {
	handler := &Handler{
		messages:     messages,
		registration: registration,
	}
	return handler
}

func (h *Handler) Run(port string) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BasicAuth(h.registration.Auth))

	e.POST("/msg", h.messages.Message)               //принимает сообщение в общий чат
	e.GET("/msg", h.messages.History)                //возвращает историю сообщений общего чата
	e.POST("/msg/:to", h.messages.PersonalMessage)   //принимает сообщение в личный чут
	e.GET("/msg/:to", h.messages.PersonalHistory)    //возвращает историю личного чата
	e.POST("/registration", h.registration.Register) //региструрет пользователя

	e.Logger.Fatal(e.Start(":" + port))
}

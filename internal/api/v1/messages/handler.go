package messages

import (
	"crypto/subtle"
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
	auth     map[string][]byte
}

func NewMessagesHandler(messages messageService) *Handler {

	return &Handler{
		messages: messages,
		auth:     make(map[string][]byte),
	}
}

func (h *Handler) Run(port string) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BasicAuth(h.Auth))

	e.POST("/msg", h.messages.Message)                  //принимает сообщение в общий чат
	e.GET("/msg", h.messages.History)                   //возвращает историю сообщений общего чата
	e.POST("/msg/:to", h.messages.PersonalMessage)      //принимает сообщение в личный чут
	e.GET("/msg/:from/:to", h.messages.PersonalHistory) //возвращает историю личного чата
	e.Logger.Fatal(e.Start(":" + port))
}

func (h *Handler) Auth(username, password string, c echo.Context) (bool, error) {
	secret, ok := h.auth[username]
	if !ok {
		return false, nil
	}
	if subtle.ConstantTimeCompare([]byte(password), secret) == 1 {
		return true, nil
	}

	return false, nil
}


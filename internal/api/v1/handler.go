package v1

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
	paroli   map[string][]byte
}

func NewMessagesHandler(messages messageService) *Handler {
	maiParoli := make(map[string][]byte)
	maiParoli["maksim"] = []byte("abobus")
	handler := &Handler{
		messages: messages,
		paroli:   maiParoli,
	}
	return handler
}

func (h *Handler) Run(port string) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BasicAuth(h.Auth))

	e.POST("/msg", h.messages.Message)             //принимает сообщение в общий чат
	e.GET("/msg", h.messages.History)              //возвращает историю сообщений общего чата
	e.POST("/msg/:to", h.messages.PersonalMessage) //принимает сообщение в личный чут
	e.GET("/msg/:to", h.messages.PersonalHistory)  //возвращает историю личного чата
	e.Logger.Fatal(e.Start(":" + port))
}

func (h *Handler) Auth(username, password string, c echo.Context) (bool, error) {
	secret, ok := h.paroli[username]
	if !ok {
		return false, nil
	}
	c.Set("username", username)
	if subtle.ConstantTimeCompare([]byte(password), secret) == 1 {
		return true, nil
	}

	return false, nil
}

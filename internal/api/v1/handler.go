package apiv1

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/HalvaPovidlo/messenger/internal/pkg/message"
)

type authService interface {
	Verify(login, password string) (uuid.UUID, bool)
	Register(name, surname, login, password string) error
}

// внутренняя логика обработки сообщений
type messageService interface {
	Message(from, to uuid.UUID, msg string) error
	History(person1, person2 uuid.UUID) ([]message.Message, error)
}

// распаковщик http сообщений
type handler struct {
	messages messageService
	auth     authService
}

func NewHandler(messages messageService, auth authService) *handler {
	return &handler{
		messages: messages,
		auth:     auth,
	}
}

func (h *handler) Run(port string) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/registration", h.Register)                                 // регистрирует пользователя
	e.POST("/msg/:to", h.PersonalMessage, middleware.BasicAuth(h.Auth)) // принимает сообщение в личный чут
	e.GET("/msg/:to", h.PersonalHistory, middleware.BasicAuth(h.Auth))  // возвращает историю личного чата

	e.Logger.Fatal(e.Start(":" + port))
}

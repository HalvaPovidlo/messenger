package apiv1

import (
	"encoding/json"
	"github.com/HalvaPovidlo/messenger/internal/pkg/message"
	"github.com/HalvaPovidlo/messenger/internal/pkg/user"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type userService interface {
	Users() ([]user.User, error)
}

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
	user     userService
}

func NewHandler(messages messageService, auth authService, user userService) *handler {
	return &handler{
		user:     user,
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
	e.GET("/users", h.Users)

	e.Logger.Fatal(e.Start(":" + port))
}

type usersOut struct {
	Users []user.User `json:"users"`
}

func (h *handler) Users(c echo.Context) error {
	users, err := h.user.Users()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	var out usersOut
	out.Users = users
	_, err = json.Marshal(out)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, out)

}

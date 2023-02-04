package message

import (
	"errors"
	"fmt"
	"github.com/HalvaPovidlo/messenger/internal/pkg/message"
	"github.com/HalvaPovidlo/messenger/internal/pkg/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type msgBody struct {
	Text string `json:"text"`
}

// внутренняя логика обработки сообщений
type PersonalService interface {
	Message(from, to, msg string) error
	History(person1, person2 string) ([]message.Item, error)
}

// распаковщик http сообщений
type Messenger struct {
	personal      PersonalService
	commonHistory []message.Item
}

func NewService(p PersonalService) *Messenger {
	return &Messenger{personal: p}

}

func (m *Messenger) PersonalMessage(c echo.Context) error {
	name := c.Param("to")
	from, err := getUser(c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}

	var body msgBody
	if err := c.Bind(&body); err != nil {
		return err
	}
	m.personal.Message(from.ID.String(), name, body.Text)
	fmt.Println(name)
	return c.String(http.StatusOK, "")

}
func (m *Messenger) PersonalHistory(c echo.Context) error {
	var otvet string
	from, err := getUser(c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}
	to := c.Param("to")
	history, err := m.personal.History(from, to)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	for i := 0; i < len(history); i++ {
		otvet += history[i].Name + ": " + history[i].Text + "\n"
	}
	return c.String(http.StatusOK, otvet)
}
func (m *Messenger) History(c echo.Context) error {
	var history string
	for i := 0; i < len(m.commonHistory); i++ {
		history += m.commonHistory[i].Name + ": " + m.commonHistory[i].Text + "\n"
	}

	return c.String(http.StatusOK, history)
}

func (m *Messenger) Message(c echo.Context) error {
	var u message.Item
	if err := c.Bind(&u); err != nil {
		return err
	}

	m.commonHistory = append(m.commonHistory, u)
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) (*user.User, error) {
	v, ok := c.Get("user").(*user.User)
	if ok {
		return v, nil
	}
	return nil, errors.New("username is empty")
}

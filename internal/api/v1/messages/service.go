package messages

import (
	"fmt"
	"github.com/HalvaPovidlo/messenger/internal/pkg/message"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PersonalService interface {
	Message(from, to, msg string) error
	History(person1, person2 string) ([]message.Item, error)
}

type Messenger struct {
	personal      PersonalService
	commonHistory []message.Item
}

func NewService(p PersonalService) *Messenger {
	return &Messenger{personal: p}

}

func (m *Messenger) PersonalMessage(c echo.Context) error {
	name := c.Param("to")

	var body message.Item
	if err := c.Bind(&body); err != nil {
		return err
	}
	m.personal.Message(body.Name, name, body.Text)
	fmt.Println(name)
	return c.String(http.StatusOK, name)

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

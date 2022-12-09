package messages

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Messenger struct {
	history []Message
}

func NewService() *Messenger {
	return &Messenger{}
}

type Message struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

func (m *Messenger) History(c echo.Context) error {
	var history string
	for i := 0; i < len(m.history); i++ {
		history += m.history[i].Name + ": " + m.history[i].Text + "\n"
	}

	return c.String(http.StatusOK, history)
}

func (m *Messenger) Message(c echo.Context) error {
	var u Message
	if err := c.Bind(&u); err != nil {
		return err
	}

	m.history = append(m.history, u)
	return c.JSON(http.StatusCreated, u)
}

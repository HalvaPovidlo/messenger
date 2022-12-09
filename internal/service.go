package internal

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Messenger struct {
	kolvo   int
	slovo   string
	history []Message
}
type Message struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

func (k *Messenger) Hello(c echo.Context) error {
	k.kolvo += 1
	s2 := strconv.Itoa(k.kolvo)
	return c.String(http.StatusOK, "aboba"+s2)
}
func (k *Messenger) Dima(c echo.Context) error {
	k.slovo = "maksim"
	return c.String(http.StatusOK, "andrusha"+k.slovo)
}
func (k *Messenger) History(c echo.Context) error {
	var history string
	for i := 0; i < len(k.history); i++ {
		history += k.history[i].Name + ": " + k.history[i].Text + "\n"
	}
	return c.String(http.StatusOK, history)
}
func (k *Messenger) Message(c echo.Context) error {
	var u Message

	if err := c.Bind(&u); err != nil {
		return err
	}
	k.history = append(k.history, u)
	fmt.Println(u.Text)
	return c.JSON(http.StatusCreated, u)
}

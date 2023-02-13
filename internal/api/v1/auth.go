package apiv1

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
)

const userIDKey = "userID"

type registerBody struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
	Login    string `json:"login"`
}

func (h *handler) Auth(username, password string, c echo.Context) (bool, error) {
	id, ok := h.auth.Verify(strings.ToLower(username), password)
	if !ok {
		return false, nil
	}
	c.Set(userIDKey, id)

	return true, nil
}
func (h *handler) Register(c echo.Context) error {
	var b registerBody
	if err := c.Bind(&b); err != nil {
		return err
	}

	b.Login = strings.ToLower(b.Login)
	r, _ := regexp.Compile("^[a-zA-Z0-9]+(?:.[a-zA-Z0-9]+)*$")
	if !r.Match([]byte(b.Login)) {
		return c.String(http.StatusBadRequest, "Login should contain only these characters: a-z, A-Z, 0-9, .")
	}

	if err := h.auth.Register(b.Name, b.Surname, b.Login, b.Password); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "Registration successful")
}

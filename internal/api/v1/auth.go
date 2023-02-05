package apiv1

import (
	"net/http"

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
	id, ok := h.auth.Verify(username, password)
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

	err := h.auth.Register(b.Name, b.Surname, b.Login, b.Password)
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}
	return c.String(http.StatusOK, "Registration successful")
}

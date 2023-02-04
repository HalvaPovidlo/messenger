package auth

import (
	"github.com/HalvaPovidlo/messenger/internal/pkg/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

type LoginService interface {
	Verify(login, password string) (*user.User, bool)
	Register(name, surname, login, password string) error
}

type Service struct {
	login LoginService
}

type registerBody struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
	Login    string `json:"login"`
}

func (s *Service) Auth(username, password string, c echo.Context) (bool, error) {
	u, ok := s.login.Verify(username, password)
	if !ok {
		return false, nil
	}
	c.Set("user", u)

	return true, nil
}
func (s *Service) Register(c echo.Context) error {
	var b registerBody
	if err := c.Bind(&b); err != nil {
		return err
	}
	err := s.login.Register(b.Name, b.Surname, b.Login, b.Password)
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}
	return c.String(http.StatusOK, "")
}

package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/purwandi/platform/pkg/flaw"
)

// Server ...
type Server struct {
	service *UserService
}

// NewServer ...
func NewServer(service *UserService) *Server {
	return &Server{service: service}
}

// Mount ...
func (s *Server) Mount(e *echo.Echo) {
	e.POST("/auth/register", s.CreateUser)
	e.POST("/auth/login", s.Login)
}

// CreateUser ...
func (s *Server) CreateUser(c echo.Context) error {
	// Given
	inpt := RegisterInput{
		Username: c.FormValue("username"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	// Process
	user, err := s.service.CreateUser(c.Request().Context(), inpt)
	if err != nil {
		if m, err := err.(*flaw.AppError); err {
			return c.JSON(m.StatusCode, m.Message)
		}

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// Login ...
func (s *Server) Login(c echo.Context) error {
	// Given
	inpt := LoginInput{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}

	// Process
	u, err := s.service.Authenticate(c.Request().Context(), inpt)
	if err != nil {
		if m, err := err.(*flaw.AppError); err {
			return c.JSON(m.StatusCode, m.Message)
		}

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set(
		echo.HeaderAuthorization,
		fmt.Sprintf("Bearer %s", u.GetAccessToken()),
	)

	return c.JSON(http.StatusOK, u)
}

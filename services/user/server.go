package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/purwandi/platform/pkg/flaw"
)

// Server ...
type Server struct {
	service *Service
}

// NewServer ...
func NewServer(service *Service) *Server {
	return &Server{service: service}
}

// Mount ...
func (server *Server) Mount(e *echo.Echo) {
	e.POST("/auth/register", server.CreateUser)
}

// CreateUser ...
func (server *Server) CreateUser(c echo.Context) error {
	// Given
	inpt := RegisterInput{
		Username:        c.FormValue("username"),
		Email:           c.FormValue("email"),
		Password:        c.FormValue("password"),
		PasswordConfirm: c.FormValue("password_confirmation"),
	}

	// Process
	user, err := server.service.CreateUser(c.Request().Context(), inpt)
	if err != nil {
		if m, err := err.(*flaw.AppError); err {
			return c.JSON(m.StatusCode, m.Message)
		}

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

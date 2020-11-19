package gateway

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/purwandi/platform/services/user"
)

// Authorization ...
type Authorization struct {
	UserService *user.UserService
}

// NewAuthorizationMiddleware ...
func NewAuthorizationMiddleware(s *user.UserService) *Authorization {
	return &Authorization{UserService: s}
}

// Check for check
func (m Authorization) Check(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		auth := c.Request().Header.Get("Authorization")
		l := len("Bearer")

		if len(auth) < l+1 || auth[:l] != "Bearer" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "missing or malformed jwt",
			})
		}

		// check the current token is exists on storage
		token := auth[l+1:]
		result, err := m.UserService.FindUserByAccessToken(ctx, token)
		if err != nil || result.ID == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "this endpoint requires you to be authenticated.",
			})
		}

		c.SetRequest(c.Request().WithContext(context.WithValue(ctx, "auth", result)))

		return next(c)
	}
}

package middleware

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type Middleware struct{}

// New :
func New() *Middleware {
	return &Middleware{}
}

// Authentication :
func (mw Middleware) Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cronjob := c.Request().Header.Get("X-Authentication")
		if cronjob != os.Getenv("ADMIN_SECRET") {
			return c.JSON(http.StatusUnauthorized, "invalid authorization")
		}
		return next(c)
	}
}

package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// Hello :
func (ctrl Controller) Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

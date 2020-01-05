package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// GetMovieByID : Get movie with ID from DB
func (ctrl Controller) GetMovieByID(c echo.Context) error {

	movieID, err := strconv.Atoi(c.Param("movieID"))
	if err != nil {
		return c.JSON(http.StatusForbidden, errors.New("invalid movie id"))
	}

	movie, err := ctrl.datacon.GetMovieByID(movieID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.Render(http.StatusOK, "index.html", movie)
}

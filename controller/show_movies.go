package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/moviedb-api/controller/transformer"
)

// GetMovieByID : Get movie with ID
func (ctrl Controller) ShowAllMovies(c echo.Context) error {

	movies, err := ctrl.datacon.GetAllMovies()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, transformer.ToMovies(movies))
}

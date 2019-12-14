package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/moviedb-api/model"
)

// GetMovieByID : Get movie with ID
func (ctrl Controller) GetMovieByIDAdmin(c echo.Context) error {
	var i struct {
		ID int
	}

	// bind input value to variable i
	if err := c.Bind(&i); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	tmMovie, err := ctrl.mdbc.GetMovieInfo(i.ID, nil)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	movie := model.Movie{}
	movie.CreatedAt = time.Now().UTC()
	movie.UpdatedAt = time.Now().UTC()
	movie.Title = tmMovie.Title
	movie.Description = tmMovie.Overview
	movie.ID = tmMovie.ID
	movie.DirectLink = tmMovie.Homepage
	ps, err := ctrl.GetDefaultPosterSize()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	movie.FilledUpImageInfo(ctrl.imageConfig.BaseURL,
		ps,
		tmMovie.PosterPath)

	// Save to folder
	go movie.DownloadImage()

	go ctrl.SaveMovieToLocal(movie)

	return c.JSON(http.StatusOK, movie)
}

package controller

import (
	"net/http"
	"time"

	"github.com/moviedb-api/model"
	"github.com/labstack/echo"
)

// GetMovieByID : Get movie with ID from DB
func (ctrl Controller) GetMovieByID(c echo.Context) error {
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

	// Save to memory
	go ctrl.SaveMovieToLocal(movie)

	// Save to folder
	go movie.DownloadImage()

	return c.JSON(http.StatusOK, movie)
}

func (ctrl *Controller) SaveMovieToLocal(m model.Movie) {
	err := ctrl.datacon.SetMovie(m)

	if err != nil {
		panic(err)
	}
	m.IsLocal = true

}

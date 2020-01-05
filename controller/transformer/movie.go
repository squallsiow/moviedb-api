package transformer

import (
	"path"
	"strings"

	"github.com/moviedb-api/controller/response"
	"github.com/moviedb-api/model"
	"github.com/ryanbradynd05/go-tmdb"
)

// ExtractImageName :
func extractImageName(fn string) string {
	return strings.TrimSuffix(strings.TrimPrefix(fn, "/"), extractImageFormat(fn))
}

// ExtractImageFormat :
func extractImageFormat(fn string) string {
	return path.Ext(fn)
}

// ToTMMovies :
func ToTMMovies(i []*tmdb.Movie) []response.Movie {
	movies := make([]response.Movie, 0)
	for _, obj := range i {
		movie := ToTMMovie(obj)
		movies = append(movies, movie)
	}
	return movies
}

// ToTMMovie :
func ToTMMovie(obj *tmdb.Movie) response.Movie {
	movie := response.Movie{}
	movie.ID = obj.ID
	movie.Title = obj.Title
	movie.Description = obj.Overview
	movie.DirectLink = obj.Homepage
	movie.Format = path.Ext(obj.PosterPath)
	movie.ImageName = extractImageName(obj.PosterPath)

	return movie
}

// ToMovies :
func ToMovies(i []*model.Movie) []response.Movie {
	movies := make([]response.Movie, 0)
	for _, obj := range i {
		movie := ToMovie(obj)
		movies = append(movies, movie)
	}
	return movies
}

// ToMovie :
func ToMovie(obj *model.Movie) response.Movie {
	movie := response.Movie{}
	movie.ID = obj.ID
	movie.Title = obj.Title
	movie.Description = obj.Description
	movie.DirectLink = obj.DirectLink
	movie.Format = obj.Format
	movie.ImageName = obj.ImageName

	return movie
}

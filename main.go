package main

import (
	"context"
	"html/template"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/moviedb-api/controller"
	mw "github.com/moviedb-api/controller/middleware"
	"github.com/ryanbradynd05/go-tmdb"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func initTMDB() *tmdb.TMDb {

	config := tmdb.Config{
		APIKey:   "b95d785d64a4e396406586a175e7955c",
		Proxies:  nil,
		UseProxy: false,
	}

	return tmdb.Init(config)
}

func main() {

	ctrl, err := controller.New(initTMDB())
	if err != nil {
		panic(err)
	}

	defer ctrl.CloseDB()

	// set templates
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	// create new echo web server
	api := echo.New()

	api.Renderer = t

	// Middleware
	api.Use(middleware.Logger())
	api.Use(middleware.Secure())
	api.Use(middleware.Recover())

	mid := mw.New()
	// Routes
	admin := api.Group("/admin", mid.Authentication)
	admin.PUT("/movie", ctrl.GetTMBMovieByID)

	// User route
	// api.GET("/", ctrl.Hello)
	api.GET("/showall", ctrl.ShowAllMovies)
	api.GET("/movie/:movieID", ctrl.GetMovieByID)

	// api.Logger.Fatal(api.Start(":8080"))

	// Start server
	go func() {
		if err := api.Start(":8080"); err != nil {
			api.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := api.Shutdown(ctx); err != nil {
		api.Logger.Fatal(err)
	}

}

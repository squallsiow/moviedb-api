package model

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

// DB_MOVIE : DB table name for movie
const DB_MOVIE = "MOVIES"

// Movie : Object type movie
type Movie struct {
	ID               int
	ImageName        string
	Format           string
	Title            string
	Description      string
	DirectLink       string
	ImageOriginalURL string
	ImageLocalPath   string
	IsLocal          bool
	Model
}

// ExtractImageName : Extract image name from original source
func extractImageName(fn string) string {
	return strings.TrimSuffix(strings.TrimPrefix(fn, "/"), extractImageFormat(fn))
}

// ExtractImageFormat : Extract image name from original source
func extractImageFormat(fn string) string {
	return path.Ext(fn)
}

// FilledUpImageInfo : Fill up image url and path
func (m *Movie) FilledUpImageInfo(baseurl string, postersize string, imagepath string) {

	m.ImageOriginalURL = baseurl + postersize + imagepath
	m.ImageName = extractImageName(imagepath)
	m.Format = extractImageFormat(imagepath)

	//open a file for writing
	// pwd, _ := os.Getwd()
	imagename := strconv.Itoa(m.ID) + "_" + m.Title + m.Format
	// imageFullPath := filepath.Join(pwd, os.Getenv("DEFAULT_IMAGE_FOLDER"), imagename)
	imageFullPath := filepath.Join(os.Getenv("DEFAULT_IMAGE_FOLDER"), imagename)
	m.ImageLocalPath = imageFullPath

}

// DownloadImage : Download image from web
func (m *Movie) DownloadImage() error {
	if _, err := os.Stat(m.ImageLocalPath); os.IsNotExist(err) {
		url := m.ImageOriginalURL
		response, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		// path/to/whatever does not exist
		log.Println(m.ImageLocalPath)
		file, err := os.Create(m.ImageLocalPath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// Use io.Copy to just dump the response body to the file. This supports huge files
		_, err = io.Copy(file, response.Body)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

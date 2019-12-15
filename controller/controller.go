package controller

import (
	"errors"

	"github.com/moviedb-api/database"
	"github.com/moviedb-api/model"
	"github.com/ryanbradynd05/go-tmdb"
)

// Controller :
type Controller struct {
	mdbc        *tmdb.TMDb
	imageConfig *model.ImageConfig
	datacon     *database.Database
}

// New :
func New(m *tmdb.TMDb) (*Controller, error) {
	ctrl := &Controller{
		mdbc: m,
	}
	err := ctrl.GenerateImageConfig()
	if err != nil {
		return nil, err
	}
	db, err := database.New()
	if err != nil {
		return nil, err
	}

	ctrl.datacon = db
	return ctrl, nil
}

func (ctrl *Controller) CloseDB() error {
	return ctrl.datacon.DB.Close()
}

// GenerateImageConfig :
func (ctrl *Controller) GenerateImageConfig() error {
	imgcf, err := ctrl.mdbc.GetConfiguration()

	if err != nil {
		return err
	}

	imageCon := &model.ImageConfig{}
	imageCon.BaseURL = imgcf.Images.BaseURL
	imageCon.SecureBaseURL = imgcf.Images.SecureBaseURL

	pssizes := []*model.PosterSize{}
	for _, v := range imgcf.Images.PosterSizes {
		pssizes = append(pssizes, &model.PosterSize{Size: v})

	}
	imageCon.PosterSizes = pssizes

	ctrl.imageConfig = imageCon

	if ctrl.imageConfig == nil {
		return errors.New("No image configuration found!")
	} else if len(ctrl.imageConfig.PosterSizes) == 0 {
		return errors.New("No Poster size available")
	}

	return nil

}

func (ctrl *Controller) GetDefaultPosterSize() (string, error) {

	size := ""
	if ctrl.imageConfig == nil {
		return size, errors.New("No image configuration found!")
	}

	size = ctrl.imageConfig.PosterSizes[0].Size
	return size, nil
}

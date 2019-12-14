package model

// ImageConfig : Required to get the linked image from TMDb
type ImageConfig struct {
	BaseURL       string
	SecureBaseURL string
	PosterSizes   []*PosterSize
	Model
}

// PosterSize  :Size available for image
type PosterSize struct {
	Size string
}

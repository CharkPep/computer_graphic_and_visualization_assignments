package ImageProcessor

import (
	"image"
)

type Format int

var FormatStrings = [...]string{"png", "jpeg", "gif", "bmp", "tiff"}

const (
	Png Format = iota
	Jpeg
	Gif
	Bmp
	Tiff
)

func (f Format) String() string {
	return FormatStrings[f]
}

type ImageSaveName = string

type ColorCorrection struct {
	RedCorrection   float64
	GreenCorrection float64
	BlueCorrection  float64
	AlphaCorrection float64
}

type ConvertOptions struct {
	Bounds          *image.Rectangle
	ColorCorrection *ColorCorrection
}

type ContrastFactor float64

type CropOptions struct {
	X      int
	Y      int
	Height *int
	Width  *int
	Invert *bool
}

// ImageConvertable defines the methods for converting images.
type ImageConvertable interface {
	Convert(Img *image.Image, Options *ConvertOptions)
}

// ImageCroppeable defines the methods for cropping images.
type ImageCroppeable interface {
	Crop(Img *image.Image, Options *CropOptions)
}

// ProcessableImage combines both the conversion and cropping functionality.
type ProcessableImage interface {
	ImageConvertable
	ImageCroppeable
	Save() error
}

type ImageProcessor interface {
	NewImageProcessor(...any) error
	CreateProcessableImage() (*ProcessableImage, error)
}

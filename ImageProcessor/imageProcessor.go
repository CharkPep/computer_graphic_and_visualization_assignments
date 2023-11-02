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

type Direction int

const (
	Horizontal Direction = iota
	Vertical
)

func (d Direction) String() string {
	return [...]string{"Horizontal", "Vertical"}[d]
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

type ConnectOptions struct {
	Images    []ProcessableImage
	Direction Direction
}

// ImageConvertable defines the methods for converting images.
type ImageConvertable interface {
	Convert(Img *image.Image, Options *ConvertOptions)
}

// ImageCroppeable defines the methods for cropping images.
type ImageCroppeable interface {
	Crop(Img *image.Image, Options *CropOptions)
}

// ImageContrastable defines the methods for increasing the contrast of images.
type ImageContrastable interface {
	IncreaseContrast(factor *ContrastFactor)
}

type ImageConnectable interface {
	Connect(Images []ProcessableImage, Direction *Direction)
}

// ProcessableImage combines both the conversion and cropping functionality.
type ProcessableImage interface {
	ImageConvertable
	ImageCroppeable
	ImageContrastable
	ImageConnectable
	Save() error
}

type ImageProcessor interface {
	NewImageProcessor(...any) error
	CreateProcessableImage() (*ProcessableImage, error)
}

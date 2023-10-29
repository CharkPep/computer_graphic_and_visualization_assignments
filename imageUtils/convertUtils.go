package imageUtils

import (
	"errors"
	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"imageProccesing/ImageProcessor"
	"io"
	"path/filepath"
)

// resolveAbsPath Function to resolve absolute path
func resolveAbsPath(path string) (string, error) {
	if !filepath.IsAbs(path) {
		absPath, err := filepath.Abs(path)
		if err != nil {
			return "", err
		}

		return absPath, nil
	}

	return path, nil
}

// Functions to convert an image to different formats
func toGif(image image.Image, dest io.Writer) error {
	if err := gif.Encode(dest, image, nil); err != nil {
		return err
	}

	return nil
}

func ToJpeg(image image.Image, dest io.Writer) error {
	if err := jpeg.Encode(dest, image, nil); err != nil {
		return err
	}

	return nil
}

func ToPng(image image.Image, dest io.Writer) error {
	if err := png.Encode(dest, image); err != nil {
		return err
	}

	return nil
}

func ToBmp(image image.Image, dest io.Writer) error {
	if err := bmp.Encode(dest, image); err != nil {
		return err
	}

	return nil
}

func ToTiff(image image.Image, dest io.Writer) error {
	if err := tiff.Encode(dest, image, nil); err != nil {
		return err
	}

	return nil
}

// ResizeImage Function to resize an image while maintaining aspect ratio
func ResizeImage(image *image.Image, bounds image.Rectangle) {
	if bounds.Dx() == 0 {
		ratio := float64(bounds.Dy()) / float64((*image).Bounds().Dy())
		newWidth := int(float64((*image).Bounds().Dx()) * ratio)
		*image = resize.Resize(uint(newWidth), uint(bounds.Dy()), *image, resize.Lanczos3)
	} else if bounds.Dy() == 0 {
		ratio := float64(bounds.Dx()) / float64((*image).Bounds().Dx())
		newHeight := int(float64((*image).Bounds().Dy()) * ratio)
		*image = resize.Resize(uint(bounds.Dx()), uint(newHeight), *image, resize.Lanczos3)
	}
	*image = resize.Resize(uint(bounds.Dx()), uint(bounds.Dy()), *image, resize.Lanczos3)
}

// adjustColorBalance Function to adjust color balance in an image
func adjustColorBalance(Img *image.Image, colorCorrectionOptions ImageProcessor.ColorCorrection) {
	bounds := (*Img).Bounds()
	outputImage := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, oldA := (*Img).At(x, y).RGBA()
			alpha := uint8(colorCorrectionOptions.AlphaCorrection)
			r = uint32(float64(r) * colorCorrectionOptions.RedCorrection)
			g = uint32(float64(g) * colorCorrectionOptions.GreenCorrection)
			b = uint32(float64(b) * colorCorrectionOptions.BlueCorrection)
			if colorCorrectionOptions.AlphaCorrection == 256 {
				alpha = uint8(oldA)
			}
			outputImage.Set(x, y, color.NRGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), alpha})
		}
	}

	*Img = outputImage
}

func convertToFormat(format *ImageProcessor.Format, image image.Image, dest io.Writer) error {
	switch (*format).String() {
	case ImageProcessor.Png.String():
		return ToPng(image, dest)
	case ImageProcessor.Jpeg.String():
		return ToJpeg(image, dest)
	case ImageProcessor.Gif.String():
		return toGif(image, dest)
	case ImageProcessor.Bmp.String():
		return ToBmp(image, dest)
	case ImageProcessor.Tiff.String():
		return ToTiff(image, dest)
	}

	return errors.New("invalid SaveFormat")

}

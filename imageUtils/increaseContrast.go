package imageUtils

import (
	"image"
	"image/color"
	"imageProccesing/ImageProcessor"
)

// IncreaseContrast Method to increase the contrast of an image
func (Img *ProcessableImage) IncreaseContrast(factor ImageProcessor.ContrastFactor) {
	bounds := (*Img.Img).Bounds()
	newImage := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := color.RGBAModel.Convert((*Img.Img).At(x, y)).(color.RGBA)
			r, g, b := float64(pixel.R), float64(pixel.G), float64(pixel.B)

			r = (r-128)*float64(factor) + 128
			g = (g-128)*float64(factor) + 128
			b = (b-128)*float64(factor) + 128

			r = clamp(r, 0, 255)
			g = clamp(g, 0, 255)
			b = clamp(b, 0, 255)

			newImage.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), pixel.A})
		}
	}

	*Img.Img = image.Image(newImage)
}

func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

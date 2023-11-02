package imageUtils

import (
	"image"
	"image/draw"
	"imageProccesing/ImageProcessor"
)

func Connect(Images []ProcessableImage, Direction *ImageProcessor.Direction) *ProcessableImage {
	switch (*Direction).String() {
	case ImageProcessor.Horizontal.String():
		maxHeight := 0
		maxWidth := 0
		for _, imageIterator := range Images {
			bounds := (*imageIterator.Img).Bounds()
			if bounds.Dy() > maxHeight {
				maxHeight = bounds.Dy()
			}
			maxWidth += bounds.Dx()
		}

		newImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{maxWidth, maxHeight}})
		currentX := 0
		for _, imageIterator := range Images {
			bounds := (*imageIterator.Img).Bounds()
			draw.Draw(newImage, image.Rectangle{image.Point{currentX, 0}, image.Point{currentX + bounds.Dx(), bounds.Dy()}}, *imageIterator.Img, image.Point{0, 0}, draw.Src)
			currentX += bounds.Dx()
		}

		var img image.Image = newImage

		return &ProcessableImage{
			Img:        &img,
			SaveFormat: Images[0].SaveFormat,
			SavePath:   Images[0].SavePath,
			Name:       Images[0].Name,
		}
	case ImageProcessor.Vertical.String():
		maxHeight := 0
		maxWidth := 0
		for _, imageIterator := range Images {
			bounds := (*imageIterator.Img).Bounds()
			if bounds.Dx() > maxWidth {
				maxWidth = bounds.Dx()
			}
			maxHeight += bounds.Dy()
		}

		newImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{maxWidth, maxHeight}})
		currentY := 0
		for _, imageIterator := range Images {
			bounds := (*imageIterator.Img).Bounds()
			draw.Draw(newImage, image.Rectangle{image.Point{0, currentY}, image.Point{bounds.Dx(), currentY + bounds.Dy()}}, *imageIterator.Img, image.Point{0, 0}, draw.Src)
			currentY += bounds.Dy()
		}

		var img image.Image = newImage
		return &ProcessableImage{
			Img:        &img,
			SaveFormat: Images[0].SaveFormat,
			SavePath:   Images[0].SavePath,
			Name:       Images[0].Name,
		}
	}

	return nil
}

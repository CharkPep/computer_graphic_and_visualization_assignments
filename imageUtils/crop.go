package imageUtils

import (
	"image"
	"image/draw"
	"imageProccesing/ImageProcessor"
)

// Crop Method to crop images
func (Img *ProcessableImage) Crop(Options *ImageProcessor.CropOptions) {
	if Options == nil {
		Options = &ImageProcessor.CropOptions{
			X:      0,
			Y:      0,
			Height: nil,
			Width:  nil,
			Invert: nil,
		}
	}
	bounds := (*Img.Img).Bounds()
	rgbaImage := image.NewRGBA(bounds)
	draw.Draw(rgbaImage, bounds, *Img.Img, image.Point{}, draw.Src)

	newBounds := (*Img.Img).Bounds()
	if Options.Height != nil && Options.Width != nil {
		newBounds = image.Rect(Options.X, Options.Y, Options.X+*Options.Width, Options.Y+*Options.Height)
	}
	if Options.Invert != nil && *Options.Invert {
		newBounds = (*Img.Img).Bounds()
	}

	*Img.Img = rgbaImage.SubImage(newBounds)
	if Options.Invert != nil && *Options.Invert {
		//Draw black rectangle (X,Y) -> (X + Width, Y + Height)
		drawRectangle := image.Rect(Options.X, Options.Y, Options.X+*Options.Width, Options.Y+*Options.Height)
		draw.Draw(rgbaImage, drawRectangle, image.NewUniform(image.Black), image.ZP, draw.Src)
		*Img.Img = rgbaImage
	}

}

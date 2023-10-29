package imageUtils

import "imageProccesing/ImageProcessor"

// Covert Method to convert images to the specified output SaveFormat
func (Img *ProcessableImage) Covert(Options *ImageProcessor.ConvertOptions) {
	if Options == nil {
		Options = &ImageProcessor.ConvertOptions{}
	}

	if Options.Bounds.Dx() != 0 || Options.Bounds.Dy() != 0 {
		ResizeImage(Img.Img, *Options.Bounds)
	}

	if Options.ColorCorrection != nil {
		adjustColorBalance(Img.Img, *Options.ColorCorrection)
	}

}

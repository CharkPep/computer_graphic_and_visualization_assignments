package imageUtils

import "github.com/sunshineplan/imgconv"

func (Img *ProcessableImage) DrawWaterMark(Options *imgconv.WatermarkOption) {
	imageWithWaterMark := imgconv.Watermark(*Img.Img, Options)
	Img.Img = &imageWithWaterMark
}

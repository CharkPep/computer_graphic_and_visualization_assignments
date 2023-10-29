package main

import (
	"github.com/alecthomas/kingpin/v2"
	"image"
	"imageProccesing/ImageProcessor"
	"imageProccesing/imageUtils"
	"log"
	"os"
	"strings"
)

//TODO add logs/debug
//TODO add loading bar

// Define global loggers for different log levels
var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

// Define available image formats
var availableFormats = []string{"png", "jpeg", "gif", "bmp", "tiff"}

// Define command-line flags using the Kingpin library
var (
	app             = kingpin.New("imgp", "A command-line tool for image processing")
	convert         = kingpin.Command("convert", "ProcessImage image(s) to another format")
	width           = convert.Flag("width", "output width (if not specified, maintain aspect ratio)").Short('w').Default("0").Int()
	height          = convert.Flag("height", "output height (if not specified, maintain aspect ratio)").Short('h').Default("0").Int()
	input           = kingpin.Flag("input", "input file(s)").Short('i').Required().Strings()
	output          = kingpin.Flag("output", "output folder").Short('o').Default(".").String()
	format          = kingpin.Flag("format", "output format").Short('f').Enum(ImageProcessor.FormatStrings[:]...)
	redCorrection   = convert.Flag("red-correction", "red color correction factor").Short('R').Default("1.0").Float64()
	greenCorrection = convert.Flag("green-correction", "green color correction factor").Short('G').Default("1.0").Float64()
	blueCorrection  = convert.Flag("blue-correction", "blue color correction factor").Short('B').Default("1.0").Float64()
	alphaCorrection = convert.Flag("alpha-correction", "alpha channel value").Short('A').Default("256.0").Float64()
	crop            = kingpin.Command("crop", "Crop image(s)")
	x               = crop.Flag("x", "x coordinate of the top left corner of the crop area").Short('X').Required().Int()
	y               = crop.Flag("y", "y coordinate of the top left corner of the crop area").Short('Y').Required().Int()
	cropWidth       = crop.Flag("width", "width of the crop area").Short('W').Required().Int()
	cropHeight      = crop.Flag("height", "height of the crop area").Short('H').Required().Int()
	cropInvert      = crop.Flag("invert", "invert the crop area").Short('I').Default("false").Bool()
	contrast        = kingpin.Command("contrast", "Adjust the contrast of image(s)")
	contrastFactor  = contrast.Flag("value", "contrast value").Short('V').Required().Float64()
)

func ImageIterator(images *imageUtils.ImagesCollection, saveFormat *ImageProcessor.Format, savePath string) func() (*imageUtils.ProcessableImage, error) {
	curImageIdx := 0
	return func() (*imageUtils.ProcessableImage, error) {
		if curImageIdx >= len(images.InputImagesPath) {
			return nil, nil
		}

		imageFile, err := os.Open(images.InputImagesPath[curImageIdx])
		if err != nil {
			return nil, err
		}

		imageDecoded, _, err := image.Decode(imageFile)
		if err != nil {
			return nil, err
		}

		curImageIdx++
		imageStat, err := imageFile.Stat()
		if err != nil {
			return nil, err
		}
		imageNameSplit := strings.Split(imageStat.Name(), ".")
		imageName := strings.Join(imageNameSplit[:len(imageNameSplit)-1], "")
		return &imageUtils.ProcessableImage{
			Img:        &imageDecoded,
			SaveFormat: saveFormat,
			SavePath:   savePath,
			Name:       &imageName,
		}, nil
	}

}

func constructColorCorrection() *ImageProcessor.ColorCorrection {
	if *redCorrection == 1.0 && *greenCorrection == 1.0 && *blueCorrection == 1.0 && *alphaCorrection == 256.0 {
		return nil
	}
	return &ImageProcessor.ColorCorrection{
		RedCorrection:   *redCorrection,
		GreenCorrection: *greenCorrection,
		BlueCorrection:  *blueCorrection,
		AlphaCorrection: *alphaCorrection,
	}
}

func main() {
	app.HelpFlag.Short('h')
	command := kingpin.Parse()
	// Initialize loggers for different log levels
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	var parsedFormat ImageProcessor.Format
	for i, formatItr := range ImageProcessor.FormatStrings {
		if formatItr == *format {
			parsedFormat = ImageProcessor.Format(i)
			break
		}
	}
	images, err := imageUtils.NewImageCollection(
		*input,
		*output)
	if err != nil {
		ErrorLogger.Fatal(err)
		return
	}

	imageItr := ImageIterator(images, &parsedFormat, *output)

	for loadedImage, err := imageItr(); loadedImage != nil && err == nil; loadedImage, err = imageItr() {
		switch command {
		case "convert":
			InfoLogger.Printf("Converting images to %s format", parsedFormat.String())
			convertOptions := &ImageProcessor.ConvertOptions{
				Bounds: &image.Rectangle{
					Min: image.Point{X: 0, Y: 0},
					Max: image.Point{X: *width, Y: *height},
				},
				ColorCorrection: constructColorCorrection(),
			}

			loadedImage.Covert(convertOptions)
			if err = loadedImage.Save(*loadedImage.Name); err != nil {
				ErrorLogger.Fatal(err)
				return
			}
		case "crop":
			cropOptions := &ImageProcessor.CropOptions{
				X:      *x,
				Y:      *y,
				Height: cropHeight,
				Width:  cropWidth,
				Invert: cropInvert,
			}
			loadedImage.Crop(cropOptions)
			if err = loadedImage.Save(*loadedImage.Name); err != nil {
				ErrorLogger.Fatal(err)
				return
			}
		case "contrast":
			loadedImage.IncreaseContrast(ImageProcessor.ContrastFactor(*contrastFactor))
			if err = loadedImage.Save(*loadedImage.Name); err != nil {
				ErrorLogger.Fatal(err)
				return
			}
		}
	}

}

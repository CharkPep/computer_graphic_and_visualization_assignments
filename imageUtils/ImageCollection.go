package imageUtils

import (
	"bytes"
	"errors"
	"image"
	"imageProccesing/ImageProcessor"
	"os"
)

// ImagesCollection Define a struct to represent image-related settings
type ImagesCollection struct {
	InputImagesPath  []string
	OutputImagesPath string
}

// ProcessableImage Define a struct to represent an image that can be processed
type ProcessableImage struct {
	Img        *image.Image
	SaveFormat *ImageProcessor.Format
	SavePath   string
	Name       *string
}

// NewImageCollection Method to create a new ImagesCollection object
func NewImageCollection(
	InputPath []string,
	OutputPath string,
) (*ImagesCollection, error) {
	OutputPath, err := resolveAbsPath(OutputPath)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(OutputPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if !fileInfo.IsDir() {
		return nil, errors.New("output OutputImagesPath is not a directory")
	}

	return &ImagesCollection{
		InputImagesPath:  InputPath,
		OutputImagesPath: OutputPath,
	}, nil
}

// Save Method to save images
func (Img *ProcessableImage) Save(ImageName ImageProcessor.ImageSaveName) error {
	buf := new(bytes.Buffer)
	if err := convertToFormat(Img.SaveFormat, *Img.Img, buf); err != nil {
		return err
	}

	file, err := os.Create(Img.SavePath + "/" + ImageName + "." + Img.SaveFormat.String())
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err = file.Write(buf.Bytes()); err != nil {
		return err
	}

	return nil
}

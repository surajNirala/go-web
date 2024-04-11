package imageprocessing

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func ReadImage(path string) image.Image {
	inputFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Could not open file: ", path)
		panic(err)
	}
	defer inputFile.Close()
	//Decode Image
	img, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Println("Path", path)
		panic(err)
	}
	return img
}

func WriteImage(path string, img image.Image) {
	outputFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	//Encode Image
	err = jpeg.Encode(outputFile, img, nil)
	if err != nil {
		fmt.Println("Path", path)
		panic(err)
	}
}

func GrayScale(img image.Image) image.Image {
	//create new gray image
	bounds := img.Bounds()
	grapImg := image.NewGray(bounds)

	//Covert each pixel to grayscale
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			orignalPixel := img.At(x, y)
			grayPixel := color.GrayModel.Convert(orignalPixel)
			grapImg.Set(x, y, grayPixel)
		}
	}
	return grapImg
}

func Resize1(img image.Image) image.Image {
	newWidth := uint(400)
	newHeight := uint(400)
	reSizedImage := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
	return reSizedImage
}

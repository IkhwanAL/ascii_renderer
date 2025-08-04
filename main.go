package main

import (
	"fmt"
	"image"
	"log"
	"math"
	"os"

	"image/color"
	"image/jpeg"
	_ "image/png"

	"github.com/ikhwanal/ascii_renderer/core"
)

func main() {
	// TODO Get Image
	imgPath := "./img/lorem.png"
	
	reader, err := os.Open(imgPath)

	if err != nil {
		log.Fatal(err)
	}

	defer reader.Close()

	img, _, err := image.Decode(reader)

	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()
	
	newImageSet := image.NewGray(bounds)
	
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldPixel := img.At(x, y)
			pixel := color.GrayModel.Convert(oldPixel)
			newImageSet.Set(x, y, pixel)
		}
	}

	// newImageSet = ScaleImage(newImageSet, bounds.Bounds().Max.X / 3, bounds.Bounds().Max.Y / 3) 
  newImageSet = ScaleImage(newImageSet, 80, 40)
	outFile, err := os.Create("./img/testResult1.jpg")

	if err != nil {
		log.Fatal(err)
	}

	defer outFile.Close()

	err = jpeg.Encode(outFile, newImageSet, nil)

	if err != nil {
		log.Fatal(err)
	}

	core.RenderToAscii(newImageSet)
	fmt.Printf("Width  : %d, To %d \n", bounds.Max.X, newImageSet.Bounds().Max.X)
	fmt.Printf("Height : %d, T0 %d \n", bounds.Max.Y, newImageSet.Bounds().Max.Y)
	fmt.Print("Done\n")

}

func ScaleImage(img *image.Gray, targetWidthScale int, targetHeightScale int) *image.Gray {
	imageScale := image.NewGray(image.Rect(0,0, targetWidthScale, targetHeightScale)) // Create New Predefined Empty Image

	Sx := float64(img.Bounds().Max.X) / float64(targetWidthScale) // Scale X Value
	Sy := float64(img.Bounds().Max.Y) / float64(targetHeightScale) // Scale Y Value

	for y := 0; y < targetHeightScale; y++ {
		for x := 0; x < targetWidthScale; x++ {
			originalX := Sx * float64(x) // Which X Position are in Original X
			originalY := Sy * float64(y) // Which Y Position Are in Original Y

			roundedX := int(math.RoundToEven(originalX)) // Round the value (Nearest Neighbor)
			roundedY := int(math.RoundToEven(originalY)) // Round the value (Nearest Neighbot)

			// This is making sure that rounded pixel value is not overlapped with max original value
			clappedX := min(max(0, roundedX), img.Bounds().Max.X-1) 
			clappedY := min(max(0, roundedY), img.Bounds().Max.Y-1)

			pixelColor := img.At(clappedX, clappedY) // Get The Value From Original Image

			imageScale.Set(x, y, pixelColor)
		}
	}

	return imageScale
}

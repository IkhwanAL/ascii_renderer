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
)

func main() {
	// TODO Get Image
	imgPath := "./img/secondTest.jpg"
	
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

	newImageSet := image.NewRGBA(bounds)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			oldPixel := img.At(x, y)
			pixel := color.GrayModel.Convert(oldPixel)
			newImageSet.Set(x, y, pixel)
		}
	}

	test := ScaleImage(newImageSet, newImageSet.Bounds().Max.X / 3, newImageSet.Bounds().Max.Y / 3) 
	
	outFile, err := os.Create("./img/testResult2.jpg")

	if err != nil {
		log.Fatal(err)
	}

	defer outFile.Close()

	err = jpeg.Encode(outFile, test, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Width  : %d, To %d \n", bounds.Max.X, test.Bounds().Max.X)
	fmt.Printf("Height : %d, T0 %d \n", bounds.Max.Y, test.Bounds().Max.Y)
	fmt.Print("Done")
}

func ScaleImage(img *image.RGBA, targetWidthScale int, targetHeightScale int) *image.RGBA {
	imageScale := image.NewRGBA(image.Rect(0,0, targetWidthScale, targetHeightScale)) // Create New Predefined Empty Image

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

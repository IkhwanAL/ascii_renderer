package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"image/color"
	"image/jpeg"
	_ "image/png"
)

func main() {
	// TODO Get Image
	imgPath := "./img/dice_5.png"
	
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

	outFile, err := os.Create("./img/testResult.jpg")

	if err != nil {
		log.Fatal(err)
	}

	defer outFile.Close()

	err = jpeg.Encode(outFile, newImageSet, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Done")
}

package utils

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

func OutputImageForDebugResult(img image.Image, filePathName string) {
	outFile, err := os.Create("./img/scaleResult.jpg")
	if err != nil {
		log.Fatal(err)
	}

	defer outFile.Close()

	err = jpeg.Encode(outFile, img, nil)

	if err != nil {
		log.Fatal(err)
	}
}

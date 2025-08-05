package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"image/jpeg"
	_ "image/png"

	"github.com/ikhwanal/ascii_renderer/core"
	"golang.org/x/term"
)

func getTerminalSize() (int, int, error) {
	fd := int(os.Stdout.Fd())

	if !term.IsTerminal(fd) {
		return 0, 0, fmt.Errorf("your not in terminal bro")
	}

	width, height, err := term.GetSize(fd)
	if err != nil {
		fmt.Print("Handle")
		return 80, 25, err
	}

	return width, height, nil
}

func main() {
	// TODO Get Image
	imgPath := "./img/Weebs1.jpg"

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
	
	newImageSet := core.ConvertToGrayScale(img, bounds)

	newImageSet = core.ScaleImage(newImageSet, bounds.Bounds().Max.X / 4, bounds.Bounds().Max.Y / 4) 
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

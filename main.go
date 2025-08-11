package main

import (
	"fmt"
	"image"
	_ "image/png"

	// "image/draw"
	"log"
	"math"
	"os"

	"github.com/ikhwanal/ascii_renderer/core"
	"github.com/ikhwanal/ascii_renderer/utils"
	"golang.org/x/term"
)

func getTerminalSize() (int, int, error) {
	fd := int(os.Stdout.Fd())

	if !term.IsTerminal(fd) {
		return 0, 0, fmt.Errorf("your not in terminal bro")
	}

	width, height, err := term.GetSize(fd)
	if err != nil {
		return 80, 25, err
	}

	return width, height, nil
}

func main() {
	// TODO Get Image
	imgPath := "./img/tenTest.jpg"

	reader, err := os.Open(imgPath)

	if err != nil {
		log.Fatal(err)
	}

	defer reader.Close()

	w, h, err := getTerminalSize()

	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(reader)

	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()

	newImageSet := core.ConvertToGrayScale(img, bounds)

	utils.OutputImageForDebugResult(newImageSet, "./img/greyResult.jpg")

	edgeImg := core.EdgeDetection(*newImageSet)

	utils.OutputImageForDebugResult(newImageSet, "./img/edgeDetectionResult.jpg")

	widthDivisor := float64(bounds.Bounds().Max.X) / float64(w)

	heightDivisor := float64(bounds.Bounds().Max.Y) / float64(h*2)

	finalDivisor := math.Round(max(widthDivisor, heightDivisor))

	newImageSet = core.BilinearScaleGray(
		newImageSet,
		bounds.Bounds().Max.X/int(finalDivisor),
		bounds.Bounds().Max.Y/int(finalDivisor),
	)

	utils.OutputImageForDebugResult(newImageSet, "./img/scaleResult.jpg")

	edgeImg = core.MaxPoolingGray(
		edgeImg,
		bounds.Bounds().Max.X/int(finalDivisor),
		bounds.Bounds().Max.Y/int(finalDivisor),
	)

	utils.OutputImageForDebugResult(newImageSet, "./img/edgeImageScale.jpg")

	core.RenderToAsciiWithEdgeContext(newImageSet, edgeImg)
	// core.RenderToAscii(newImageSet)
	fmt.Printf("Width  : %d, To %d Max %d\n", bounds.Max.X, newImageSet.Bounds().Max.X, w)
	fmt.Printf("Height : %d, T0 %d Max %d\n", bounds.Max.Y, newImageSet.Bounds().Max.Y, h)
	fmt.Print("Done\n")

}

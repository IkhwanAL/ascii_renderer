package core

import (
	"fmt"
	"image"
)
// const DENSITY = " .:-=+*#%@"
// const DENSITY = " .'`^\",:;Il!i><~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"
// Sort By The Brigest To The Darkest Color
const DENSITY = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "

func RenderToAscii(img *image.Gray) {
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {

			grey := img.GrayAt(x, y)

			density_index := int(grey.Y) * (len(DENSITY) - 1) / 255
			character := DENSITY[density_index]
			fmt.Printf("%s",string(character))
		}
		fmt.Println()
	}
}

func RenderToAsciiWithEdgeContext(img *image.Gray, edgeImg *image.Gray) {
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			// gray := img.GrayAt(x, y)

			edgeGray := edgeImg.GrayAt(x, y)

			fmt.Printf("%d ", edgeGray.Y)
		}
		fmt.Println()
	}
}

package core

import (
	"fmt"
	"image"
)
const DENSITY = " .:-=+*#%@"
// const DENSITY = " .'`^\",:;Il!i><~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"

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

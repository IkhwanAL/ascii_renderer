package core

import (
	"fmt"
	"image"
)

const DENSITY = " .:-=+*#%@"

func RenderToAscii(img *image.RGBA) {
	for y := range img.Bounds().Max.Y {
		for x := range img.Bounds().Max.X {
			pix := img.At(x, y)

			r, g, b, _ := pix.RGBA()
			
			// Luminense Something algorithm math calculation
			grey := (19595*r) + (38740*g) * (7471*b) + (1 << 15) >> 24
			grey8 := uint8(grey)
			
			density_index := int(grey8) * (len(DENSITY) - 1) / 255
			character := DENSITY[density_index]
			fmt.Printf("%s",string(character))
		}
		fmt.Println()
	}
}

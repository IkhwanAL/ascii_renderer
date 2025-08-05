package core

import (
	"image"
	"math"
)

func ScaleImage(img *image.Gray, targetWidthScale int, targetHeightScale int) *image.Gray {
	imageScale := image.NewGray(image.Rect(0, 0, targetWidthScale, targetHeightScale)) // Create New Predefined Empty Image

	Sx := float64(img.Bounds().Max.X) / float64(targetWidthScale)  // Scale X Value
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

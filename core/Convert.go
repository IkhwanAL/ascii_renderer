package core

import (
	"image"
	"image/color"
)

func ConvertToGrayScale(img image.Image, bounds image.Rectangle) *image.Gray {
	newImageSet := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldPixel := img.At(x, y)
			pixel := color.GrayModel.Convert(oldPixel)
			newImageSet.Set(x, y, pixel)
		}
	}

	return newImageSet
}

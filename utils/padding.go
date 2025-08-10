package utils

import (
	"image"
	"image/draw"
)

// Add Padding in Image
// Replicate the pixel to the Padded to the image 
// it easy when you try to do kernel convolution image technique
func AddPaddingImage(img *image.Gray) *image.Gray {
	newImg := image.NewGray(image.Rect(0, 0, img.Bounds().Dx()+2, img.Bounds().Dy()+2))

	grayImage := image.Rect(1, 1, img.Bounds().Dx()+1, img.Bounds().Dy()+1)

	draw.Draw(newImg, grayImage, img, img.Bounds().Min, draw.Src)

	for x := 1; x < newImg.Bounds().Max.X-1; x++ {
		topGrayColor := newImg.GrayAt(x, 1)
		newImg.SetGray(x, 0, topGrayColor)

		bottomGrayColor := newImg.GrayAt(x, newImg.Bounds().Max.Y-2)
		newImg.SetGray(x, newImg.Bounds().Max.Y-1, bottomGrayColor)
	}

	for y := 0; y < newImg.Bounds().Max.Y; y++ {
		leftGrayColor := newImg.GrayAt(1, y)
		newImg.SetGray(0, y, leftGrayColor)

		rightGrayColor := newImg.GrayAt(newImg.Bounds().Max.X-2, y)
		newImg.SetGray(img.Bounds().Max.X-1, y, rightGrayColor)
	}

	return newImg
}

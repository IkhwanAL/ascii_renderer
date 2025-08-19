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

func AddEdgePaddingExtenstion(img *image.Gray, edgeTopBottom int, edgeLeftRight int) *image.Gray {
	oldW := img.Bounds().Dx()
	oldH := img.Bounds().Dy()

	newW := oldW + 2*edgeLeftRight
	newH := oldH + 2*edgeTopBottom

	newImg := image.NewGray(image.Rect(0, 0, newW, newH))

	dstRect := image.Rect(edgeLeftRight, edgeTopBottom, edgeLeftRight+oldW, edgeTopBottom+oldH)

	draw.Draw(newImg, dstRect, img, img.Bounds().Min, draw.Src)

	topLeftColor := img.GrayAt(0,0)
	topRightColor := img.GrayAt(oldW - 1, 0)
	bottomLeftColor := img.GrayAt(0, oldH - 1)
	bottomRigthColor := img.GrayAt(oldW - 1, oldH - 1)

	// Add Color In Corner
	for y := 0; y < edgeTopBottom; y++ {
		for x := 0; x < edgeLeftRight; x++ {
			newImg.SetGray(x, y, topLeftColor)
			newImg.SetGray(newW - 1 - x, y, topRightColor)
			newImg.SetGray(x, newImg.Bounds().Dy() - 1 - y, bottomLeftColor)
			newImg.SetGray(newImg.Bounds().Dx() - 1 - x, newImg.Bounds().Dy() - 1 - y, bottomRigthColor)
		}
	}

	// Add Color In Top Bottom
	for y := 0; y < edgeTopBottom; y++ {
		for x := edgeLeftRight; x < edgeLeftRight + oldW; x++ {
			oldGrayTop := img.GrayAt(x - edgeLeftRight, 0)
			newImg.SetGray(x, y, oldGrayTop)

			oldGrayBottom := img.GrayAt(x - edgeLeftRight, img.Bounds().Dy()-1)
			newImg.SetGray(x, newImg.Bounds().Dy() - 1 - y, oldGrayBottom)
		}
	}

	// Add Color In Left Right
	for x := 0; x < edgeLeftRight; x++ {
		for y := edgeTopBottom; y < edgeTopBottom + oldH; y++ {
			oldGrayLeft := img.GrayAt(x, y - edgeTopBottom)
			newImg.SetGray(x, y, oldGrayLeft)

			oldGrayRight := img.GrayAt(img.Bounds().Dx()-1, y - edgeTopBottom)
			newImg.SetGray(newImg.Bounds().Dx() - 1 - x, y, oldGrayRight)
		}
	}
	return newImg
}

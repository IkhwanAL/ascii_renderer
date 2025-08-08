package core

import (
	"image"
	"image/color"
	"image/draw"
	"math"
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

func allocNewArray(x0, y0 int) [][]float64 {
	arr := make([][]float64, y0)
	for y := 0; y < y0; y++ {
		arr[y] = make([]float64, x0)
	}

	return arr
}

func EdgeDetection(img image.Gray) *image.Gray {
	Gx := allocNewArray(img.Bounds().Dx(), img.Bounds().Dy())
	Gy := allocNewArray(img.Bounds().Dx(), img.Bounds().Dy())

	newPaddedImage := addPaddingImage(&img)

	horintalKernelConvluation := [3][3]int{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}

	verticalKernelConvluation := [3][3]int{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}

	for y := newPaddedImage.Bounds().Min.Y + 1; y < newPaddedImage.Bounds().Max.Y - 1; y++ {
		for x := newPaddedImage.Bounds().Min.X + 1; x < newPaddedImage.Bounds().Max.X - 1; x++ {
			horizontal := EdgeCalculation(newPaddedImage, x, y, horintalKernelConvluation)

			Gx[y-1][x-1] = horizontal

			vertical := EdgeCalculation(newPaddedImage,x, y, verticalKernelConvluation)

			Gy[y-1][x-1] = vertical

		}
	}

	imgGray := image.NewGray(image.Rect(0,0, img.Bounds().Dx(), img.Bounds().Dy()))

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			gx := Gx[y][x]
			gy := Gy[y][x]

			magnitude := math.Round(math.Sqrt(math.Pow(gx, 2) + math.Pow(gy, 2)))

			gray := color.Gray{Y: uint8(magnitude)}

			imgGray.SetGray(x, y, gray)
		}
	}

	return imgGray
}

func addPaddingImage(img *image.Gray) *image.Gray {
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

func EdgeCalculation(img *image.Gray, x, y int, kernelConvolution [3][3]int) float64 {
	pixelTopLeft := img.GrayAt(x-1, y-1).Y
	pixelTopCenter := img.GrayAt(x, y-1).Y
	pixelTopRight := img.GrayAt(x+1, y-1).Y

	pixelMiddleLeft := img.GrayAt(x-1, y).Y
	pixelMiddleCenter := img.GrayAt(x, y).Y
	pixelMiddleRight := img.GrayAt(x+1, y).Y

	pixelBottomLeft := img.GrayAt(x-1, y+1).Y
	pixelBottomCenter := img.GrayAt(x, y+1).Y
	pixelBottomRight := img.GrayAt(x+1, y+1).Y

	topLeft := int(pixelTopLeft) * kernelConvolution[0][0]
	topCenter := int(pixelTopCenter) * kernelConvolution[0][1]
	topRight := int(pixelTopRight) * kernelConvolution[0][2]

	middleLeft := int(pixelMiddleLeft) * kernelConvolution[1][0]
	middleCenter := int(pixelMiddleCenter) * kernelConvolution[1][1]
	middleRight := int(pixelMiddleRight) * kernelConvolution[1][2] 

	bottomLeft := int(pixelBottomLeft) * kernelConvolution[2][0]
	bottomCenter := int(pixelBottomCenter) * kernelConvolution[2][1]
	bottomRight := int(pixelBottomRight) * kernelConvolution[2][2]

	sumPix := float64((topLeft + topCenter + topRight + int(middleLeft) + int(middleCenter) + int(middleRight) + int(bottomLeft) + int(bottomCenter) + int(bottomRight))) / 9

	return sumPix
}

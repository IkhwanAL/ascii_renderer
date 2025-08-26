package core

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/ikhwanal/ascii_renderer/utils"
)

func AddDitheringAlgo(img image.Image) *image.Paletted {
	newImg := image.NewPaletted(img.Bounds(), color.Palette{color.Black, color.White})

	draw.FloydSteinberg.Draw(newImg, img.Bounds(), img, img.Bounds().Min)

	return newImg
}

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

func allocNewArray[V int | float64 | uint8](x0, y0 int) [][]V {
	arr := make([][]V, y0)
	for y := 0; y < y0; y++ {
		arr[y] = make([]V, x0)
	}

	return arr
}

func EdgeDetection(img image.Gray) *image.Gray {
	Gx := allocNewArray[float64](img.Bounds().Dx(), img.Bounds().Dy())
	Gy := allocNewArray[float64](img.Bounds().Dx(), img.Bounds().Dy())

	newPaddedImage := utils.AddPaddingImage(&img)

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

	for y := newPaddedImage.Bounds().Min.Y + 1; y < newPaddedImage.Bounds().Max.Y-1; y++ {
		for x := newPaddedImage.Bounds().Min.X + 1; x < newPaddedImage.Bounds().Max.X-1; x++ {
			horizontal := EdgeCalculation(newPaddedImage, x, y, horintalKernelConvluation)

			Gx[y-1][x-1] = horizontal

			vertical := EdgeCalculation(newPaddedImage, x, y, verticalKernelConvluation)

			Gy[y-1][x-1] = vertical

		}
	}

	imgGray := image.NewGray(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))

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

func GaussianBlur(img *image.Gray, sigma float64) (*image.Gray, error) {

	if img == nil {
		return nil, fmt.Errorf("mage Is nil")
	}

	if sigma <= 0 {
		fmt.Print(sigma)
		return img, fmt.Errorf("Error sigma should above 0")
	}

	// kernelSize := 6*sigma + 1
	kernelSize := int(2*math.Ceil(3*sigma) + 1) // Find The Corresping Size | How Blur The Image Are
	kernelGapFromCenter := int(math.Floor(float64(kernelSize) / 2)) // The Gap From 0 to Inf

	kernelConvolution := allocNewArray[float64](kernelSize, kernelSize)

	// fmt.Println(kernelGapFromCenter)
	sum := 0.0
	preCalculatedGaussianLeftFunc := 1 / (2 * math.Pi * math.Pow(sigma, 2)) 
	preExponentCalculation := 2.0*math.Pow(sigma,2) 

	for x := 0; x < kernelSize; x++ {
		for y := 0; y < kernelSize; y++ {
			xGap := x - kernelGapFromCenter
			yGap := y - kernelGapFromCenter

			// fmt.Printf("%d, %d \n", xGap, yGap)

			power := -(math.Pow(float64(xGap), 2) + math.Pow(float64(yGap), 2))

			exponent := power/preExponentCalculation

			gaussianFuncRightSize := math.Pow(math.E, exponent)

			gaussianFuncResult := preCalculatedGaussianLeftFunc * gaussianFuncRightSize

			kernelConvolution[x][y] = gaussianFuncResult

			sum += gaussianFuncResult
		}
	}

	sum = normalizationKernelGaussion(sum, kernelConvolution)

	// Edge Padding
	paddedImg := utils.AddEdgePaddingExtenstion(img, kernelGapFromCenter, kernelGapFromCenter)

	newImg := image.NewGray(img.Bounds())
	for y := 0 + kernelGapFromCenter; y < paddedImg.Bounds().Dy() - kernelGapFromCenter; y++ {
		for x := 0 + kernelGapFromCenter; x < paddedImg.Bounds().Dx() - kernelGapFromCenter; x++ {
 
			sumPixel := 0.0
			
			for kernelX := 0; kernelX  < len(kernelConvolution); kernelX++ {
				for kernelY := 0; kernelY < len(kernelConvolution[kernelX]); kernelY++ {
					colorGray := paddedImg.GrayAt(x - kernelGapFromCenter + kernelX, y-kernelGapFromCenter + kernelY).Y
					
					// fmt.Println(kernelConvolution[kernelX][kernalY])
					appliedKernel := float64(colorGray) * kernelConvolution[kernelX][kernelY]
					sumPixel += appliedKernel
				}
			}

			clr := color.Gray{
				Y: utils.ClampUint8(sumPixel, 0.0, 255.0),
			}
			newImg.Set(x - kernelGapFromCenter, y - kernelGapFromCenter, clr)
		}
	}

	return newImg, nil
}

func normalizationKernelGaussion(sum float64, kernelConvolution [][]float64) float64 {

	newSum := 0.0

	for y := range kernelConvolution {
		for x := range kernelConvolution[y] {
			kernelConvolution[x][y] = kernelConvolution[x][y] / sum
			newSum += kernelConvolution[x][y]
		}
	}

	return newSum
}

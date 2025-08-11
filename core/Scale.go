package core

import (
	// "fmt"	
	"image"
	"image/color"
	"math"
	"slices"
)

// NearestNeighbor Logic
func NearesetNeighborScale(img *image.Gray, targetWidthScale int, targetHeightScale int) *image.Gray {
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

// Bilinear Interpolation
func BilinearScaleGray(img *image.Gray, targetWidthScale int, targetHeightScale int) *image.Gray {
	imageScale := image.NewGray(image.Rect(0, 0, targetWidthScale, targetHeightScale))

	maxX := float64(img.Bounds().Max.X)
	maxY := float64(img.Bounds().Max.Y)

	Sx := maxX / float64(targetWidthScale)  // Scale X Value
	Sy := maxY / float64(targetHeightScale) // Scale Y Value

	for y := 0; y < targetHeightScale; y++ {
		for x := 0; x < targetWidthScale; x++ {
			// Possible Value Are Negative
			originalX := math.Max(0, math.Min(maxX-1, Sx*(float64(x)+0.5)-0.5)) // Update Image Sampling Strategy
			originalY := math.Max(0, math.Min(maxY-1, Sy*(float64(y)+0.5)-0.5)) // Update Image Sampling Strategy

			x1 := int(math.Floor(originalX))
			y1 := int(math.Floor(originalY))

			x2 := min(x1+1, img.Bounds().Max.X-1)
			y2 := min(y1+1, img.Bounds().Max.Y-1)

			// Find Distribution Weigth Between To Point
			dx := originalX - float64(x1)
			dy := originalY - float64(y1)

			bilinear := BILINEAR_RGBA{
				img: img,
				x1:  x1,
				x2:  x2,
				y1:  y1,
				y2:  y2,
			}

			gray := bilinear.ScaleGray(dx, dy)

			clr := color.Gray{
				Y: uint8(math.Round(gray)),
			}

			imageScale.Set(x, y, clr)
		}
	}

	return imageScale
}

type BILINEAR_RGBA struct {
	img image.Image
	x1  int
	x2  int
	y1  int
	y2  int
}

func (b BILINEAR_RGBA) RGBA_TOP_LEFT_NODE() color.Color {
	return b.img.At(b.x1, b.y1)
}

func (b BILINEAR_RGBA) RGBA_TOP_RIGHT_NODE() color.Color {
	return b.img.At(b.x2, b.y1)
}

func (b BILINEAR_RGBA) RGBA_BOTTOM_LEFT_NODE() color.Color {
	return b.img.At(b.x1, b.y2)
}

func (b BILINEAR_RGBA) RGBA_BOTTOM_RIGHT_NODE() color.Color {
	return b.img.At(b.x2, b.y2)
}

func (rgba BILINEAR_RGBA) ScaleGray(dx float64, dy float64) float64 {
	y1 := float64(rgba.img.(*image.Gray).GrayAt(rgba.x1, rgba.y1).Y)
	y2 := float64(rgba.img.(*image.Gray).GrayAt(rgba.x2, rgba.y1).Y)
	y3 := float64(rgba.img.(*image.Gray).GrayAt(rgba.x1, rgba.y2).Y)
	y4 := float64(rgba.img.(*image.Gray).GrayAt(rgba.x2, rgba.y2).Y)

	y := (y1 * ((1 - dx) * (1 - dy))) + 
		(y2 * ((dx) * (1 - dy))) + 
		(y3 * ((1 - dx) * (dy))) + 
		(y4 * ((dx) * (dy)))

	return y
}

func (rgba BILINEAR_RGBA) ScaleRGBA(dx float64, dy float64) (r, g, b, a float64) {
	clr1 := rgba.RGBA_TOP_LEFT_NODE()
	r1, g1, b1, a1 := clr1.RGBA()

	clr2 := rgba.RGBA_TOP_RIGHT_NODE()
	r2, g2, b2, a2 := clr2.RGBA()

	clr3 := rgba.RGBA_BOTTOM_LEFT_NODE()
	r3, g3, b3, a3 := clr3.RGBA()

	clr4 := rgba.RGBA_BOTTOM_RIGHT_NODE()
	r4, g4, b4, a4 := clr4.RGBA()

	// Distribution Weight Of Value Between Four Nodes
	r = (float64(r1) * ((1 - dx) * (1 - dy))) + (float64(r2) * ((dx) * (1 - dy))) + (float64(r3) * ((1 - dx) * (dy))) + (float64(r4) * ((dx) * (dy)))
	g = (float64(g1) * ((1 - dx) * (1 - dy))) + (float64(g2) * ((dx) * (1 - dy))) + (float64(g3) * ((1 - dx) * (dy))) + (float64(g4) * ((dx) * (dy)))
	b = (float64(b1) * ((1 - dx) * (1 - dy))) + (float64(b2) * ((dx) * (1 - dy))) + (float64(b3) * ((1 - dx) * (dy))) + (float64(b4) * ((dx) * (dy)))
	a = (float64(a1) * ((1 - dx) * (1 - dy))) + (float64(a2) * ((dx) * (1 - dy))) + (float64(a3) * ((1 - dx) * (dy))) + (float64(a4) * ((dx) * (dy)))

	return r, g, b, a
}

func BilinearScaleRGBA(img *image.RGBA, targetWidthScale int, targetHeightScale int) *image.RGBA {
	imageScale := image.NewRGBA(image.Rect(0, 0, targetWidthScale, targetHeightScale))

	maxX := float64(img.Bounds().Max.X)
	maxY := float64(img.Bounds().Max.Y)

	Sx := maxX / float64(targetWidthScale)  // Scale X Value
	Sy := maxY / float64(targetHeightScale) // Scale Y Value

	for y := 0; y < targetHeightScale; y++ {
		for x := 0; x < targetWidthScale; x++ {
			// Possible Value Are Negative
			originalX := math.Max(0, math.Min(maxX-1, Sx*(float64(x)+0.5)-0.5)) // Update Image Sampling Strategy
			originalY := math.Max(0, math.Min(maxY-1, Sy*(float64(y)+0.5)-0.5)) // Update Image Sampling Strategy

			x1 := int(math.Floor(originalX))
			y1 := int(math.Floor(originalY))

			x2 := min(x1+1, img.Bounds().Max.X-1)
			y2 := min(y1+1, img.Bounds().Max.Y-1)

			dx := originalX - float64(x1)
			dy := originalY - float64(y1)

			bilinear := BILINEAR_RGBA{
				img: img,
				x1:  x1,
				x2:  x2,
				y1:  y1,
				y2:  y2,
			}

			r, g, b, a := bilinear.ScaleRGBA(dx, dy)
			clr := color.RGBA64{
				R: uint16(math.Round(r)),
				G: uint16(math.Round(g)),
				B: uint16(math.Round(b)),
				A: uint16(math.Round(a)),
			}

			imageScale.Set(x, y, clr)
		}
	}

	return imageScale
}

// Value Return 2 ^ size
func powerOf2N(size float64) int {
	return int(math.Pow(float64(2), math.Log2(size)))
} 

func MaxPoolingGray(img *image.Gray, targetWidthScale, targetHeightScale int) *image.Gray {

	// It's Better To Make the Stride have a value of 2 ^ n
	// This allowed to have Optimze with SIMD instruction
	// This thing can be improved for later
	strideY := powerOf2N(float64(img.Bounds().Max.Y) / float64(targetHeightScale))
  strideX := powerOf2N(float64(img.Bounds().Max.X) / float64(targetWidthScale))

	actualTargetWidth := int(img.Bounds().Max.X / strideX)
	actualTargetHeight := int(img.Bounds().Max.Y / strideY)

	imgScale := image.NewGray(image.Rect(0, 0, actualTargetWidth, actualTargetHeight))

	kernel := make([]uint8, strideX*strideY)

	// Step 1 Max Pooling Thing
	for y := 0; y < imgScale.Bounds().Max.Y; y++ {
		for x := 0; x < imgScale.Bounds().Max.X; x++ {

			kernelStep := 0
			startY := y * strideY
			startX := x * strideX

			endY := min(img.Bounds().Max.Y, startY + strideY)
			endX := min(img.Bounds().Max.X, startX + strideX)

			// Kernel Convolution / Spatial Pooling
			for yy := startY; yy < endY; yy++ {
				for xx := startX; xx < endX; xx++ {
					gray := img.GrayAt(xx, yy).Y
					kernel[kernelStep] = gray
					kernelStep++
				}
			}

			maxValue := slices.Max(kernel[:kernelStep])

			imgScale.SetGray(x, y, color.Gray{Y: maxValue})
			kernelStep = 0
		}
	}

	if (targetHeightScale != actualTargetHeight) || (targetWidthScale != actualTargetWidth) {
		imgScale = BilinearScaleGray(imgScale, targetHeightScale, targetHeightScale)
	}

	return imgScale
}

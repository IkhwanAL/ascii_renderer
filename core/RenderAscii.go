package core

import (
	"fmt"
	"image"
	"math"
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
			fmt.Printf("%s", string(character))
		}
		fmt.Println()
	}
}

func mean(allEdgeMagnitude *[]uint8) float64 {
	sum := 0

	for _, v := range *allEdgeMagnitude {
		sum += int(v)
	}

	return float64(sum) / float64(len(*allEdgeMagnitude))
}

func std(mean float64, allEdgeMagnitude *[]uint8) float64 {
	sum := 0.0

	for _, v := range *allEdgeMagnitude {
		x := float64(v) - mean

		sample := math.Pow(x, 2)

		sum += sample
	}

	return math.Sqrt(sum / float64(len(*allEdgeMagnitude)-1))
}

func getCorrespondingEdgeMagnitude(edgeImg *image.Gray, variantInput float64) float64 {
	var activeEdge []uint8

	for y := edgeImg.Bounds().Min.Y; y < edgeImg.Bounds().Max.Y; y++ {
		for x := edgeImg.Bounds().Min.X; x < edgeImg.Bounds().Max.X; x++ {
			edgeGray := edgeImg.GrayAt(x, y).Y

			activeEdge = append(activeEdge, edgeGray)
		}
	}

	meanMagnitude := mean(&activeEdge)
	stdMagnitude := std(meanMagnitude, &activeEdge)

	return meanMagnitude + (variantInput * stdMagnitude)
}

// Dark to Light
// const DENSITY_CHARS = " .:-=+*#%@"
// const EDGE_CHARS      = ".,:;!?\"')([]}{/<>\\|+=*#@"

// Light to Dark
const DENSITY_CHARS = "@%#*+=-:. "
const EDGE_CHARS = "@#*+=|\\<>/}{][)(\"'?!:;,.";

func RenderToAsciiWithEdgeContext(img *image.Gray, edgeImg *image.Gray) {
	thresholdMagnitude := getCorrespondingEdgeMagnitude(edgeImg, 1.5)

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			gray := img.GrayAt(x, y).Y
			edgeGray := edgeImg.GrayAt(x, y).Y

			if float64(edgeGray) > thresholdMagnitude {
				edgeIndex := int(math.Floor((float64(gray) / 255.0) * float64(len(EDGE_CHARS)-1)))
				character := EDGE_CHARS[edgeIndex]
				fmt.Printf("%s", string(character))
			} else {
				densityIndex := int(math.Floor((float64(gray) / 255.0) * float64(len(DENSITY_CHARS) - 1)))
				character := DENSITY_CHARS[densityIndex]
				fmt.Printf("%s", string(character))
			}
		}
		fmt.Println()
	}
}

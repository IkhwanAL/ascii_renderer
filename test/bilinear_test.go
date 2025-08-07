package test

import (
	"image"
	"image/color"
	"testing"

	"github.com/ikhwanal/ascii_renderer/core"
)

func createTestImage(w, h int) *image.Gray {
	img := image.NewGray(image.Rect(0, 0, w, h))

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, color.Gray{Y: 128})
		}
	}

	return img
}

// Checkerboard pattern - good for testing interpolation
func createCheckerboard(width, height int) *image.Gray {
	img := image.NewGray(image.Rect(0, 0, width, height))
	squareSize := 4 // Size of each square

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Determine if we're in a black or white square
			squareX := x / squareSize
			squareY := y / squareSize
			isBlack := (squareX+squareY)%2 == 0

			if isBlack {
				img.Set(x, y, color.Gray{Y: 0}) // Black
			} else {
				img.Set(x, y, color.Gray{Y: 255}) // White
			}
		}
	}
	return img
}

func imagesIdentical(img1, img2 *image.Gray) bool {
	if !img1.Bounds().Eq(img2.Bounds()) {
		return false
	}

	bounds := img1.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if img1.GrayAt(x, y).Y != img2.GrayAt(x, y).Y {
				return false
			}
		}
	}
	return true
}

// ////To Make Sure Not Crash The App If We Choose To Scale Aggresively
// func TestScaleBoundaryImage(t *testing.T) {
// 	img := createTestImage(10, 10)
//
// 	result := core.BiliniarScale(img, 50, 50) // Scale up
// 	if result == nil {
// 		t.Fatal("Scaling up should not crash")
// 	}
//
// 	result = core.BiliniarScale(img, 3, 3) // Scale down to very small
// 	if result == nil {
// 		t.Fatal("Scaling to tiny size should not crash")
// 	}
//
// 	result = core.BiliniarScale(img, 1, 1) // Edge case: 1x1
// 	if result == nil {
// 		t.Fatal("Scaling to 1x1 should not crash")
// 	}
// }
//
// ////A Normal Test Just To make Sure Bilinear Work Fine Generally
// func TestBilinearSmoke(t *testing.T) {
// 	// Create checkerboard pattern - sensitive to coordinate errors
// 	img := createCheckerboard(20, 20)
//
// 	// Test multiple scale factors
// 	scales := []int{5, 10, 40, 80} // Down, down, up, up
//
// 	for _, scale := range scales {
// 		result := core.BiliniarScale(img, scale, scale)
//
// 		// Basic sanity checks
// 		if result == nil {
// 			t.Fatalf("Scale to %dx%d returned nil", scale, scale)
// 		}
//
// 		// Does not give use a correct scale value
// 		if result.Bounds().Dx() != scale || result.Bounds().Dy() != scale {
// 			t.Fatalf("Scale to %dx%d gave wrong dimensions", scale, scale)
// 		}
//
// 		// Should be different from nearest neighbor (unless very small)
// 		if scale > 10 {
// 			nearest := core.NearesetNeighborScale(img, scale, scale)
// 			if imagesIdentical(result, nearest) {
// 				t.Errorf("Bilinear identical to nearest neighbor at %dx%d", scale, scale)
// 			}
// 		}
// 	}
// }

//// This Function Is To Test The Math Equation Are Correct
func TestWeightDistribution(t *testing.T) {
    // Create image where we can verify weight calculation
    img := image.NewGray(image.Rect(0, 0, 2, 2))
    img.Set(0, 0, color.Gray{Y: 50})   // Use small distinct values
    img.Set(1, 0, color.Gray{Y: 100})   // so we can see the math clearly
    img.Set(0, 1, color.Gray{Y: 150})
    img.Set(1, 1, color.Gray{Y: 200})
    
    result := core.BiliniarScale(img, 3, 3)
 		   
    // Check that interpolated values are within reasonable bounds
    for y := 0; y < 3; y++ {
        for x := 0; x < 3; x++ {
            pixel := result.GrayAt(x, y).Y
            // All interpolated values should be between min and max of corners
            if pixel < 50 || pixel > 200 {
                t.Errorf("Pixel (%d,%d) = %d is outside valid range [50,200]", x, y, pixel)
            }
        }
    }
}

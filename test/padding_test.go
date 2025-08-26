package test

import (
	"testing"

	"github.com/ikhwanal/ascii_renderer/utils"
)

func TestPaddingImage(t *testing.T) {
	testImage := createTestImage(10, 10)

	afterPaddedImage := utils.AddEdgePaddingExtenstion(testImage, 3, 3)

	if testImage.Bounds().Dx() == afterPaddedImage.Bounds().Dx() {
		t.Fatalf("new image Width (%d) is suppposed not the same as test image (%d)", afterPaddedImage.Bounds().Dx(), testImage.Bounds().Dx())
	}
	
	if testImage.Bounds().Dy() == afterPaddedImage.Bounds().Dy() {
		t.Fatalf("new image height (%d) is suppposed not the same as test image (%d)", afterPaddedImage.Bounds().Dy(), testImage.Bounds().Dy())
	}

	if testImage.Bounds().Dx() >= afterPaddedImage.Bounds().Dx() {
		t.Fatalf("new image width (%d) is must bigger than test image (%d)", afterPaddedImage.Bounds().Dx(), testImage.Bounds().Dx())
	}

	if testImage.Bounds().Dy() >= afterPaddedImage.Bounds().Dy() {
		t.Fatalf("new image height (%d) is must bigger than test image (%d)", afterPaddedImage.Bounds().Dy(), testImage.Bounds().Dy())

	}
}

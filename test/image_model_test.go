package test

import (
	"testing"

	"github.com/ashalfarhan/random-gallery/api/models"
)

func TestModelValidation(t *testing.T) {
	t.Run("should be failed if no label and url", func(t *testing.T) {
		img := &models.Image{}
		if err := img.Validate(); err == nil {
			t.Fatalf("\nexpected error, but got: %v", err.Error())
		}
	})

	t.Run("should be failed if no label or url", func(t *testing.T) {
		img := &models.Image{Label: "asd"}
		if err := img.Validate(); err == nil {
			t.Fatalf("\nexpected error, but got: %v", err.Error())
		}
	})

	t.Run("should be success", func(t *testing.T) {
		img := &models.Image{Label: "My Image", Url: "https://example.com/asd.png"}
		if err := img.Validate(); err != nil {
			t.Fatalf("\nexpected error to be %v, but got: %v", nil, err.Error())
		}
	})
}

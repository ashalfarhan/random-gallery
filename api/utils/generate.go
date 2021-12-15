package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ashalfarhan/random-gallery/api/models"
	"github.com/google/uuid"
	"syreclabs.com/go/faker"
)

func GenerateRandomImages(len int) []models.Image {
	images := make([]models.Image, len)

	rand.Seed(time.Now().Unix())
	for i := 0; i < len; i++ {
		r := rand.Intn(240) - i
		h := r + 420
		image := models.Image{
			ID:     uuid.NewString(),
			Label:  faker.Lorem().Sentence(2),
			Url:    fmt.Sprintf("https://picsum.photos/seed/%d/640/%d.webp", r, h),
			Height: h,
			Width:  640,
		}

		images[i] = image
	}

	return images
}

package repository

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/ashalfarhan/gallery-api/api/models"
	"github.com/ashalfarhan/gallery-api/api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func CreateImage(b io.ReadCloser) (*models.Image, error) {
	var ni *models.Image
	err := json.NewDecoder(b).Decode(&ni)
	if err != nil {
		return nil, err
	}
	err = ni.Validate()
	if err != nil {
		return nil, err
	}
	ni.ID = uuid.NewString()
	return ni, nil
}

func GetAllImages(lim int, page int) ([]models.Image, error) {
	v := validator.New()
	err := v.Var(&lim, "gt=0,lt=25")
	if err != nil {

		return nil, errors.New("limit must be gt 0 or lt 25")
	}
	return utils.GenerateRandomImages(lim), nil
}

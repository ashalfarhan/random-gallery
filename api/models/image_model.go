package models

import (
	"github.com/go-playground/validator/v10"
)

type Image struct {
	ID     string `json:"id"`
	Label  string `json:"label" validate:"required"`
	Url    string `json:"url" validate:"required,url"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

func (i *Image) Validate() error {
	v := validator.New()
	return v.Struct(i)
}

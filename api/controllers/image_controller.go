package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/ashalfarhan/random-gallery/api/repository"
	"github.com/ashalfarhan/random-gallery/api/responses"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	responses.Success(w, http.StatusOK, "Hello")
}

func Ping(w http.ResponseWriter, r *http.Request) {
	responses.Success(w, http.StatusOK, "pong!")
}

func CreateImage(w http.ResponseWriter, r *http.Request) {
	ni, err := repository.CreateImage(r.Body)
	if err != nil {
		log.Println("Failed to create an image", err)
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusCreated, ni, nil)
}

func DeleteImage(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusNoContent, nil, nil)
}

func GetAllImages(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	l, p := q.Get("limit"), q.Get("page")
	if l == "" {
		l = "5"
	}
	if p == "" {
		p = "1"
	}

	page, err := strconv.Atoi(p)
	if err != nil {
		responses.Error(w,
			http.StatusBadRequest,
			errors.New("invalid page value"),
		)
		return
	}

	limit, err := strconv.Atoi(l)
	if err != nil {
		responses.Error(w,
			http.StatusBadRequest,
			errors.New("invalid limit value"),
		)
		return
	}

	imgs, err := repository.GetAllImages(limit, page)
	if err != nil {
		responses.Error(w,
			http.StatusBadRequest,
			err,
		)
		return
	}

	responses.JSON(w, http.StatusOK, imgs, nil)
}

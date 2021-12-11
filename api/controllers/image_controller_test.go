package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ashalfarhan/random-gallery/api/responses"
)

func TestDeleteImage(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/api/images/asd", nil)
	w := httptest.NewRecorder()
	DeleteImage(w, req)
	res := w.Result()
	defer res.Body.Close()
	var result responses.Response
	json.NewDecoder(res.Body).Decode(&result)
	if !result.Ok {
		t.Fatalf("\nexpected ok to be %v, but got: %v\n", true, result.Ok)
	}
	if result.StatusCode != http.StatusNoContent {
		t.Fatalf("\nexpected statusCode to be %v, but got: %v\n", http.StatusNoContent, result.StatusCode)
	}
}

func TestGetImages(t *testing.T) {
	t.Run("should be success", func(t *testing.T) {
		expected := 2
		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/images?limit=%d", expected), nil)
		w := httptest.NewRecorder()
		GetAllImages(w, req)
		res := w.Result()
		var result responses.Response
		json.NewDecoder(res.Body).Decode(&result)
		if !result.Ok {
			t.Fatalf("\nexpected ok to be %v, but got: %v\n", true, result.Ok)
		}
		l := len(result.Data.([]interface{}))
		if l != expected {
			t.Fatalf("\nexpected data len to be %v, but got: %v\n", expected, l)
		}
	})

	t.Run("should be fail", func(t *testing.T) {
		expected := -10000000
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/images?limit=%d", expected), nil)
		w := httptest.NewRecorder()
		GetAllImages(w, req)
		res := w.Result()
		var result responses.Response
		json.NewDecoder(res.Body).Decode(&result)
		if result.Ok {
			t.Fatalf("\nexpected ok to be %v, but got: %v\n", false, result.Ok)
		}
		if result.Error == nil {
			t.Fatalf("\nexpected error, but got: %v\n", result.Error)
		}
	})
}

package responses

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Ok         bool        `json:"ok"`
	StatusCode int         `json:"statusCode"`
	Error      interface{} `json:"error,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}, err error) {
	w.WriteHeader(statusCode)
	
	w.Header().Set("Content-Type", "application/json")

	res := Response{StatusCode: statusCode}

	if err != nil {
		res.Ok = false
		res.Error = err.Error()
	} else {
		res.Ok = true
		res.Data = data
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to encode json response of %v, Error: %s\n", res, err.Error())
		fmt.Fprint(w, err.Error())
	}
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, nil, err)
}

func Success(w http.ResponseWriter, statusCode int, data interface{}) {
	JSON(w, statusCode, data, nil)
}

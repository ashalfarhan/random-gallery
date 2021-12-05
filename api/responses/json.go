package responses

import (
	"encoding/json"
	"fmt"
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
	res := Response{StatusCode: statusCode, Data: data}
	if err != nil {
		res.Ok = false
		res.Error = err.Error()
		res.Data = nil
	} else {
		res.Ok = true
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		fmt.Fprint(w, err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, nil, err)
}

package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type ResponseAuth struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Token string `json:"token"`
	Data interface{} `json:"data"`
}

func ResponseToken(w http.ResponseWriter, message string, token string, data interface{}, status int) {
	res := ResponseAuth{
		Status: status,
		Message: message,
		Token: token,
		Data: data,
	}

	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Fatal(err)
	}
}

func ResponseMessage(w http.ResponseWriter, message string, data interface{}, status int) {
	res := Response{
		Status: status,
		Message: message,
		Data: data,
	}

	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Fatal(err)
	}
}

func ResponseError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	res := Response{
		Message: err.Error(),
		Data: nil,
	}

	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Fatal(err)
	}
}
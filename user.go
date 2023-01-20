package main

import (
	"encoding/json"
	"net/http"

	"github.com/tune-bot/database"
)

func Register(w http.ResponseWriter, req *http.Request) {
	user := database.User{}
	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		errorResponse(err, w)
		return
	}

	err = user.Create()

	if err != nil {
		errorResponse(err, w)
	} else {
		user.Password = ""
		data, err := json.Marshal(user)

		if err != nil {
			errorResponse(err, w)
		} else {
			successResponse(data, w)
		}
	}
}

func Login(w http.ResponseWriter, req *http.Request) {
	user := database.User{}
	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		errorResponse(err, w)
		return
	}

	err = user.Read()

	if err != nil {
		errorResponse(err, w)
	} else {
		user.Password = ""
		data, err := json.Marshal(user)

		if err != nil {
			errorResponse(err, w)
		} else {
			successResponse(data, w)
		}
	}
}

package main

import (
	"encoding/json"
	"net/http"

	"github.com/tune-bot/database"
)

func Link(w http.ResponseWriter, req *http.Request) {
	device := database.Device{}
	err := json.NewDecoder(req.Body).Decode(&device)

	if err != nil {
		errorResponse(err, w)
		return
	}

	err = device.Link()

	if err != nil {
		errorResponse(err, w)
	} else {
		data, err := json.Marshal(device)

		if err != nil {
			errorResponse(err, w)
		} else {
			successResponse(data, w)
		}
	}
}

func Get(w http.ResponseWriter, req *http.Request) {
	device := database.Device{}
	err := json.NewDecoder(req.Body).Decode(&device)

	if err != nil {
		errorResponse(err, w)
		return
	}

	user, err := device.GetUser()

	if err != nil {
		errorResponse(err, w)
	} else {
		data, err := json.Marshal(user)

		if err != nil {
			errorResponse(err, w)
		} else {
			successResponse(data, w)
		}
	}
}

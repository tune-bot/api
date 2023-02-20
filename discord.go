package main

import (
	"encoding/json"
	"net/http"

	"github.com/tune-bot/database"
)

func LinkDiscord(w http.ResponseWriter, req *http.Request) {
	discord := database.Discord{}
	err := json.NewDecoder(req.Body).Decode(&discord)

	if err != nil {
		errorResponse(err, w)
		return
	}

	err = discord.Link()

	if err != nil {
		errorResponse(err, w)
	} else {
		data, err := json.Marshal(discord)

		if err != nil {
			errorResponse(err, w)
		} else {
			successResponse(data, w)
		}
	}
}

func GetDiscord(w http.ResponseWriter, req *http.Request) {
	discord := database.Discord{}
	err := json.NewDecoder(req.Body).Decode(&discord)

	if err != nil {
		errorResponse(err, w)
		return
	}

	user, err := discord.GetUser()

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

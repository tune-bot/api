package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tune-bot/database"
)

func main() {
	port := "80"
	err := database.Connect()
	if err != nil {
		return
	}
	defer database.Disconnect()

	router := mux.NewRouter()

	router.HandleFunc("/user/register/", Register).Methods("POST")
	router.HandleFunc("/user/login/", Login).Methods("POST")
	router.HandleFunc("/user/device/link", LinkDevice).Methods("LINK")
	router.HandleFunc("/user/device/", GetDevice).Methods("GET")
	router.HandleFunc("/user/discord/link/", LinkDiscord).Methods("LINK")
	router.HandleFunc("/user/discord/", GetDiscord).Methods("GET")
	router.HandleFunc("/playlist/create/", Create).Methods("POST")
	router.HandleFunc("/playlist/update/", Update).Methods("PATCH")
	router.HandleFunc("/playlist/delete/", Delete).Methods("DELETE")
	router.HandleFunc("/playlist/song/add/", Add).Methods("PUT")
	router.HandleFunc("/playlist/song/remove/", Remove).Methods("DELETE")
	router.HandleFunc("/song/download/", Download).Methods("GET")

	http.ListenAndServe(":"+port, router)

	printLnColor("tune-bot api listening on port "+port, rotateSuccessColor())
}

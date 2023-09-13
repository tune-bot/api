package main

import (
	"net/http"

	"github.com/gorilla/mux"
	core "github.com/tune-bot/core/src"
)

func main() {
	port := "80"
	err := core.Connect()
	if err != nil {
		core.PrintError("Failed to connect to database: " + err.Error())
		return
	}
	defer core.Disconnect()

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
	router.HandleFunc("/song/search/", Search).Methods("GET")

	core.PrintSuccess("tune-bot api listening on port " + port)

	err = http.ListenAndServe(":"+port, router)

	if err != nil {
		core.PrintError("tune-bot api failed to start: " + err.Error())
		return
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/tune-bot/database"
	"github.com/valyala/fastjson"
)

func Add(w http.ResponseWriter, req *http.Request) {
	var parser fastjson.Parser
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(req.Body)
	raw, err := parser.Parse(buf.String())

	if err != nil {
		errorResponse(err, w)
		return
	}

	song := database.Song{Url: string(raw.GetStringBytes("url"))}

	err = song.AddToPlaylist(string(raw.GetStringBytes("playlistId")))
	if err != nil {
		errorResponse(err, w)
		return
	}
	data, err := json.Marshal(song)

	successResponse(data, w)
}

func Remove(w http.ResponseWriter, req *http.Request) {
	var parser fastjson.Parser
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(req.Body)
	raw, err := parser.Parse(buf.String())

	if err != nil {
		errorResponse(err, w)
		return
	}

	song := database.Song{Id: string(raw.GetStringBytes("songId"))}

	err = song.RemoveFromPlaylist(string(raw.GetStringBytes("playlistId")))
	if err != nil {
		errorResponse(err, w)
		return
	}

	successResponse([]byte("{}"), w)
}

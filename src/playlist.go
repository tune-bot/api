package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	core "github.com/tune-bot/core/src"
	"github.com/valyala/fastjson"
)

func Create(w http.ResponseWriter, req *http.Request) {
	var parser fastjson.Parser
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(req.Body)
	raw, err := parser.Parse(buf.String())

	if err != nil {
		errorResponse(err, w)
		return
	}

	playlist := core.Playlist{Name: string(raw.GetStringBytes("name")), Enabled: true, Songs: []core.Song{}}

	err = playlist.Create(string(raw.GetStringBytes("userId")))
	if err != nil {
		errorResponse(err, w)
		return
	}
	data, err := json.Marshal(playlist)
	successResponse(data, w)
}

func Update(w http.ResponseWriter, req *http.Request) {
	playlist := core.Playlist{}
	err := json.NewDecoder(req.Body).Decode(&playlist)

	if err != nil {
		errorResponse(err, w)
		return
	}

	err = playlist.Update()

	if err != nil {
		errorResponse(err, w)
		return
	}

	data, err := json.Marshal(playlist)
	successResponse(data, w)
}

func Delete(w http.ResponseWriter, req *http.Request) {
	playlist := core.Playlist{}
	err := json.NewDecoder(req.Body).Decode(&playlist)

	if err != nil {
		errorResponse(err, w)
		return
	}

	err = playlist.Delete()
	if err != nil {
		errorResponse(err, w)
		return
	}
	successResponse([]byte("{}"), w)
}
